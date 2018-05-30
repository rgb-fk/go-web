package controller

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
)

func StartWebServer_iris() {
	app := iris.New()
	app.Use(logger.New(logger.Config{
		Status: true,
		IP:     true,
		Method: true,
		Path:   true,
	}))

	// url分组
	// demo := app.Party("/v1/demo", middlewares.ErrorHandler)
	// {
	// 	demo.Get("/{key}", handler.DemoGet)
	// 	demo.Post("/", handler.DemoPost)
	// 	demo.Patch("/{key}", handler.DemoPatch)
	// 	demo.Delete("/{key}", handler.DemoDelete)
	// }

	app.OnErrorCode(404, func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"code": "404",
		})
	})

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch,
			os.Interrupt,
			syscall.SIGINT,
			os.Kill,
			syscall.SIGKILL,
			syscall.SIGTERM,
		)
		select {
		case <-ch:
			println("shutdown...")
			timeout := 5 * time.Second
			ctx, cancel := context.WithTimeout(context.Background(), timeout)
			defer cancel()
			app.Shutdown(ctx)
		}
	}()

	app.Run(iris.Addr(":50000"), iris.WithoutVersionChecker)
}
