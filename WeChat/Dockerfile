# 使用国内源的 Go 镜像
FROM registry.cn-hangzhou.aliyuncs.com/imooc-study/golang:1.15.4 AS builder

# 设置工作目录
WORKDIR /usr/ipad849

# 设置 Go 环境变量
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct

# 将当前目录的所有文件添加到工作目录中
ADD . /usr/ipad849

# 下载 Go 依赖
RUN go mod download

# 构建 Go 应用程序
RUN go build -o main main.go

# 使用国内源的 Redis 镜像
FROM registry.cn-hangzhou.aliyuncs.com/imooc-study/redis:6.0.9

# 将构建阶段的 Go 应用程序复制到运行阶段
COPY --from=builder /usr/ipad849 /usr/ipad849

# 设置工作目录
WORKDIR /usr/ipad849

# 暴露应用程序端口
EXPOSE 8005

# 启动时执行 Go 应用程序和 Redis 服务器
CMD ["sh", "-c", "redis-server & ./main"]
