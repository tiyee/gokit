# GoKit

一套go的项目模板，支持docker,尽可能的减少依赖.


## Install

1. 直接clone代码 `git clone https://github.com/tiyee/gokit.git your_app`
2. 全局替换`github.com/tiyee/gokit/`,改成你自己的命名空间
3. **Makefile**文件的`SERVICE`变量改成需要的名称
4. **Dockerfile**和**docker-compose**里的`gokit`字样也需要更改
5. 不同的运行平台，对应的Makefile里的`GOOS`环境变量也要改。
6. Dockerfile里的`alpine`镜像地址是阿里云的镜像源，请根据实际情况修改

## config

请自行更改 `pkg/consts`文件夹里文件内容，配置文件最好做成独立的文件配置化。


## 基于公共或私有镜像仓库部署(本文档以阿里云的acr为例)

可以自行执行docker命令打包并上传，也可以利用代码仓库自带的pipeline或action自动触发打包或上传。

部署的时候，只需要执行两步

1. `sudo docker-compose -f docker-compose.acr.yml pull `
2. `sudo docker-compose -f docker-compose.acr.yml up -d `

> 较新版本的docker，可以用`docker compose`命令组合来代替`docker-compose`





## Principle

1. 尽可能减少依赖
2. 白名单原则
3. 只保留唯一且简洁的实现


## 组件(componet)的使用

目前内置了`log,cache,redis,iden,orm,jwt,mysql`组件，其中`log,redis,mysql,cache`需要再项目启动的时候初始化才能使用。


## 自动生成go文件

1. 可以参考`cmd/log_gen/main.go`,在项目根目录执行`go run cms/log_gen/main.log`即可生成包装了zap的log


## Demo

自带了一个微信公众平台的echo服务，和一个分片上传oss的例子。

## Changelog

* 2023-05-26 v2.0.0版本将`http server`由`fasthttp`改回go自带的`http server`模块
