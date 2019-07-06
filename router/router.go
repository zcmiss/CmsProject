package router

import (
	"fmt"

	"github.com/kataras/iris"
)

const (
	adminPrefix    = "admin"
	categoryPrefix = "category"
	shoppingPrefix = "shopping"
	v1Prefix       = "v1"
)

func RouteAdmin(app *iris.Application) {
	fmt.Println("这是router  . . . ")
}
