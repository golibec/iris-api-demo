package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var public_url = []string{
	"/debug",
}


func main() {
	app := iris.New()

	app.Use(recover.New())
	app.Use(logger.New())

	app.Get("*", handler)
	app.Post("*", handler)

	app.Run(iris.Addr(":8081"), iris.WithoutServerError(iris.ErrServerClosed))
}

func handler(ctx iris.Context){
	if public_url[ctx.Path()] == nil {

	}
	ctx.Writef("Hello from method: %s and path: %s", ctx.Method(), ctx.Path())
}