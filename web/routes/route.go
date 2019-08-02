package routes

import (
	"eilieili/bootstrap"
	"eilieili/services"
	"eilieili/web/controllers/admincon"
	"eilieili/web/controllers/indexcon"
	"eilieili/web/middleware"

	"github.com/kataras/iris/mvc"
)

// Configure registers the necessary routes to the app.
func Configure(b *bootstrap.Bootstrapper) {
	accountService := services.NewAccountService()
	accountContentvice := services.NewAccountContentService()
	auctionService := services.NewAuctionService()
	voteService := services.NewVoteService()
	bidwinnerService := services.NewBidwinnerService()

	userService := services.NewUserService()
	userdayService := services.NewUserdayService()
	blackipService := services.NewBlackipService()

	// 路由分组IndexController/AdminController
	// "/" IndexController
	index := mvc.New(b.Party("/"))
	index.Register(
		accountService,
		accountContentvice,
		auctionService,
		voteService,
		bidwinnerService,
		userService,
		userdayService,
		blackipService)
	index.Handle(new(indexcon.IndexController))

	// 管理用户/ip黑名单,
	// "/admin" AdminController
	admin := mvc.New(b.Party("/admin"))
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(
		accountService,
		accountContentvice,
		auctionService,
		voteService,
		bidwinnerService,
		userService,
		userdayService,
		blackipService)
	admin.Handle(new(admincon.AdminController))

	// "/admin/user" AdminController
	adminUser := admin.Party("/user")
	adminUser.Register(userService)
	adminUser.Handle(new(admincon.AdminUserController))

	// "/admin/user" AdminController
	adminBlackip := admin.Party("/blackip")
	adminBlackip.Register(blackipService)
	adminBlackip.Handle(new(admincon.AdminBlackipController))

	// "/rpc"
	// rpc := mvc.New(b.Party("rpc"))
	// rpc.Register(userService, giftService, codeService, resultService, userdayService, blackipService)
	// rpc.Handle(new(indexcon.IndexController))

	// 传统设置路由
	//b.Get("/follower/{id:long}", GetFollowerHandler)
	//b.Get("/following/{id:long}", GetFollowingHandler)
	//b.Get("/like/{id:long}", GetLikeHandler)
}
