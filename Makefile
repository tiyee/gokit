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
	@echo "make proto \t\t\t 根据proto更新当前项目的pb.go和pb.gw.go"
	@echo "make vet  \t\t\t 当前项目代码静态检查"
	@echo "make build \t\t\t 构建当前项目"
	@echo "make run  \t\t 运行当前项目"
	@echo "make forever \t\t\t deamon 运行"
	@echo
proto:
	@echo 'create proto'
env:
	@go mod tidy
vet:
	@echo '->[$(SERVICE)] 正在检查代码'
	@go vet ./pkg
	@echo '->[$(SERVICE)] 检查代码完成'

build: vet
	@echo '->[$(SERVICE)] 正在构建'
	@$(if $(wildcard bin), , mkdir bin)
	@go build -o bin$(DS)$(SERVICE) $(GCFLAG) cmd/main.go
	@echo '->[$(SERVICE)] 构建完成'
cron:
	@go build -o bin$(DS)notice $(GCFLAG) script/schedular_notice.go

run: build
	@echo '->[$(SERVICE)] 正在启动'
	@bin$(DS)$(SERVICE)  -usr1 default
forever:build
	@echo '->[$(SERVICE)] 正在启动'
	@nohup bin/${SERVICE} -usr1 default  >> logs/output.log 2>&1 &


clean:
	@rm -rf  $(PWD)$(DS)bin








