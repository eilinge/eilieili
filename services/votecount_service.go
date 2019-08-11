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
var cachedvotecountList = make(map[string]*models.Votecount)
var cachedvotecountLock = sync.Mutex{}

// VotecountService interface methods
type VotecountService interface {
	Get(id int) *models.Votecount
	Update(data *models.Votecount, columns []string) error
	Create(data *models.Votecount) error
	GetByTokenid(id int) *models.Votecount
	GetAll() []models.Votecount
}

type votecountService struct {
	dao *dao.VotecountDao
}

// NewvotecountService votecountService entance
func NewvotecountService() VotecountService {
	return &votecountService{
		dao: dao.NewVotecountDao(datasource.InstanceDbMaster()),
	}
}

func (s *votecountService) Get(id int) *models.Votecount {
	return s.dao.Get(id)
}

func (s *votecountService) Update(data *models.Votecount, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *votecountService) Create(data *models.Votecount) error {
	return s.dao.Create(data)
}

func (s *votecountService) GetByTokenid(id int) *models.Votecount {
	return s.dao.GetByTokenid(id)
}

func (s *votecountService) GetAll() []models.Votecount {
	return s.dao.GetAll()
}
