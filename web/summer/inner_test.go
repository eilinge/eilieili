package summer

import (
	"eilieili/datasource"
	"fmt"
	"log"

	"github.com/go-xorm/xorm"
)

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
	User `xorm:"extends"` // {Id:0, Name:"eilinge", GroupId:0}
	Id   int64            // group.Id(默认)
	Name string           // group.string(默认)
}

func (UserGroup) TableName() string {
	return "user" //
}

type ContentinfoDao struct {
	engine *xorm.Engine
}

func NewContentinfoService(engine *xorm.Engine) *ContentinfoDao {
	return &ContentinfoDao{
		engine: engine,
	}
}

/*
[]main.UserGroup{main.UserGroup{User:main.User{Id:0, Name:"eilinge", GroupId:0}, Name:"A"}, main.UserGroup{User:main.User{Id:0, Name:"lin", GroupId:0}, Name:"B"}}
[]main.UserGroup{main.UserGroup{User:main.User{Id:0, Name:"", GroupId:0}, Name:"eilinge"}, main.UserGroup{User:main.User{Id:0, Name:"", GroupId:0}, Name:"lin"}}
*/

func (c *ContentinfoDao) InnerContent() {
	var users []UserGroup
	err := c.engine.Cols("user.name", "`group.name`").Join("INNER", "`group`", "`group`.id = user.groupid").Find(&users)
	// err := c.engine.Join("INNER", "`group`", "`group`.id = user.groupid").Find(&users)
	if err != nil {
		log.Println("failed to join ....", err)
		return
	}
	for _, v := range users {
		fmt.Printf("the users: %s in group %#v \n", v.User.Name, v)
	}

}
func main() {
	// var engine datasource.InstanceDbMaster()
	dao := NewContentinfoService(datasource.InstanceDbMaster())
	dao.InnerContent()
}
