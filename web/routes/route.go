package routes

import (
	"github.com/kataras/iris/mvc"

	"eilieili/bootstrap"
	"eilieili/services"
	"eilieili/web/controllers/admincon"
	"eilieili/web/controllers/indexcon"
	"eilieili/web/middleware"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	userService := services.NewUserService()
	giftService := services.NewGiftService()
	codeService := services.NewCodeService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewBlackipService()

	// 路由分组IndexController/AdminController
	// "/" IndexController
	index := mvc.New(b.Party("/"))
	index.Register(userService,
		giftService,
		codeService,
		resultService,
		userdayService,
		blackipService)
	index.Handle(new(indexcon.IndexController))

	// "/admin" AdminController
	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(userService,
		giftService,
		codeService,
		resultService,
		userdayService,
		blackipService)
	admin.Handle(new(admincon.AdminController))

	// "/admin/gift" AdminController
	adminGift := admin.Party("/gift")
	adminGift.Register(giftService)
	adminGift.Handle(new(admincon.AdminGiftController))

	// "/admin/code" AdminController
	adminCode := admin.Party("/code")
	adminCode.Register(codeService)
	adminCode.Handle(new(admincon.AdminCodeController))

	// "/admin/result" AdminController
	adminResult := admin.Party("/result")
	adminResult.Register(resultService)
	adminResult.Handle(new(admincon.AdminResultController))

	// "/admin/user" AdminController
	adminUser := admin.Party("/user")
	adminUser.Register(userService)
	adminUser.Handle(new(admincon.AdminUserController))

	// "/admin/user" AdminController
	adminBlackip := admin.Party("/blackip")
	adminBlackip.Register(blackipService)
	adminBlackip.Handle(new(admincon.AdminBlackipController))

	// "/rpc"
	rpc := mvc.New(b.Party("rpc"))
	rpc.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	rpc.Handle(new(indexcon.IndexController))

	// 传统设置路由
	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
