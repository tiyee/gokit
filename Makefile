.PHONY: usage  build run forever cron

PWD := $(shell pwd)
DS := /
SERVICE := gokit


GCFLAG := -gcflags='all=-N -l'
PID=./logs/$(SERVICE).pid
export GO111MODULE=on
export GOPROXY=https://goproxy.cn
export GOOS=linux
export GOARCH=amd64
export CGO_ENABLED=0

default: usage

usage:
	@echo
	@echo "-> usage:"
	@echo "make build \t\t\t 编译"
	@echo "make vet  \t\t\t 当前项目代码静态检查"
	@echo "make clean \t\t\t 清理"
env:
	@go mod tidy
	@go mod download
vet:env
	@echo '->[$(SERVICE)] 正在检查代码'
	@go vet ./cmd
	@echo '->[$(SERVICE)] 检查代码完成'

build: vet
	@echo '->[$(SERVICE)] 正在构建'
	@$(if $(wildcard bin), , mkdir -p bin)
	@go build -o bin$(DS)$(SERVICE) $(GCFLAG) cmd/main.go
	@echo '->[$(SERVICE)] 构建完成'
clean:
	@rm -rf  $(PWD)$(DS)bin








