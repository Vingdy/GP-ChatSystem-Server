#镜像
FROM golang:1.12.14
#作者
MAINTAINER Ving<"705105278@qq.com">
#将服务器的go工程代码加入到docker容器中
COPY GP /bin/GP
#最终运行docker的命令
CMD ["GP"]