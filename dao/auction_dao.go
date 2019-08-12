package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type AuctionDao struct {
	engine *xorm.Engine
}

func NewAuctionDao(engine *xorm.Engine) *AuctionDao {
	return &AuctionDao{
		engine: engine,
	}
}
func (d *AuctionDao) Get(id int) *models.Auction {
	data := &models.Auction{TokenId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.TokenId = 0
	return data
}

func (d *AuctionDao) GetAll(page, size int) []models.Auction {
	offset := (page - 1) * size
	datalist := make([]models.Auction, 0)
	err := d.engine.
		Desc("token_id"). // ?
		Limit(size, offset).
		Find(&datalist)
	if err != nil {
		return datalist
	}
	return datalist
}

func (d *AuctionDao) Update(data *models.Auction, columns []string) error {
	_, err := d.engine.Id(data.TokenId).MustCols(columns...).Update(data)
	return err
}

func (d *AuctionDao) Delete(id int) error {
	data := &models.Auction{TokenId: id, Status: 1}
	_, err := d.engine.Id(data.TokenId).Update(&data)
	return err
}

func (d *AuctionDao) Create(data *models.Auction) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *AuctionDao) GetByStatus(id, status int) *models.Auction {
	data := &models.Auction{TokenId: id, Status: status}
	_, err := d.engine.Get(data)
	if err != nil {
		return nil
	}
	return data
}

func (d *AuctionDao) GetByContentHash(hash string) *models.Auction {
	data := &models.Auction{ContentHash: hash}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.ContentHash = ""
	return data
}
