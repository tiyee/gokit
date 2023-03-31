# GoKit

一套go的项目模板，支持docker,尽可能的减少依赖


## Install

1. 直接clone代码 `git clone https://github.com/tiyee/gokit.git your_app`
2. 全局替换`github.com/tiyee/gokit/`,改成你自己的命名空间
3. **Makefile**文件的`SERVICE`变量改成需要的名称
4. **Dockerfile**和**docker-compose**里的`gokit`字样也需要更改

## Principle
1. 尽可能减少依赖
2. 白名单原则
3. 只保留唯一且简介的实现

## 自动生成go文件
1. 可以参考`cmd/go_log.go`