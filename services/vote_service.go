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
var cachedvoteList = make(map[string]*models.Vote)
var cachedvoteLock = sync.Mutex{}

// VoteService interface methods
type VoteService interface {
	Get(id int) *models.Vote
	Update(data *models.Vote, columns []string) error
	Create(data *models.Vote) error
	GetByTokenid(id int) *models.Vote
	GetAll() []models.Vote
}

type voteService struct {
	dao *dao.VoteDao
}

// NewVoteService voteService entance
func NewVoteService() VoteService {
	return &voteService{
		dao: dao.NewVoteDao(datasource.InstanceDbMaster()),
	}
}

func (s *voteService) Get(id int) *models.Vote {
	return s.dao.Get(id)
}

func (s *voteService) Update(data *models.Vote, columns []string) error {
	return s.dao.Update(data, columns)
}

func (s *voteService) Create(data *models.Vote) error {
	return s.dao.Create(data)
}

func (s *voteService) GetByTokenid(id int) *models.Vote {
	return s.dao.GetByTokenid(id)
}

func (s *voteService) GetAll() []models.Vote {
	return s.dao.GetAll()
}
