package models

import (
	"time"
)

type AccountContent struct {
	ContentHash string    `xorm:"not null default '' comment('资产hash') VARCHAR(100)"`
	TokenId     int       `xorm:"not null pk default 0 comment('资产_id') unique INT(10)"`
	Address     string    `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	Ts          time.Time `xorm:"not null default '2019-08-16 00:00:00' comment('当前时间') TIMESTAMP"`
}
