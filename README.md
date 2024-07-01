# ncr-tool

[ncr-mirror](https://github.com/nekoimi/ncr-mirror) 命令行容器镜像拉取工具

### Install

- windows

在[Release](https://github.com/nekoimi/ncr-tool/releases)页面下载最新版并添加到环境变量

- Linux

运行以下命令安装最新

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

### Uninstall

- linux

直接删除 ncr-tool 文件

```shell
rm -rf /usr/local/bin/ncr-tool
```


##### Tips

仅适用于Docker容器服务
