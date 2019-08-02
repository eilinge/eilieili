package indexcon

import (
	"crypto/sha256"
	"fmt"
	"log"

	"eilieili/comm"
	"eilieili/conf"
	"eilieili/eths"
	"eilieili/models"
	"eilieili/web/utils"
	"eilieili/web/viewmodels"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

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

// GetRegister 登录 GET /register
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
func (c *IndexController) PostRegister() {
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
		return
	}
	// fmt.Printf("fromValue: %#v \n", account)

	//3. 操作geth创建账户(account.IdentitiyId->pass)
	// passwd := account.IdentitiyID
	address, err := eths.NewAcc(account.IdentitiyID, conf.Config.Eth.Connstr)
	if err != nil {
		// fmt.Println("failed to NewAcc: ", err)
		resp.Errno = utils.RECODE_IPCERR
		return
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
		return
	}
	//5. cookies处理
	// 每次随机生成一个登录用户信息
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: account.UserName,
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}

	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
	c.setSession(account.UserName, account.IdentitiyID)
	// username, passwd := c.getSession()
	// fmt.Println("the user: ", username, passwd)
	comm.Redirect(c.Ctx.ResponseWriter(), "/login")
}

// PostLogin ...
func (c *IndexController) PostLogin() {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)
	//2. 解析数据
	account := viewmodels.Accounts{}
	err := c.Ctx.ReadJSON(&account)
	if err != nil {
		fmt.Println("failed index.PostLogin c.Ctx.ReadJSON err: ", err)
		resp.Errno = utils.RECODE_IPCERR
		return
	}

	//3. 操作Mysql查询数据
	pwd := fmt.Sprintf("%x", sha256.Sum256([]byte(account.IdentitiyID)))
	acc, err := c.ServiceAccount.GetByUserName(account.UserName)
	if err != nil || pwd != acc.IdentityId {
		fmt.Println("failed index.PostLogin GetByUserName err: ", err)
		resp.Errno = utils.RECODE_IPCERR
		return
	}
	//5. cookies处理
	c.setSession(account.UserName, account.IdentitiyID)
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:      uid,
		Username: account.UserName,
		Now:      comm.NowUnix(),
		Ip:       comm.ClientIP(c.Ctx.Request()),
	}
	// username, passwd := c.getSession()
	// fmt.Println("the user: ", username, passwd)
	comm.SetLoginuser(c.Ctx.ResponseWriter(), &loginuser)
}

// GetLogout 退出: 清除cookies GET /logout
func (c *IndexController) GetLogout() {
	refer := c.Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/login"
	}
	comm.SetLoginuser(c.Ctx.ResponseWriter(), nil)
	c.deleteSession()
	comm.Redirect(c.Ctx.ResponseWriter(), refer)
}

// setSession ...
func (c *IndexController) setSession(username, passwd string) {
	//set session values
	s := comm.MySessions.Start(c.Ctx)
	log.Println("start store my session ...")
	s.Set("name", username)
	s.Set("passwd", passwd)
}

// getSession ...
func (c *IndexController) getSession() (username, passwd string) {
	s := comm.MySessions.Start(c.Ctx)
	username = s.GetString("name")
	passwd = s.GetString("passwd")
	log.Println("get name, passwd from session", username, passwd)
	return
}

// deleteSession ...
func (c *IndexController) deleteSession() {
	s := comm.MySessions.Start(c.Ctx)
	s.Delete("name")
	s.Delete("passwd")
}
