# 使用官方的 Golang 镜像作为基础镜像
FROM golang:1.21.3

# 设置工作目录
WORKDIR /client

# 将本地代码复制到容器中
COPY . .

# 修改镜像源
RUN go env -w GO111MODULE=on && \
    go env -w  GOPROXY=https://goproxy.cn,direct
    
# 编译 Go 应用程序  
RUN cd greeter_client && \
    go build -o client

# 启动应用程序
CMD ["./greeter_client/client"]

