package main

import (
	_ "learnbeego/learnbeego/routers"
	"github.com/astaxie/beego"
)

func main() {
	fmt.print('hello')
	beego.Run()
}

