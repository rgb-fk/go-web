# go-web

一个基于 Go语言 的Web项目模板. 致力于构建完善的Go-Web开发框架, 以便快速高效的使用Go构建服务端. 示例网站: [瞎搞瞎玩](www.xiagaoxiawan.com)

## 特性
1. 使用 dep 管理项目依赖
1. 完整的项目结构, 并且可以使用Docker容器技术一键生成镜像并上传到阿里云服务器.
    - 阿里云私有容器服务上传需要先登录 `sudo docker login --username=name registry.cn-hangzhou.aliyuncs.com`
2. 使用 `.ini` 配置文件管理配置
3. 使用dao层管理数据库的查询
    - 目前已经封装了 mysql/redis/influx 数据库. 如需连接池支持, 还需开发.
4. 部分消息队列封装
    - 目前支持 kafka
5. 服务层分别采用 `net/http` 和 iris 框架实现. 二次开发时可以根据需要选择框架.
6. 支持微信公众号, 可以接收/响应微信公众号消息.

## 使用
1. 全项目替换 `github.com/everywan/go-web-demo` 为自己的项目路径.
2. 使用dep初始化项目, 确保所有的引用初始化到vendor中(dep/手动均可)
3. 符合 restful 规范

### 各框架比较
5. iris 封装了 GET/POST/PUT.. 等HTTP方法, 不需要像 net/http 一样在service函数里判断请求方法
#### iris
#### net/http
#### gin

## TODO
1. 替换掉 github.com/stackcats 内容

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

### AMQP
1. [kafka介绍]()
