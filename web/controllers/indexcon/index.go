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

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"

	"eilieili/comm"
	"eilieili/conf"
	"eilieili/eths"
	"eilieili/models"
	"eilieili/services"
	"eilieili/web/utils"
	"eilieili/web/viewmodels"
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
func (c *IndexController) Get() mvc.Result {
	return mvc.View{
		Name: "shared/login.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "shared/layout.html",
	}
}

// PostRegister http://localhost:8080/
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
	err := c.Ctx.ReadForm(&account)
	if err != nil {
		log.Println("c.Ctx.ReadForm ERR: ", err)
	}
	// fromValue := c.Ctx.Request().PostForm // 解析form表单
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
	sql := fmt.Sprintf("insert into account(email, username, identity_id, address) values('%s', '%s', '%s', '%s')",
		account.Email, account.UserName, pwd, address)
	fmt.Println(sql)
	// _, err = dbs.Create(sql)
	if err != nil {
		resp.Errno = utils.RECODE_DBERR
		return nil
	}
	//5. session处理
	// sess, _ := session.Get("session", c.Ctx)
	// sess.Options = &sessions.Options{
	// 	Path:     "/",
	// 	MaxAge:   86400 * 7, // 7 days
	// 	HttpOnly: true,
	// }
	// sess.Values["address"] = address
	// sess.Values["username"] = account.UserName
	// sess.Save(c.Ctx.Request(), c.Ctx.Response())
	return nil

	return mvc.View{
		Name: "shared/login.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "",
		},
		Layout: "shared/layout.html",
	}
}

// GetLogin 登录 GET /login
func (c *IndexController) GetLogin() {
	// 每次随机生成一个登录用户信息
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: fmt.Sprintf("admin-%d", uid),
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/public/index.html?from=login"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
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
