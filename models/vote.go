package models

import (
	"time"
)

type Vote struct {
	VoteId   int       `xorm:"not null pk autoincr INT(11)"`
	Address  string    `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	TokenId  int       `xorm:"not null default 0 comment('资产_id') unique INT(10)"`
	VoteTime time.Time `xorm:"not null default '2019-08-16 00:00:00' comment('当前时间') TIMESTAMP"`
	Comment  string    `xorm:"not null default '' comment('备注') VARCHAR(100)"`
}
