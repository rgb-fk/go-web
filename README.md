# go-web-demo

1. 更改项目时: 全项目搜索 `github.com/everywan/go-web-demo` 改为自己的项目路径
2. 使用dep初始化项目, 确保所有的引用初始化到vendor中(dep/手动均可)
3. 符合 restful 规范

4. 替换掉 github.com/stackcats 内容

## 各文件介绍
### docker.sh
> [源码](docker.sh)
1. 用于自动构建docker镜像, 并且上传到阿里云容器.
    - 默认镜像名称: `basename $PWD`:$version`

### Dockerfile
> [源码](Dockerfile)
1. 实际构建docker的程序. 使用 Dockerfile 方式构建镜像

### Makefile
> [源码](Makefile)
1. 用于构建可执行的二进制文件

### cache
1. redis 的客户端, 用于缓存

### config
1. 配置各数据库客户端
