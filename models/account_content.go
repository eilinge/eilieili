package models

type AccountContent struct {
	ContentHash string `xorm:"not null default '' comment('资产hash') VARCHAR(100)"`
	TokenId     int    `xorm:"not null pk default 0 comment('资产_id') unique INT(10)"`
	Address     string `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	Ts          int    `xorm:"not null default 0 comment('修改时间') INT(20)"`
}
