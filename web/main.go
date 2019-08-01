package indexcon

import (
	"eilieili/datasource"
	"log"

	"github.com/go-xorm/xorm"
)

// content_hash,weight,a.title,b.token_id content a, account_content
// a.content_hash = b.content_hash and address
type ContentInfo struct {
	ContentHash string `xorm:"extends"`
	Weight      int    `xorm:"extends"`
	Title       string `xorm:"extends"`
	TokenID     int    `xorm:"extends"`
}

func (ContentInfo) TableName() (string, int, string, int) {
	return "content_hash", 0, "title", 0
}

var contentinfo []ContentInfo

type Group struct {
	Id   int64
	Name string
}
type User struct {
	Id      int64
	Name    string
	GroupId int64 `xorm:"index"`
}

type UserGroup struct {
	User `xorm:"extends"`
	Name string
}

func (UserGroup) TableName() string {
	return "user"
}

type ContentinfoDao struct {
	engine *xorm.Engine
}

func NewContentinfoService(engine *xorm.Engine) *ContentinfoDao {
	return &ContentinfoDao{
		engine: engine,
	}
}

func (c *ContentinfoDao) InnerContent() {
	var users []UserGroup
	err := c.engine.Join("INNER", "group", "group.id = user.group_id").Find(&users)
	if err != nil {
		log.Println("failed to join ....")
		return
	}
}
func main() {
	// var engine datasource.InstanceDbMaster()
	dao := NewContentinfoService(datasource.InstanceDbMaster())
	dao.InnerContent()
}
