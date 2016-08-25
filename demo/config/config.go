package main

import (
	"fmt"
	"github.com/astaxie/beego/config"
)

func main() {
	// API
	confTest, err := config.NewConfig("ini", "../../conf/app.conf")
	val_bool, err := confTest.Bool("test_bool")
	if err == nil {
		fmt.Println("test bool", val_bool)
	} else {
		fmt.Println(confTest)
	}
}
