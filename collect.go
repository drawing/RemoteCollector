package main

import (
	_ "./routers"
	"github.com/astaxie/beego"
)

func main() {
	beego.TemplateLeft = "{%"
	beego.TemplateRight = "%}"

	beego.Run()
}
