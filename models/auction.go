package models

type Auction struct {
	ContentHash string `xorm:"not null default '' comment('资产hash') VARCHAR(256)"`
	Address     string `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	TokenId     int    `xorm:"not null pk default 0 comment('资产_id') unique INT(10)"`
	Percent     int    `xorm:"not null default 0 comment('权重') INT(10)"`
	Price       int    `xorm:"not null default 0 comment('价格') INT(10)"`
	Status      int    `xorm:"not null default 0 comment('状态') INT(10)"`
	Ts          int    `xorm:"not null default 0 comment('修改时间') INT(20)"`
}
