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
var cachedContentList = make(map[string]*models.Content)
var cachedContentLock = sync.Mutex{}

// ContentService interface methods
type ContentService interface {
	Get(id int) *models.Content
	Update(data *models.Content, columns []string) error
	Create(data *models.Content) error
	GetByContentHash(hash string) *models.Content
	// InnerConTentHash(address string) sql.Result
}

type contentService struct {
	dao *dao.ContentDao
}

// NewContentService ContentService entance
func NewContentService() ContentService {
	return &contentService{
		dao: dao.NewContentDao(datasource.InstanceDbMaster()),
	}
}

func (s *contentService) Get(id int) *models.Content {
	return s.dao.Get(id)
}
func (s *contentService) Update(data *models.Content, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *contentService) Create(data *models.Content) error {
	return s.dao.Create(data)
}

func (s *contentService) GetByContentHash(hash string) *models.Content {
	return s.dao.GetByContentHash(hash)
}
