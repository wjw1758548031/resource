# 使用官方的 Go 作为基础镜像
FROM golang:1.18 AS builder

# 设置工作目录
WORKDIR /app

# 复制 Go 模块和依赖
#COPY go.mod ./
#COPY go.sum ./
RUN #go mod download

# 复制源代码
COPY . ./

# 构建 Go 应用
RUN go build -o main main.go

# 使用一个更小的基础镜像（如alpine）
FROM alpine:latest

# 复制编译好的二进制文件到新镜像中
COPY --from=builder /app/main /app/main

# 确保二进制文件有执行权限
RUN chmod +x /app/main

# 设置容器启动时要执行的命令
ENTRYPOINT ["/app/main"]
