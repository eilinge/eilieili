package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type VoteDao struct {
	engine *xorm.Engine
}

func NewVoteDao(engine *xorm.Engine) *VoteDao {
	return &VoteDao{
		engine: engine,
	}
}

func (d *VoteDao) Get(id int) *models.Vote {
	data := &models.Vote{VoteId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.VoteId = 0
	return data
}

func (d *VoteDao) Update(data *models.Vote, columns []string) error {
	_, err := d.engine.Id(data.VoteId).MustCols(columns...).Update(data)
	return err
}

func (d *VoteDao) Create(data *models.Vote) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *VoteDao) GetByTokenid(id int) *models.Vote {
	data := &models.Vote{}
	ok, err := d.engine.Where("token_id=?", id).Get(data)
	if ok && err == nil {
		return data
	}
	data.TokenId = 0
	return data
}

func (d *VoteDao) GetAll() []models.Vote {
	datalist := []models.Vote{}
	err := d.engine.Distinct("token_id").Find(&datalist)
	if err != nil {
		return nil
	}
	// log.Println("data: ", datalist)
	return datalist
}
