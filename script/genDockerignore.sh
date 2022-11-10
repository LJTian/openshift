#!/bin/bash

filePath=$1

AllFileName=${filePath}/.dockerignore

cat << EOF | tee ${AllFileName}
Dockerfile-utccp-operators*
EOF