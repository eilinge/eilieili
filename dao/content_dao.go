package dao

import (
	"eilieili/models"

	"github.com/go-xorm/xorm"
)

type ContentDao struct {
	engine *xorm.Engine
}

func NewContentDao(engine *xorm.Engine) *ContentDao {
	return &ContentDao{
		engine: engine,
	}
}

func (d *ContentDao) Get(id int) *models.Content {
	data := &models.Content{ContentId: id}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.ContentId = 0
	return data
}

func (d *ContentDao) Update(data *models.Content, columns []string) error {
	_, err := d.engine.Id(data.ContentId).MustCols(columns...).Update(data)
	return err
}

func (d *ContentDao) Create(data *models.Content) error {
	n, err := d.engine.Insert(data)
	if n > 0 && err == nil {
		return nil
	}
	return err
}

func (d *ContentDao) GetByContentHash(hash string) *models.Content {
	data := &models.Content{ContentHash: hash}
	ok, err := d.engine.Get(data)
	if ok && err == nil {
		return data
	}
	data.ContentHash = ""
	return data
}
