package utils

import (
	"eilieili/models"
	"log"

	"github.com/go-xorm/xorm"
)

type ContentInfo struct {
	models.Content        `xorm:"extends"`
	models.AccountContent `xorm:"extends"`
	models.Auction        `xorm:"extends"`
}

type ResContentInfo struct {
	ContentHash string
	Weight      int
	Title       string
	TokenID     int
}

type ResAuctionInfo struct {
	Percent     int
	Weight      int
	Price       int
	ContentHash string
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
func (c *ContentinfoDao) InnerAuction(id int) (*map[int]ResAuctionInfo, int, error) {
	var contentinfo []ContentInfo
	err := c.engine.Join("INNER", "auction", "auction.content_hash = content.content_hash").
		Where("token_id=?", id).Find(&contentinfo)
	if err != nil {
		log.Println("failed to join ....", err)
		return nil, 0, err
	}
	log.Println("contentinfo: ", contentinfo)
	res := make(map[int]ResAuctionInfo)
	for i, info := range contentinfo {
		res[i] = ResAuctionInfo{
			ContentHash: info.Content.ContentHash,
			Weight:      info.Content.Weight,
			Percent:     info.Percent,
			Price:       info.Content.Price,
		}
	}
	log.Println("res: ", res)
	num := len(res)
	return &res, num, nil
}

// func main() {
// 	dao := NewContentinfoService(datasource.InstanceDbMaster())
// 	dao.InnerContent()
// }
