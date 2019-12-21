#镜像
FROM golang:1.12.14
#作者
MAINTAINER Ving<"705105278@qq.com">
#将服务器的go工程代码加入到docker容器中

RUN mkdir -p /src/GP

COPY . /src/GP

WORKDIR /src/GP

#COPY . /go/src/

ENV PORT 9090

EXPOSE 9090

RUN go build src/GP/main.go

RUN chmod 777 /src/GP/main

CMD ["/src/GP/main"]

#EXPOSE 80
#最终运行docker的命令
#CMD ["GP"]
