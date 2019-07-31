package models

type Vote struct {
	VoteId   int    `xorm:"not null pk autoincr INT(11)"`
	Address  string `xorm:"not null default '' comment('ether地址') VARCHAR(100)"`
	TokenId  int    `xorm:"not null default 0 comment('资产_id') unique INT(10)"`
	VoteTime int    `xorm:"not null default 0 comment('投票时间') INT(20)"`
	Comment  string `xorm:"not null default '' comment('备注') VARCHAR(100)"`
}
