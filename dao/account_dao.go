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

func (d *AccountDao) GetByEmail(email string) (*models.Account, error) {
	data := &models.Account{Email: email}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data, nil
	}
	data.Email = ""
	return data, err
}

func (d *AccountDao) GetByUserName(usr string) (*models.Account, error) {
	data := &models.Account{}
	ok, err := d.engine.Where("username=?", usr).Get(data)
	if ok && err == nil {
		return data, nil
	}
	data.Username = ""
	return data, err
}

func (d *AccountDao) GetByUserAddr(usr string) (*models.Account, error) {
	data := &models.Account{}
	ok, err := d.engine.Cols("address").Where("username=?", usr).Get(data)
	if ok && err == nil {
		return data, nil
	}
	data.Address = ""
	return data, err
}
