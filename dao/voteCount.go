package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type VotecountDao struct {
	engine *xorm.Engine
}

func NewVotecountDao(engine *xorm.Engine) *VotecountDao {
	return &VotecountDao{
		engine: engine,
	}
}

func (d *VotecountDao) Get(id int) *models.Votecount {
	data := &models.Votecount{VoteId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.VoteId = 0
	return data
}

func (d *VotecountDao) Update(data *models.Votecount, columns []string) error {
	_, err := d.engine.Id(data.VoteId).MustCols(columns...).Update(data)
	return err
}

func (d *VotecountDao) Create(data *models.Votecount) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *VotecountDao) GetByTokenid(id int) *models.Votecount {
	data := &models.Votecount{TokenId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.TokenId = 0
	return data
}

func (v *VotecountDao) GetAll() []models.Votecount {
	datalist := []models.Votecount{}
	err := v.engine.Distinct("token_id").Find(&datalist)
	if err != nil {
		return nil
	}
	return datalist
}
