package models

type Votecount struct {
	Address  string `xorm:"not null default '' comment('ether地址') VARCHAR(256)"`
	Amount   int    `xorm:"not null default 0 comment('投票票数') INT(20)"`
	TokenId  int    `xorm:"not null default 0 comment('资产_id') INT(10)"`
	VoteId   int    `xorm:"not null pk autoincr INT(11)"`
	VoteTime int    `xorm:"not null default 0 comment('投票时间') INT(20)"`
}
