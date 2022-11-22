#!/bin/bash

dirName=$1
echo ${dirName}

IS_INSTALLED=$(rpm -qa |grep rpcbind)

if [ $? -eq 0 ] ; then
    echo 'rpcbind installed'
else
    echo 'rpcbind not installed'
fi

if [ $? -eq 0 ] ; then
    echo 'rpcbind installed'
else
    echo 'rpcbind not installed'
fi

if [ ! -n "${dirName}" ]; then
  echo "入参为NULL"
  echo "注意：请输入全路径，并使用root用户执行！"
fi

if [ -d ${dirName} ] ; then
  mkdir -p ${dirName}
fi

chmod 777 ${dirName}
echo "${dirName} *(rw,sync,root_squash,insecure)" >> /etc/exports

systemctl restart nfs-server