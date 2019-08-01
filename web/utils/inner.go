package utils

import (
	"eilieili/models"
	"log"

	"github.com/go-xorm/xorm"
)

// content_hash,weight,a.title,b.token_id content a, account_content
// a.content_hash = b.content_hash and address
type ContentInfo struct {
	models.Content        `xorm:"extends"`
	models.AccountContent `xorm:"extends"`
}

type ResContentInfo struct {
	ContentHash string
	Weight      int
	Title       string
	TokenID     int
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
			Title:       info.Title,
			TokenID:     info.TokenId,
		}
		// res[i].Weight = info.Weight
		// res[i].Title = info.Title
		// res[i].TokenID = info.TokenId
	}
	// log.Println("res: ", res)
	num := len(res)
	return &res, num, nil
}

// func main() {
// 	// var engine datasource.InstanceDbMaster()
// 	dao := NewContentinfoService(datasource.InstanceDbMaster())
// 	dao.InnerContent()
// }
