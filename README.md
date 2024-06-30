# ncr-tool

[ncr-mirror](https://github.com/nekoimi/ncr-mirror) 命令行容器镜像拉取工具

### 使用

```shell
# 以拉取nginx的容器镜像为例

docker pull docker.io/library/nginx:latest
```

- 使用[ncr-mirror](https://github.com/nekoimi/ncr-mirror)代理拉取:

```shell
# 1、先通过 ncr-mirror 代理拉取镜像
docker pull docker.mirror.403forbidden.run/library/nginx:latest
# 2、重命名为原tag
docker tag docker.mirror.403forbidden.run/library/nginx:latest docker.io/library/nginx:latest
```

- 使用ncr-tool直接拉取

```shell
ncr-tool pull docker.io/library/nginx:latest
```

##### Tips

仅适用于Docker容器服务
