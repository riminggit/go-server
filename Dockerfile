#编译和运行都是用这个镜像
FROM golang:latest

#设置工作路径
WORKDIR /build

#把项目所有文件复制到镜像的根目录文件夹中
ADD . ./

# 设置Go语言的环境变量，打开Go Module模式。设置包下载源，有利于快速下载包
ENV GO111MODULE=on \
GOPROXY=https://goproxy.cn

#下载go.mod里面的包
RUN go mod download

#编译
RUN go build -o gin_docker .

#打开端口
EXPOSE 8080

#运行项目
ENTRYPOINT  ["./go-server-template"]
