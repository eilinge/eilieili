package indexcon

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"strconv"

	"eilieili/comm"
	"eilieili/datasource"
	"eilieili/eths"
	"eilieili/models"
	"eilieili/services"
	"eilieili/web/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// GetIndex 登录 GET /index
func (c *IndexController) GetIndex() mvc.Result {
	return mvc.View{
		Name: "user/userindex.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "index",
		},
		Layout: "shared/indexlayout.html",
	}
}

// GetImage 登录 GET /image
func (c *IndexController) GetImage() mvc.Result {
	return mvc.View{
		Name: "user/imageAuthor.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "index",
		},
		Layout: "shared/indexlayout.html",
	}
}

// PostContent ...
func (c *IndexController) PostContent() error {
	//1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)
	//2. 解析数据
	content := models.Content{}

	price, _ := strconv.ParseInt(c.Ctx.FormValue("price"), 10, 32)
	weight, _ := strconv.ParseInt(c.Ctx.FormValue("weight"), 10, 32)
	// fmt.Printf("fromValue: %#v, %#v\n", price, weight)
	content.Price = int(price)
	content.Weight = int(weight)

	_, h, err := c.Ctx.FormFile("fileName")
	if err != nil {
		fmt.Println("failed to FormFile: ", err)
		resp.Errno = utils.RECODE_PARAMERR
		return err
	}

	src, err := h.Open()
	defer src.Close()
	// 3. 打开一个目标文件用于保存
	content.Content = "public/photo/" + h.Filename
	dst, err := os.Create(content.Content)
	if err != nil {
		fmt.Println("failed to create file: ", err)
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	defer dst.Close()

	// 4. get hash
	cData := make([]byte, h.Size)
	n, err := src.Read(cData)
	if err != nil || h.Size != int64(n) {
		resp.Errno = utils.RECODE_IOERR
		return err
	}
	content.ContentHash = fmt.Sprintf("%x", sha256.Sum256(cData))
	dst.Write(cData) // 图片存储

	content.Title = h.Filename
	// 5. write to dbs / 给上传图片页面, 添加weight, price, 并一起传入
	// content.AddContent()

	// 6. 操作以太坊
	user := comm.GetLoginUser(c.Ctx.Request())
	// username, passwd := c.getSession()
	// fmt.Println("the user: ", username)
	userobj, _ := c.ServiceAccount.GetByUserName(user.Username)
	if err != nil {
		log.Println("user_index.PostContent GetByUserName err: ", err)
		resp.Errno = utils.RECODE_USERERR
		return err
	}
	fromAddr := userobj.Address
	log.Println("fromAddr: ", fromAddr)
	passwd := "eilinge"
	// from, pass, hash, data string
	// fmt.Printf("price: %d, weight: %d\n", price, weight)

	// 使用go func开启协程, 则当挖矿失败, 无法返回resp.Errno
	err = eths.Upload(fromAddr, passwd, content.ContentHash, content.Title, price, weight)
	if err != nil {
		resp.Errno = utils.RECODE_MINTERR
		return err
	}
	ts := comm.NowUnix()
	content.Ts = ts
	err = services.NewContentService().Create(&content)
	if err != nil {
		log.Println("user_index.PostContent create err ", err)
		resp.Errno = utils.RECODE_DATAERR
		return err
	}
	return nil
}

// GetContents ...
func (c *IndexController) GetContents() mvc.Result {
	//1. 通过cookies, 获取address, 然后读取出其所有资产
	userObj := comm.GetLoginUser(c.Ctx.Request())
	acc, err := c.ServiceAccount.GetByUserAddr(userObj.Username)
	if err != nil || acc.Address == "" {
		log.Println("failed to GetByUserAddr err ", err)
		return nil
	}
	log.Println("acc obj address: ", acc.Address)
	dao := utils.NewContentinfoService(datasource.InstanceDbMaster())
	contents, num, err := dao.InnerContent(acc.Address)
	if err != err || num <= 0 {
		log.Println("failed to GetContents err ", err)
		return nil
	}
	fmt.Printf("contents: %v \n", contents)
	return mvc.View{
		Name: "user/balancelist.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "balance",
			"Data":    contents,
		},
		Layout: "shared/indexlayout.html",
	}
}
