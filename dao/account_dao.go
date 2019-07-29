package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type AccountDao struct {
	engine *xorm.Engine
}

func NewAccountDao(engine *xorm.Engine) *AccountDao {
	return &AccountDao{
		engine: engine,
	}
}
func (d *AccountDao) Get(id int) *models.Account {
	data := &models.Account{AccountId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.AccountId = 0
	return data
}

func (d *AccountDao) Update(data *models.Account, columns []string) error {
	_, err := d.engine.Id(data.AccountId).MustCols(columns...).Update(data)
	return err
}

func (d *AccountDao) Create(data *models.Account) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *AccountDao) GetByEmail(email string) *models.Account {
	data := &models.Account{Email: email}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Email = ""
	return data
}

func (d *AccountDao) GetByUserName(usr string) *models.Account {
	data := &models.Account{Username: usr}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Username = ""
	return data
}
