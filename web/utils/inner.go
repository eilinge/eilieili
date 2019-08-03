package utils

import (
	"log"
	"strconv"

	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type ContentInfo struct {
	models.Content        `xorm:"extends"`
	models.AccountContent `xorm:"extends"`
	models.Auction        `xorm:"extends"`
	models.Vote           `xorm:"extends"`
}

type ResContentInfo struct {
	ContentHash string
	Weight      int
	Title       string
	TokenID     int
}

type ResAuctionInfo struct {
	TokenID     int
	Title       string
	Percent     int
	Weight      int
	Price       int
	ContentHash string
	Content     string
}

type ResAddressInfo struct {
	Address string
}

func (ContentInfo) TableName() string {
	return "content"
}

type ContentinfoDao struct {
	engine *xorm.Engine
}

func NewContentinfoService(engine *xorm.Engine) *ContentinfoDao {
	return &ContentinfoDao{
		engine: engine,
	}
}

// InnerContent content_hash,weight,a.title,b.token_id content a, account_content a.content_hash = b.content_hash and address
func (c *ContentinfoDao) InnerContent(address string) (*map[int]ResContentInfo, int, error) {
	var contentinfo []ContentInfo
	err := c.engine.Alias("a").Join("INNER", "account_content", "a.content_hash = account_content.content_hash").
		Where("address=?", address).Find(&contentinfo)
	if err != nil {
		log.Println("failed to join ....", err)
		return nil, 0, err
	}
	// log.Println("contentinfo: ", contentinfo)
	res := make(map[int]ResContentInfo)
	for i, info := range contentinfo {
		res[i] = ResContentInfo{
			ContentHash: info.Content.ContentHash,
			Weight:      info.Weight,
			Title:       info.Content.Content,
			TokenID:     info.AccountContent.TokenId,
		}
	}
	// log.Println("res: ", res)
	num := len(res)
	return &res, num, nil
}

// InnerAuction "select a.percent,b.weight,b.price,a.content_hash from auction a, content b where a.content_hash = b.content_hash and token_id ='%d'"
func (c *ContentinfoDao) InnerAuction(s map[string]int) (*map[int]ResAuctionInfo, int, error) {
	var contentinfo []ContentInfo
	var err error
	// status/token_id
	// start, offset
	if s != nil {
		for i, v := range s {
			start, err := strconv.ParseInt(i, 10, 32)
			// int
			if err == nil {
				// limitSQL := fmt.Sprintf("select token_id, title, a.content_hash from account_content a,content b where a.content_hash = b.content_hash limit %d, %d", startCount, stopCount)
				log.Println("v, int(start): ", v, int(start))
				err = c.engine.Join("INNER", "auction", "auction.content_hash = content.content_hash").Limit(int(start), v).Find(&contentinfo)
			} else {
				// string
				err = c.engine.Join("INNER", "auction", "auction.content_hash = content.content_hash").
					Where(i+"=?", v).Find(&contentinfo)
			}
		}
	} else {
		err = c.engine.Join("INNER", "auction", "auction.content_hash = content.content_hash").Find(&contentinfo)
	}

	if err != nil {
		log.Println("failed to join ....", err)
		return nil, 0, err
	}
	// log.Println("contentinfo: ", contentinfo)
	res := make(map[int]ResAuctionInfo)
	for i, info := range contentinfo {
		res[i] = ResAuctionInfo{
			TokenID:     info.Auction.TokenId,
			Title:       info.Content.Title,
			ContentHash: info.Content.ContentHash,
			Content:     info.Content.Content,
			Weight:      info.Content.Weight,
			Percent:     info.Percent,
			Price:       info.Content.Price,
		}
	}
	log.Println("res: ", res)
	num := len(res)
	return &res, num, nil
}

// InnerAddress select distinct a.address from auction a, vote b where a.token_id= b.token_id and b.token_id='%d'
func (c *ContentinfoDao) InnerAddress(id int) (string, int, error) {
	var contentinfo ContentInfo
	var err error
	ok, err := c.engine.Join("INNER", "auction", "auction.content_hash = content.content_hash").
		Where("tokenid=?", id).Get(&contentinfo)

	if !ok || err != nil {
		log.Println("failed to join ....", err)
		return "", 0, err
	}
	res := contentinfo.Vote.Address
	return res, 0, nil
}

// func main() {
// 	dao := NewContentinfoService(datasource.InstanceDbMaster())
// 	dao.InnerContent()
// }
