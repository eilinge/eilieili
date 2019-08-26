/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package admincon

import (
	"github.com/kataras/iris"

	"eilieili/services"

	"github.com/kataras/iris/mvc"
)

// AdminController 其他用户访问界面
type AdminController struct {
	Ctx                   iris.Context // 解析前端传来的数据/响应数据给后端
	ServiceAccount        services.AccountService
	ServiceAccountContent services.AccountContentService
	ServiceAuction        services.AuctionService
	ServiceBidwinner      services.BidwinnerService
	ServiceVote           services.VoteService

	ServiceUser    services.UserService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080/admin
func (c *AdminController) Get() mvc.Result {
	return mvc.View{
		Name: "admin/index.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "admin/layout.html",
	}
}
