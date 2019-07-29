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
var cachedaccountContentList = make(map[string]*models.AccountContent)
var cachedaccountContentLock = sync.Mutex{}

// accountContentService interface methods
type AccountContentService interface {
	Get(id int) *models.AccountContent
	Create(data *models.AccountContent) error
	GetByContentHash(hash string) *models.AccountContent
}

type accountContentService struct {
	dao *dao.accountContentDao
}

// NewaccountContentService accountContentService entance
func NewaccountContentService() accountContentService {
	return &accountContentService{
		dao: dao.NewaccountContentDao(datasource.InstanceDbMaster()),
	}
}

func (s *accountContentService) Get(id int) *models.AccountContent {
	return s.dao.Get(id)
}

func (s *accountContentService) Create(data *models.AccountContent) error {
	return s.dao.Create(data)
}

func (s *accountContentService) GetByContentHash(hash string) *models.AccountContent {
	return s.dao.GetByContentHash(hash)
}
