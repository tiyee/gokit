package main

import (
	"fmt"
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/component/log"
)

func main() {
	// 需要什么引用什么
	log.InitLogger()
	if err := component.InitMysql(); err != nil {
		fmt.Println("init mysql err:", err.Error())
		panic(err)
		//os.Exit(1)

	}
	log.Info("hello world")
}
