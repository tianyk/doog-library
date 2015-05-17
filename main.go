package main

import (
	"github.com/astaxie/beego"
	_ "github.com/tianyk/doog-library/models"
	_ "github.com/tianyk/doog-library/routers"
)

func main() {
	beego.Run()
}
