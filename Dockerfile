FROM golang:1.19 AS builder

# 为我们的镜像设置必要的环境变量
ENV	GOPROXY="https://goproxy.cn,direct"
	
WORKDIR /app

# 将代码复制到容器中
COPY go.mod go.sum ./

# 将我们的代码编译成二进制可执行文件  可执行文件名为 app
RUN go mod download

COPY . .

# Build the Go app for specific architecture
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -o kubefox .

# Stage 2: Copy the binary file into a scratch container
FROM scratch

WORKDIR /app

# Copy the binary file from builder
COPY --from=builder /app/kubefox .
COPY --from=builder /app/setting.yaml  .
COPY --from=builder /app/static ./static
# Expose port 8081 to the outside
EXPOSE 8081

# Command to run the executable
ENTRYPOINT ["./kubefox"]