package main

import (
	"blog/global"
	_ "blog/inits"
	_ "blog/routers"
	_ "blog/validate"
	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	beego.Run()
	defer global.SqlDB.Close()
}
