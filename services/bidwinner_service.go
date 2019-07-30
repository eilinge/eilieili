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
var cachedbidwinnerList = make(map[string]*models.Bidwinner)
var cachedbidwinnerLock = sync.Mutex{}

// BidwinnerService interface methods
type BidwinnerService interface {
	Get(id int) *models.Bidwinner
	Update(data *models.Bidwinner, columns []string) error
	Create(data *models.Bidwinner) error
	GetByTokenId(id int) *models.Bidwinner
}

type bidwinnerService struct {
	dao *dao.BidwinnerDao
}

// NewBidwinnerService bidwinnerService entance
func NewBidwinnerService() BidwinnerService {
	return &bidwinnerService{
		dao: dao.NewBidwinnerDao(datasource.InstanceDbMaster()),
	}
}

func (s *bidwinnerService) Get(id int) *models.Bidwinner {
	return s.dao.Get(id)
}

func (s *bidwinnerService) Update(data *models.Bidwinner, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *bidwinnerService) Create(data *models.Bidwinner) error {
	return s.dao.Create(data)
}

func (s *bidwinnerService) GetByTokenId(id int) *models.Bidwinner {
	return s.dao.GetByTokenId(id)
}
