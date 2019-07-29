package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type AccountContentDao struct {
	engine *xorm.Engine
}

func NewAccountContentDao(engine *xorm.Engine) *AccountContentDao {
	return &AccountContentDao{
		engine: engine,
	}
}
func (d *AccountContentDao) Get(id int) *models.AccountContent {
	data := &models.AccountContent{TokenId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.TokenId = 0
	return data
}

func (d *AccountContentDao) Create(data *models.AccountContent) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *AccountContentDao) GetByContentHash(hash string) *models.AccountContent {
	data := &models.AccountContent{ContentHash: hash}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.ContentHash = ""
	return data
}
