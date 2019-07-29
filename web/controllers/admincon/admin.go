/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package admincon

import (
	"github.com/kataras/iris"

	"lottery/services"

	"github.com/kataras/iris/mvc"
)

// AdminController 其他用户访问界面
type AdminController struct {
	Ctx            iris.Context // 解析前端传来的数据
	ServiceUser    services.UserService
	ServiceGift    services.GiftService
	ServiceCode    services.CodeService
	ServiceResult  services.ResultService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080/
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