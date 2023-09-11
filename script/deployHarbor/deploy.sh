#/bin/bash

# 加载入参环境变量文件
source `pwd`/env_etc
toolsList=("podman" "openssl" "tar" "md5sum")
dataAllName=`pwd`/data/$DATA_NAME
# 输出入参值
printENV(){
	echo " 入参内容如下，如果为 NULL 请先进行配置。配置文件为：[`pwd`/env_etc]"
	echo ""
	echo ""
	echo " DATA_NAME        is [$DATA_NAME]"
	echo " DATA_MD5         is [$DATA_MD5]"
	echo " DOMAIN           is [$DOMAIN]"
	echo " INSTALL_PATH     is [$INSTALL_PATH]"
	echo " CONTAINER_NAME   is [$CONTAINER_NAME]"
	echo ""
}

checkTool(){
	# echo "校验需要的工具是否安装"
	# 使用 which 命令检查工具是否存在
	if which $1 >/dev/null 2>&1; then
		echo "$1 已安装."
	else
		echo "$1 未安装."
		exit 64
	fi
}

checkTools(){

	for tool in "${toolsList[@]}"
	do
		checkTool $tool 
	done
}

checkData(){
	echo "开始计算 MD5 值，时间比较久，请耐性等待..."

	if [ 0`md5sum $dataAllName` != 0$DATA_MD5 ] ; then 
		echo "md5 值不一致"
		exit 64
	fi
}

loadContain(){
	podman load -i `pwd`/data/$CONTAINER_NAME
	if [ 0$? != 00 ]; then 
		echo "加载失败 返回值为[$?]"
		exit 64
	fi
}

creatCerts(){
	mkdir -p $INSTALL_PATH/certs
	cd $INSTALL_PATH/certs
	openssl genrsa -out ca.key 2048
	openssl req -x509 -new -nodes -key ca.key -subj "/CN=$DOMAIN" -days 3650 -out ca.crt
	openssl req -new -sha256  -key ca.key -subj    "/C=CN/ST=GD/L=SZ/O=Global Security/OU=IT Department/CN=$DOMAIN" -reqexts SAN  -config <(cat /etc/pki/tls/openssl.cnf <(printf "[SAN]\nsubjectAltName=DNS:$DOMAIN"))  -out registry.csr
	openssl x509 -req -days 3650  -in registry.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extfile <(printf "subjectAltName=DNS:$DOMAIN") -out  registry.crt
	cp *.crt /etc/pki/ca-trust/source/anchors/
	update-ca-trust
}

tarData(){
	echo "解压tar包，时间比较久，请耐性等待..."
	mkdir -p $INSTALL_PATH/data/
	tar xf $dataAllName -C $INSTALL_PATH/data/
	if [ 0$? != 00 ]; then
		echo "解压失败 返回值为[$?]"
		echo "清理垃圾缓存"
		rm -rf $INSTALL_PATH/data/
		exit 64
	fi	
}

deployData(){
	if [ -d $INSTALL_PATH ] ; then 
		echo "$INSTALL_PATH 已经存在"
		#exit 64
	fi

	mkdir -p $INSTALL_PATH
	#tarData
	creatCerts
	dataDirName=`ls $INSTALL_PATH/data/`
	podman run -d --restart=always --name uccps-registry -v $INSTALL_PATH/certs:/certs:z -e REGISTRY_HTTP_ADDR=0.0.0.0:443 -e REGISTRY_HTTP_TLS_CERTIFICATE=/certs/registry.crt -e REGISTRY_HTTP_TLS_KEY=/certs/ca.key -p 443:443 -v $INSTALL_PATH/data/$dataDirName/registry:/var/lib/registry:z localhost/registry
}

testHub(){
	echo 127.0.0.1 $INSTALL_PATH
	podman pull $DOMAIN/$TEST_IMAGE
	if [ 0$? != 00 ] ; then 
		echo "拉取镜像失败 返回值为[$?]"
		exit 64
	fi
}

# 打印环境变量
printENV
# 检验所需要的工具是否安装
checkTools
# 检验离线仓库数据MD5
#checkData
# 加载镜像仓库
loadContain
# 部署离线仓库
deployData
# 测试署情况
testHub

