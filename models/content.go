package models

type Content struct {
	ContentId   int    `xorm:"not null pk autoincr INT(11)"`
	Title       string `xorm:"not null default '' comment('名称') VARCHAR(100)"`
	Content     string `xorm:"not null default '' comment('资产') VARCHAR(256)"`
	ContentHash string `xorm:"not null default '' comment('资产hash') VARCHAR(100)"`
	Price       int    `xorm:"not null default 0 comment('价格') INT(100)"`
	Weight      int    `xorm:"not null default 0 comment('权重') INT(100)"`
	Ts          int    `xorm:"not null default 0 comment('修改时间') INT(20)"`
}
