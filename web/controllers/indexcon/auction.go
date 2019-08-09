package indexcon

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"eilieili/comm"
	"eilieili/datasource"
	"eilieili/eths"
	"eilieili/models"
	"eilieili/web/utils"
	"eilieili/web/viewmodels"

	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

var (
	endBid = false
	mutex  sync.Mutex
)

// PostAuction ...
func (c *IndexController) PostAuction() error {
	// 0.5. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)
	// 1. 解析参数
	fmt.Printf("start parse from ......................\n")
	auction := viewmodels.Auction{}
	err := c.Ctx.ReadJSON(&auction)
	if auction.Percent <= 0 || auction.Price <= 0 {
		log.Println("failed to auction.PostAuction ReadJSON(auction) err ", err)
		return nil
	}

	// 2. 插入拍卖(auction)数据库
	ts := comm.NowUnix()
	address := c.GetAddress()
	newAuction := models.Auction{
		ContentHash: auction.ContentHash,
		Address:     address,
		TokenId:     int(auction.TokenID),
		Percent:     int(auction.Percent),
		Price:       int(auction.Price),
		Status:      1,
		Ts:          ts,
	}
	// fmt.Println("auction: ", newAuction)
	// TODO: 需要将资产列表的状态更新为1
	err = c.ServiceAuction.Create(&newAuction)
	if err != nil {
		log.Println("failed to Create(newAuction) err: ", err)
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
			<-ticker.C
			// c.EndBid(int(auction.TokenID), int(auction.Percent), resp)
		}
	}()
	return nil
	// return mvc.View{
	// 	Name: "user/autionlist.html",
	// 	Data: iris.Map{
	// 		"Channel": "autcion",
	// 		"Data":    "",
	// 	},

	// 	Layout: "shared/indexlayout.html",
	// }
}

// EndBid ...
func (c *IndexController) EndBid(tokenID, weight int, resp utils.Resp) error {
	// TODO: 数据库读取过多, Bidwinner/auction 数据的操作, 可以直接从缓存中读取
	// 1 根据tokenId, 查询出最高价者的address, price
	winDetail := c.ServiceBidwinner.GetByTokenId(tokenID)
	price := winDetail.Price
	fmt.Println("price: ", price)
	// address := winDetail.Address
	// 2. 数据库操作, price
	// 2.1 获取拍卖时的价格
	auctPrice := c.ServiceAuction.GetByStatus(tokenID, 1)
	_priceAuct := auctPrice.Price
	fmt.Println("_priceAuct: ", _priceAuct)
	auctcont := models.Auction{TokenId: tokenID, Status: 0}
	err := c.ServiceAuction.Update(&auctcont, []string{"status"})
	if err != nil {
		log.Println("failed to Create(&bid) err: ", err)
		return nil
	}
	// 3.1 更新content数据库: percent(总的-参与拍卖的)/price()
	dao := utils.NewContentinfoService(datasource.InstanceDbMaster())
	s := make(map[string]int)
	s["token_id"] = tokenID
	auctions, num, err := dao.InnerAuction(s)
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
			err = eths.EthSplitAsset(conf.Config.Eth.Fundation, conf.Config.Eth.FundationPWD, address, tokenID, weight)
			if err != nil {
				resp.Errno = utils.RECODE_MINTERR
				fmt.Println("failed to eths.EthSplitAsset ", err)
				return
			}

			err = eths.EthErc20Transfer(address, conf.Config.Eth.FundationPWD, to, price)
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
func (c *IndexController) GetAuctions() mvc.Result {
	// 1. 查看拍卖
	// 自动识别出查询字段所在tables
	// sql := fmt.Sprintf("select a.*,b.title from auction a, content b where a.content_hash=b.content_hash and a.status=1;")
	dao := utils.NewContentinfoService(datasource.InstanceDbMaster())
	s := make(map[string]int)
	s["status"] = 1
	log.Println("get auction start ...")
	auctions, num, err := dao.InnerAuction(s)
	if err != err || num <= 0 {
		log.Println("failed to GetContents err ", err)
		return nil
	}
	fmt.Printf("contents: %v \n", auctions)
	return mvc.View{
		Name: "user/auctionlist.html",
		Data: iris.Map{
			"Title":   "管理后台",
			"Channel": "balance",
			"Data":    auctions,
		},
		Layout: "shared/indexlayout.html",
	}
}

// GetBid ...
func (c *IndexController) GetBid() error {
	// 1. 响应数据结构初始化
	var resp utils.Resp
	resp.Errno = utils.RECODE_OK
	defer utils.ResponseData(c.Ctx, &resp)

	// 2. 获取参数
	// price,tokenid
	price := c.Ctx.URLParam("price")
	tokenID := c.Ctx.URLParam("tokenid")

	var err error
	if endBid {
		fmt.Printf("this bid is  end: %s\n", tokenID)
		resp.Errno = utils.RECODE_DBERR
		return err
	}

	// 进行比较竞拍值, 保存该address
	_price, _ := strconv.ParseInt(price, 10, 32)
	_tokenid, _ := strconv.ParseInt(tokenID, 10, 32)

	// 3.5 获取该address响应erc20 余额, 保证其有足够(>=30)的token进行该次投票
	address := c.GetAddress()
	log.Println("get address: ", address)
	erc20Balance, _ := eths.GetPxcBalance(address)
	if erc20Balance < _price || err != nil {
		fmt.Printf("%s: your erc20 balance is poor, connot operate this vote\n", address)
		resp.Errno = utils.RECODE_ERC20POORERR
		return err
	}

	// 从bidwinner中取出当前最大的price, 然后进行比较
	value := c.ServiceBidwinner.GetByTokenId(int(_tokenid))
	Price := value.Price
	id := value.Id

	// 同步锁, 防止多人同时修改数据
	mutex.Lock()
	defer mutex.Unlock()

	// fmt.Printf("the price: %d and Price: %d\n", _price, Price)
	if int(_price) > Price {

		log.Println("_price, Price: ", _price, Price)
		Price = int(_price)
		// TODO: xorm update-- 只能根据id进行更新, 无法定义其他所需的字段
		theWinner := models.Bidwinner{
			Id:      int(id), // 必须传入id, 否则无法准确定位到该数据
			TokenId: int(_tokenid),
			Price:   Price,
			Address: address,
			Ts:      comm.NowUnix(),
		}
		err = c.ServiceBidwinner.Update(&theWinner, []string{"price", "address", "ts"})
		if err != nil {
			log.Println("auction Update failed err: ", err)
			resp.Errno = utils.RECODE_DBERR
			return err
		}
		fmt.Printf("the account: %s Join bid success ...", address)
	} else {
		resp.Errno = utils.RECODE_DATAERR
		return err
	}
	return nil
}

// GetAddress ...
func (c *IndexController) GetAddress() string {
	user := comm.GetLoginUser(c.Ctx.Request())
	// username, passwd := c.getSession()
	// fmt.Println("the user: ", username)
	// TODO: 获得address, 直接从缓存中读取/cookies
	userobj, err := c.ServiceAccount.GetByUserName(user.Username)
	if err != nil {
		log.Println("user_index.PostContent GetByUserName err: ", err)
		return ""
	}
	return userobj.Address
}
