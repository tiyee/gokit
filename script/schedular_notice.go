package main

import (
	"fmt"
	"github.com/tiyee/gokit/pkg/component"
)

func main() {
	// 需要什么引用什么
	component.InitLogger()
	if err := component.InitMysql(); err != nil {
		fmt.Println("init mysql err:", err.Error())
		panic(err)
		//os.Exit(1)

	}
	component.Logger.Info("hello world")
}
