package main

/*
curl http://localhost:8888/
curl --data "users=eilinge, duzi, lin" http://localhost:8888/import
curl http://localhost:8888/lucky
*/
import (
	"fmt"

	"eilieili/bootstrap"
	"eilieili/conf"
	"eilieili/eths"
	"eilieili/web/middleware/identity"
	"eilieili/web/routes"

	"github.com/kataras/iris"
	// "github.com/kataras/iris/mvc"
)

type eilieiliController struct {
	Ctx iris.Context
}

var port = 8080

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Go抽奖系统", "eilinge")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)

	return app
}
func main() {
	app := newApp()
	go eths.EventSubscribe("ws://localhost:8546", conf.Config.Eth.PxaAddr)
	app.Listen(fmt.Sprintf(":%d", port))
}
