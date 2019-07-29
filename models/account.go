package models

type Account struct {
	AccountId  int    `xorm:"not null pk autoincr INT(10)"`
	Email      string `xorm:"not null default '' comment('email地址') unique VARCHAR(50)"`
	Username   string `xorm:"not null default '' comment('用户名') unique VARCHAR(30)"`
	IdentityId string `xorm:"not null default '' comment('验证身份id') VARCHAR(100)"`
	Address    string `xorm:"not null default '' comment('ether地址') VARCHAR(256)"`
}
