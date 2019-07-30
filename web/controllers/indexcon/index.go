/**
 * 首页根目录的Controller
 * http://localhost:8080/
 */
package indexcon

import (
	// "time"
	// "strconv"
	"crypto/sha256"
	"fmt"
	"log"

	"eilieili/comm"
	"eilieili/conf"
	"eilieili/eths"
	"eilieili/models"
	"eilieili/services"
	"eilieili/web/utils"
	"eilieili/web/viewmodels"

	"github.com/gorilla/sessions"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"github.com/labstack/echo-contrib/session"
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

// GetLogin 登录 GET /login
func (c *IndexController) GetLogin() mvc.Result {
	return mvc.View{
		Name: "login.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "shared/layout.html",
	}
}

// GGetRegister 登录 GET /register
func (c *IndexController) GetRegister() mvc.Result {
	return mvc.View{
		Name: "register.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "shared/layout.html",
	}
}

// PostRegister http://localhost:8080/register
func (c *IndexController) PostRegister() mvc.Result {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)
	//2. 解析数据
	account := viewmodels.Accounts{}

	/*
		将前端传过来的数据, 与dbs.Accounts进行数据绑定
		&dbs.Account{
			Email       `json:"email"`			name="email"
			IdentitiyID `json:"identity_id"`	name="identity_id"
			UserName 	`json:"username"`		name="username"
		}
	*/

	// 读取表单("submit")
	err := c.Ctx.ReadJSON(&account)
	// fromValue := c.Ctx.Request().PostForm // 解析form表单
	if err != nil {
		// fmt.Println("failed to NewAcc: ", err)
		resp.Errno = utils.RECODE_IPCERR
		return nil
	}
	fmt.Printf("fromValue: %#v \n", account)

	//3. 操作geth创建账户(account.IdentitiyId->pass)
	// passwd := account.IdentitiyID
	address, err := eths.NewAcc(account.IdentitiyID, conf.Config.Eth.Connstr)
	if err != nil {
		// fmt.Println("failed to NewAcc: ", err)
		resp.Errno = utils.RECODE_IPCERR
		return nil
	}
	go func() {
		err = eths.EthErc20Transfer(conf.Config.Eth.Fundation, conf.Config.Eth.FundationPWD, address, 5)
		if err != nil {
			fmt.Println("Transfer failed when register err: ", err)
			return
		}
		// _, err = eths.EtherTransfer(conf.Config.Eth.Fundation, address)
	}()
	//4. 操作Mysql插入数据
	pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(account.IdentitiyID)))
	newAccount := &models.Account{
		Email:      account.Email,
		Username:   account.UserName,
		IdentityId: pwd,
		Address:    address,
	}
	err = c.ServiceAccount.Create(newAccount)
	if err != nil {
		log.Println("index.PostRegister c.ServiceAccount.Create err ", err)
		resp.Errno = utils.RECODE_DBERR
		return nil
	}
	//5. session处理
	sess, _ := session.Get("session", c.Ctx)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
	}
	sess.Values["address"] = address
	sess.Values["username"] = account.UserName
	sess.Save(c.Ctx.Request(), c.Ctx.Response())

}

// GetLogout 退出 GET /logout
func (c *IndexController) GetLogout() {
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=logout"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}

// GetSession session
func (c *IndexController) GetSession() error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c, &resp)

	sess, err := session.Get("session", c.Context)
	if err != nil {
		fmt.Println("failed to get session")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	address := sess.Values["address"]
	// username := sess.Values["username"]
	if address == nil {
		fmt.Println("failed to get address")
		resp.Errno = utils.RECODE_SESSIONERR
		return err
	}
	return nil
}
