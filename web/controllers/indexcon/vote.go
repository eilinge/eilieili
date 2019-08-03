package indexcon

import (
	"fmt"
	"log"
	"math"
	"strconv"

	"eilieili/comm"
	"eilieili/conf"
	"eilieili/datasource"
	"eilieili/eths"
	"eilieili/models"
	"eilieili/web/utils"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

// VotePlace ...
func (c *IndexController) GetVoteplace() mvc.Result {
	// 1. 查看所有资产
	dao := utils.NewContentinfoService(datasource.InstanceDbMaster())
	log.Println("get vote start ...")
	_, count, err := dao.InnerAuction(nil)
	countPage := 2
	// int64(3/2) 1
	// math.Ceil(1.2) 2
	// math.Floor(1.2) 1
	pageCount := math.Ceil(float64(count) / float64(countPage))
	pageIndex := 1

	// <a href="/admin?id=1">first</a>

	pageIndex, _ = c.Ctx.URLParamInt("id")

	countStart := countPage * (pageIndex - 1)
	// limitSQL := fmt.Sprintf("select token_id, title, a.content_hash from account_content a,content b where a.content_hash = b.content_hash limit %d, %d", startCount, stopCount)
	// datalist := c.Service.GetLimit(countPage, countStart)

	limit := make(map[string]int)
	start := strconv.Itoa(countPage)
	// limit, start
	log.Println("countPage, start: ", countPage, countStart)
	limit[start] = countStart
	datalist, _, err := dao.InnerAuction(limit)
	if err != nil {
		log.Println("failed to vote.InnerAuction(limit) err ", err)
		return nil
	}
	firstPage := false
	if pageIndex == 1 {
		firstPage = true
	}

	lastPage := false
	if pageIndex == int(pageCount) {
		lastPage = true
	}

	return mvc.View{
		Name: "user/vote.html", // view template
		Data: iris.Map{
			"Title":     "管理后台",
			"Datalist":  datalist,
			"count":     count,
			"pageCount": pageCount,
			"current":   pageIndex,
			"firstPage": firstPage,
			"lastPage":  lastPage,
		},
		Layout: "shared/indexlayout.html",
	}
}

// Vote ...
func (c *IndexController) GetVote() error {
	// 1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)

	// 2. 获取参数
	obj := c.Ctx.URLParams()
	tokenID, _ := strconv.ParseInt(obj["token_id"], 10, 32)
	log.Println("tokenid: ", tokenID)
	// 3. address
	address := c.GetAddress()
	// 3.5 获取该address响应erc20 余额, 保证其有足够(>=30)的token进行该次投票
	log.Println("the address: ", address)
	erc20Balance, err := eths.GetPxcBalance(address)
	if erc20Balance < 30 || err != nil {
		fmt.Printf("%s: your erc20 balance is poor, connot operate this vote\n", address)
		resp.Errno = utils.RECODE_ERC20POORERR
		return err
	}
	// 4. 存储到数据库
	voteObj := models.Vote{
		Address:  address,
		TokenId:  int(tokenID),
		VoteTime: comm.NowUnix(),
	}
	err = c.ServiceVote.Create(&voteObj)
	if err != nil {
		fmt.Println("failed to VoteSQL")
		resp.Errno = utils.RECODE_DATAERR
		return err
	}
	// 5. 操作以太坊, 进行投票, 只能合约内将erc20 token转给tokenID的地址
	eths.VoteTo(address, conf.Config.Eth.FundationPWD, tokenID)
	return nil
}
