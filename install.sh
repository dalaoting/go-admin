#!/bin/bash

BIN=$1

echo "开始构建 go-admin"
go build
echo "构建完成"

TAR_FILE=video-merge.tar.gz
MV_TAR=/usr/local/go-admin/api/

echo "开始安装video-merge"
echo "打包中..."
tar -czvf $TAR_FILE go-admin ./config restart.sh
echo "打包完成, 开始安装"

mv $TAR_FILE $MV_TAR
cd $MV_TAR
mv go-admin go-admin-bak
rm -rf ./config
tar -zxvf $TAR_FILE
sh restart.sh $BIN