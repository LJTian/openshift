#!/bin/bash

NFS_SERVER_IP=$1
NFS_SERVER_DIR=$2
if [ ! -n "${NFS_SERVER_IP}" ] ; then
    echo "nfs server IP 没有填写！"
    echo "格式 ：install_nfs.sh IP DIR"
    exit  64
fi

if [ ! -n "${NFS_SERVER_DIR}" ] ; then
    echo "nfs server PATH 没有填写！"
    echo "提示：$(showmount -e ${NFS_SERVER_IP})"
    echo "格式 ：install_nfs.sh IP DIR"
    exit  64
fi

NFS_SERVER_PATH=`showmount -e ${NFS_SERVER_IP} | grep ${NFS_SERVER_DIR} | awk '{print $1}'`
if [ ! -n "${NFS_SERVER_PATH}" ] ; then
    echo "未获取到正确的nfs PATH, 请检查!"
    exit 64
fi

echo "IP = [${NFS_SERVER_IP}] \nPATH = [$NFS_SERVER_PATH]\n"

mkdir -p ./yaml/
cp ./nfs_temp/*.yaml ./yaml/

echo "设置 nfs IP 和 PATH"
sed -i "s#<SED_NFS_IP>#${NFS_SERVER_IP}#g" ./yaml/deployment.yaml
sed -i "s#<SED_NFS_PATH>#${NFS_SERVER_DIR}#g" ./yaml/deployment.yaml

echo "设置完成，开始部署"
oc new-project nfs-storage
oc create -f ./yaml/rbac.yaml
oc adm policy add-scc-to-user hostmount-anyuid system:serviceaccount:nfs-storage:nfs-client-provisioner
oc create -f ./yaml/deployment.yaml -f ./yaml/class.yaml
echo "完成！！！！"