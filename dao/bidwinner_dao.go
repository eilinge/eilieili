package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type BidwinnerDao struct {
	engine *xorm.Engine
}

func NewBidwinnerDao(engine *xorm.Engine) *BidwinnerDao {
	return &BidwinnerDao{
		engine: engine,
	}
}
func (d *BidwinnerDao) Get(id int) *models.Bidwinner {
	data := &models.Bidwinner{Id: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.Id = 0
	return data
}

func (d *BidwinnerDao) Update(data *models.Bidwinner, columns []string) error {
	_, err := d.engine.Id(data.Id).MustCols(columns...).Update(data)
	return err
}

func (d *BidwinnerDao) Create(data *models.Bidwinner) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *BidwinnerDao) GetByTokenId(id int) *models.Bidwinner {
	data := &models.Bidwinner{TokenId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.TokenId = 0
	return data
}
