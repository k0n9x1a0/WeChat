# go docker服务


now := time.Now().Unix()


# 远程调试

1. 复制代码到服务器;
2. 服务器进入代码目录并执行dlv;
```bash
cd /root/go_project/wic-go/
dlv debug --headless --listen=:2346 --api-version=2
```
3.客户端运行go-remote远程调试

# http专用

cd /home/docker-go
docker build -f ./Dockerfile.http -t wic-go-http:1.0.2 .
docker service remove wic-go-http
docker service create --replicas 16 --network wic-business --name wic-go-http -p 8005:8005 wic-go-http:1.0.2

# tcp专用

cd /home/docker-go
docker build -f ./Dockerfile.tcp -t wic-go-tcp:1.0.2 .
docker service remove wic-go-tcp
docker service create --replicas 1 --network wic-business --name wic-go-tcp -p 8006:8006 wic-go-tcp:1.0.2




Windows 下编译 Mac 和 Linux 64位可执行程序

SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build -o mac  main.go


SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build -o linuxService main.go


go build -o main.exe main.go
