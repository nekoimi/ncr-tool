# ncr-tool

[ncr-mirror](https://github.com/nekoimi/ncr-mirror) 命令行容器镜像拉取工具

### Install

- windows

在[Release](https://github.com/nekoimi/ncr-tool/releases)页面下载最新版并添加到环境变量

- Linux

```shell
curl -L -o /usr/local/bin/ncr-tool https://github.com/nekoimi/ncr-tool/releases/download/v0.0.3/ncr-tool-v0.0.3-linux-amd64 && chmod +x /usr/local/bin/ncr-tool
```

### Used

```shell
# 以拉取nginx的容器镜像为例
docker pull nginx:latest
```

- 使用ncr-tool直接拉取

```shell
ncr-tool pull nginx:latest
```

##### Tips

仅适用于Docker容器服务
