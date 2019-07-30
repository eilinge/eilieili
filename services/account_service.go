/**
 * 抽奖系统数据处理（包括数据库，也包括缓存等其他形式数据）
 */
package services

import (
	"sync"

	"eilieili/dao"
	"eilieili/datasource"
	"eilieili/models"
)

// IP信息，可以缓存(本地或者redis)，有更新的时候，再根据具体情况更新缓存
var cachedAccountList = make(map[string]*models.Account)
var cachedAccountLock = sync.Mutex{}

// AccountService interface methods
type AccountService interface {
	Get(id int) *models.Account
	Update(data *models.Account, columns []string) error
	Create(data *models.Account) error
	GetByEmail(email string) *models.Account
	GetByUserName(usr string) *models.Account
}

type accountService struct {
	dao *dao.AccountDao
}

// NewAccountService AccountService entance
func NewAccountService() AccountService {
	return &accountService{
		dao: dao.NewAccountDao(datasource.InstanceDbMaster()),
	}
}

func (s *accountService) Get(id int) *models.Account {
	return s.dao.Get(id)
}
func (s *accountService) Update(data *models.Account, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *accountService) Create(data *models.Account) error {
	return s.dao.Create(data)
}

func (s *accountService) GetByEmail(email string) *models.Account {
	return s.dao.GetByEmail(email)
}

func (s *accountService) GetByUserName(usr string) *models.Account {
	return s.dao.GetByUserName(usr)
}
