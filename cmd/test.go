package main

import (
	"fmt"
	"github.com/tiyee/gokit/pkg/helps"
	"time"
)

func main() {
	//u := &User{Name: "a"}
	ts := time.Now().Unix()
	for i := int64(0); i < 24; i++ {
		_ts := ts + i*3600
		fs := helps.FirstSecond2(_ts)
		fmt.Println(helps.FormatTime(_ts), helps.FormatTime(fs))
	}

}
