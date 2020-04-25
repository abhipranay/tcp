#!/usr/bin/env bash
project_root=`pwd`
cmd_dir=$project_root/cmd
bin_dir=$project_root/bin
mkdir -p $bin_dir

cd $project_root
go get .
go build -o $bin_dir/server main.go
cd $cmd_dir
go build -o $bin_dir/client client.go