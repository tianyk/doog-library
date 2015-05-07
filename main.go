package main

import (
	_ "doog-library/models"
	_ "doog-library/routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}
