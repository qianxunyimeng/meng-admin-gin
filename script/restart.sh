#!/bin/bash
echo "删除进程"
killall meng-admin #杀死运行中的go-admin服务进程
echo "启动进程"
nohup ./meng-admin -c ./config.prod.yaml >> access.log 2>&1 & #后台启动服务将日志写入access.log文件
ps -aux | grep meng-admin #查看运行用的进程