# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.21.3

# 设置工作目录
WORKDIR /server

# 将本地代码复制到容器中
COPY . .

# 修改镜像源
RUN go env -w GO111MODULE=on && \
    go env -w  GOPROXY=https://goproxy.cn,direct
    
# 编译 Go 应用程序
RUN cd greeter_server && \
    go build -o server

# 暴露应用程序的端口
EXPOSE 50051

# 启动应用程序
CMD ["./greeter_server/server"]

