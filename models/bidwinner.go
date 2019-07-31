package models

type Bidwinner struct {
	Id      int    `xorm:"not null pk autoincr INT(11)"`
	TokenId int    `xorm:"not null default 0 comment('资产_id') unique INT(10)"`
	Price   int    `xorm:"not null default 0 comment('价格') INT(10)"`
	Address string `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	Ts      int    `xorm:"not null default 0 comment('修改时间') INT(20)"`
}
