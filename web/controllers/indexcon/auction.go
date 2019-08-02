package indexcon

import (
	"fmt"
	"log"
	"time"

	"eilieili/comm"
	"eilieili/datasource"
	"eilieili/dbs"
	"eilieili/models"
	"eilieili/web/utils"
	"eilieili/web/viewmodels"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var (
	endBid = false
)

// PostAuction ...
func (c *IndexController) PostAuction() mvc.Result {
	// 0.5. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)
	// 1. 解析参数
	fmt.Printf("start parse from ......................\n")
	auction := viewmodels.Auction{}
	err := c.Ctx.ReadJSON(&auction)
	if auction.Percent <= 0 || auction.Price <= 0 {
		log.Println("failed to ReadJSON(&auction) err ", err)
		return nil
	}

	// 2. 插入拍卖(auction)数据库
	ts := comm.NowUnix()
	user := comm.GetLoginUser(c.Ctx.Request())
	// username, passwd := c.getSession()
	// fmt.Println("the user: ", username)
	// TODO: 获得address, 直接从缓存中读取/cookies
	userobj, _ := c.ServiceAccount.GetByUserName(user.Username)
	if err != nil {
		log.Println("user_index.PostContent GetByUserName err: ", err)
		return nil
	}
	address := userobj.Address
	newAuction := models.Auction{
		ContentHash: auction.ContentHash,
		Address:     address,
		TokenId:     int(auction.TokenID),
		Percent:     int(auction.Percent),
		Price:       int(auction.Price),
		Status:      1,
		Ts:          ts,
	}
	fmt.Println("auction: ", newAuction)
	err = c.ServiceAuction.Create(&newAuction)
	if err != nil {
		log.Println("failed to Create(&newAuction) err: ", err)
		return nil
	}
	fmt.Println("start insert into bidWinner...")
	// 3 插入bidwinner数据库
	bid := models.Bidwinner{
		TokenId: int(auction.TokenID),
		Price:   int(auction.Price),
		Address: address,
		Ts:      comm.NowUnix(),
	}
	err = c.ServiceBidwinner.Create(&bid)
	if err != nil {
		log.Println("failed to Create(&bid) err: ", err)
		return nil
	}

	if err != nil {

		return nil
	}
	fmt.Println("--------------------------------------")
	fmt.Println("start bid, this asset 3 minute over")
	// 4. 开始拍卖执行后, 设置定时器, 3分钟后, 时间结束, 自动完成财产的分割/erc20转账
	ticker := time.NewTicker(time.Minute * 3)
	go func() {
		for i := 1; i > 0; i-- {
			// for {
			<-ticker.C
			// c.EndBid(int(auction.TokenID), int(auction.Percent), resp)
		}
	}()
	return mvc.View{
		Name: "user/autionlist.html",
		Data: iris.Map{
			"Channel": "autcion",
			"Data":    "",
		},

		Layout: "shared/indexlayout.html",
	}
}

// EndBid ...
func (c *IndexController) EndBid(tokenID, weight int, resp utils.Resp) error {
	// TODO: 数据库读取过多, Bidwinner/auction 数据的操作, 可以直接从缓存中读取
	// 3.5 根据tokenId, 查询出最高价者的address, price
	winDetail := c.ServiceBidwinner.GetByTokenId(tokenID)
	price := winDetail.Price
	fmt.Println("price: ", price)
	// address := winDetail.Address
	// 4. 数据库操作, price
	// 4.1 获取拍卖时的价格
	auctPrice := c.ServiceAuction.GetByStatus(tokenID, 1)
	_priceAuct := auctPrice.Price
	fmt.Println("_priceAuct: ", _priceAuct)
	auctcont := models.Auction{TokenId: tokenID, Status: 0}
	err := c.ServiceAuction.Update(&auctcont, []string{"status"})
	if err != nil {
		log.Println("failed to Create(&bid) err: ", err)
		return nil
	}
	// 4.1 更新content数据库: percent(总的-参与拍卖的)/price()
	dao := utils.NewContentinfoService(datasource.InstanceDbMaster())
	auctions, num, err := dao.InnerAuction(tokenID)
	if err != err || num <= 0 {
		log.Println("failed to GetContents err ", err)
		return nil
	}
	fmt.Printf("contents: %v \n", auctions)
	/*
		// aPercent, _ := strconv.ParseInt(auctionWeight[0]["percent"], 10, 32)
		bWeight, _ := strconv.ParseInt(auctionWeight[0]["weight"], 10, 32)
		bPrice, _ := strconv.ParseInt(auctionWeight[0]["price"], 10, 32)
		contentHash := auctionWeight[0]["content_hash"]
		_newPriceAuct, _ := strconv.ParseInt(_priceAuct, 10, 32)
		// 总的权重  - 拍卖时的权重
		newWeight := bWeight - weight
		// 拍卖之后的价格 - 拍卖时的价格 + 原本的价格
		newPrice := (price - _newPriceAuct) + bPrice
		// content_hash 不唯一(分割出的新资产与旧资产content_hash)
		UpConSQL := fmt.Sprintf("update content set price='%d' ,weight='%d' where content_hash ='%s';", newPrice, newWeight, contentHash)
		fmt.Println(UpconSQL)
		// 获取token_id最高竞拍者的price
		bidSQL := fmt.Sprintf("select b.price,b.address from auction a, bidwinner b where b.token_id='%d' and a.status=1;", tokenID)
		fmt.Println(bidSQL)
		to := value[0]["address"]

		fmt.Println("---------------------------------------")
		fmt.Println("aleady EndBid, Waiting SpiltAsset and transfer .....")
		// 5. 操作以太坊: 资产分割, erc20转账
		go func() {
			err = eths.EthSplitAsset(configs.Config.Eth.Fundation, configs.Config.Eth.FundationPWD, address, tokenID, weight)
			if err != nil {
				resp.Errno = utils.RECODE_MINTERR
				fmt.Println("failed to eths.EthSplitAsset ", err)
				return
			}

			err = eths.EthErc20Transfer(address, configs.Config.Eth.FundationPWD, to, price)
			if err != nil {
				resp.Errno = utils.RECODE_ERC20ERR
				fmt.Println("failed to eths.EthErc20Transfer ", err)
				return
			}
			fmt.Println("---------------------------------------")
			fmt.Println("Success SpiltAsset and transfer .....")
			endBid = true
		}()

	*/
	return nil
}

// GetAuctions ...
func (c *IndexController) GetAuctions() error {
	// 3. 查看拍卖
	// 自动识别出查询字段所在tables
	// sql := fmt.Sprintf("select a.*,b.title from auction a, content b where a.content_hash=b.content_hash and a.status=1;")
	sql := fmt.Sprintf("select a.content_hash,title,b.price,b.percent,token_id from content a, auction b where a.content_hash=b.content_hash and b.status=1")
	fmt.Println(sql)
	values, num, err := dbs.DBQuery(sql)
	if err != nil || num <= 0 {
		return err
	}
	mapResp := make(map[string]interface{})
	fmt.Printf("the values: %#v\n", values)
	mapResp["data"] = values

	return nil
	//1. 获取所有资产
	dao := utils.NewContentinfoService(datasource.InstanceDbMaster())
	contents, num, err := dao.InnerContent("0x127abc67e63ceb4dfeb3e066b9ee4297c12a8100")
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
