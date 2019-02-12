package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
)

var public_url = map[string]string{
	"/debug":"GET",
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
	checkSign(ctx)

	//if _, ok := public_url[ctx.Path()];ok {
	//	fmt.Println(ok)
	//}else{
	//	checkSign(ctx)
	//}

	ctx.Writef("Hello from method: %s and path: %s \n", ctx.Method(), ctx.Path())
}

func checkSign(ctx iris.Context){

	if method,ok := public_url[ctx.Path()];ok {
		if method == ctx.Method() {
			ctx.Writef("Right method: %s", method)
		}else{
			ctx.Writef("Wrond method: %s", method)
		}
	}else{

	}

	ctx.Writef("Method: %s \n", ctx.Method())
	ctx.Writef("Host: %s \n", ctx.Host())
	ctx.Writef("RemoteAddr: %s \n", ctx.RemoteAddr())
	ctx.Writef("GetHeader: %s \n", ctx.GetHeader("sign"))

}