/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package indexcon

import (
	// "time"
	// "strconv"

	"eilieili/comm"
	"eilieili/services"
	"eilieili/web/utils"

	"github.com/kataras/iris"
)

// IndexController ...
type IndexController struct {
	Ctx                   iris.Context
	ServiceAccount        services.AccountService
	ServiceAccountContent services.AccountContentService
	ServiceBidwinner      services.BidwinnerService
	ServiceVote           services.VoteService

	ServiceUser    services.UserService
	ServiceUserday services.UserdayService
	ServiceBlackip services.BlackipService
}

// Get http://localhost:8080/
func (c *IndexController) Get() {
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html"
	}
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}

// GetSession /session
func (c *IndexController) GetSession() error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)
	//2. 通过session 获取用户名
	c.getSession()
	return nil
}
