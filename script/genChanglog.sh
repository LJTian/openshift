#!/bin/bash

filePath=$1

AllFileName=${filePath}/utccp-changelog.json



cat << EOF | tee ${AllFileName}
{
    "changelog":[
        {
            "date":"`date '+%Y-%m-%d'`",
            "version":"1.2.1-rc7-1",
            "author":"tianlijun",
            "changes":"build"
        }
    ]
}
EOF