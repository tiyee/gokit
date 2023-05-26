package main

import (
	"github.com/tiyee/gokit/pkg/component"
	"github.com/tiyee/gokit/pkg/component/log"
	"github.com/tiyee/gokit/pkg/component/redis"
	"github.com/tiyee/gokit/pkg/engine"
	"github.com/tiyee/gokit/pkg/router"
)

func main() {
	e := engine.New()
	log.InitLogger()
	if err := component.InitMysql(); err != nil {
		log.Error("init mysql error", log.String("error", err.Error()))
		panic(err)
	} else {
		log.Info("init mysql success")
	}
	if err := redis.InitRedis(); err != nil {
		log.Error("init redis error", log.String("error", err.Error()))
		panic(err)
	} else {
		log.Info("init redis success")
	}
	router.LoadRouter(e)
	if err := e.Run(); err != nil {
		panic(err)
	}

}
