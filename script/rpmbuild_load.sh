#!/bin/bash

echo "rpmbuild_load start"

PWD=`pwd`
ACTION=$1
echo ${ACTION}
if [ 1"${ACTION}" == 1 ]; then
        ACTION="bp"
fi

mkdir -p ${PWD}/rpmbuild/{BUILD,BUILDROOT,RPMS,SOURCES,SPECS,SRPMS}
cp *.spec ${PWD}/rpmbuild/SPECS
files=`grep -E "Source|Patch" *.spec | awk '{print $2}'`

for file in $files
do
        echo ${file}
        cp ${file} ${PWD}/rpmbuild/SOURCES/
done

rpmbuild --define="_topdir `pwd`/rpmbuild" --nodebuginfo -${ACTION} ${PWD}/rpmbuild/SPECS/*.spec
