package main

import (
	"fmt"
	"github.com/tiyee/gokit/pkg"
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/engine"
)

func main() {
	e := engine.New()
	pkg.LoadRouter(e)

	component.InitLogger()
	if err := component.InitMysql(); err != nil {
		fmt.Println("init mysql err:", err.Error())
		panic(err)
		//os.Exit(1)

	}
	if err := e.Run(); err != nil {
		fmt.Println(err.Error())
	}

}
