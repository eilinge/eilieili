package models

import (
	"time"
)

type Bidwinner struct {
	Id      int       `xorm:"not null pk autoincr INT(11)"`
	TokenId int       `xorm:"not null default 0 comment('资产_id') unique INT(10)"`
	Price   int       `xorm:"not null default 0 comment('价格') INT(10)"`
	Address string    `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	Ts      time.Time `xorm:"not null default '2019-08-16 00:00:00' comment('当前时间') TIMESTAMP"`
}
