package indexcon

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
