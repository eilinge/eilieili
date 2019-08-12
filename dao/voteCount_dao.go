package dao

import (
	"eilieili/models"
	"log"

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
	data := &models.Votecount{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

func (d *VotecountDao) Update(data *models.Votecount, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
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
	data := &models.Votecount{}
	ok, err := d.engine.Where("token_id=?", id).Get(data)
	if ok && err == nil {
		log.Println("GetByTokenid data: ", data)
		return data
	}
	data.TokenId = 0
	return data
}

// GetAll get all token id of asset
func (d *VotecountDao) GetAll() []models.Votecount {
	datalist := []models.Votecount{}
	// err := d.engine.Cols("token_id", "amount").Find(&datalist)
	err := d.engine.Cols("token_id").Find(&datalist)
	if err != nil {
		return nil
	}
	return datalist
}
