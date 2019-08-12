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
var cachedauctionList = make(map[string]*models.Auction)
var cachedauctionLock = sync.Mutex{}

// AuctionService interface methods
type AuctionService interface {
	Get(id int) *models.Auction
	GetAll(page, size int) []models.Auction
	Update(data *models.Auction, columns []string) error
	Delete(id int) error
	Create(data *models.Auction) error
	GetByStatus(id, status int) *models.Auction
	GetByContentHash(hash string) *models.Auction
	GetAllTokenId() []models.Auction
}

type auctionService struct {
	dao *dao.AuctionDao
}

// NewAuctionService auctionService entance
func NewAuctionService() AuctionService {
	return &auctionService{
		dao: dao.NewAuctionDao(datasource.InstanceDbMaster()),
	}
}

func (s *auctionService) Get(id int) *models.Auction {
	return s.dao.Get(id)
}

func (s *auctionService) GetAll(page, size int) []models.Auction {
	return s.dao.GetAll(page, size)
}

func (s *auctionService) Update(data *models.Auction, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *auctionService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *auctionService) Create(data *models.Auction) error {
	return s.dao.Create(data)
}
func (s *auctionService) GetByStatus(id, status int) *models.Auction {
	return s.dao.GetByStatus(id, status)
}
func (s *auctionService) GetByContentHash(hash string) *models.Auction {
	return s.dao.GetByContentHash(hash)
}

func (s *auctionService) GetAllTokenId() []models.Auction {
	return s.dao.GetAllTokenId()
}
