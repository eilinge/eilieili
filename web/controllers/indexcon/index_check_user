package indexcon

import (
	"time"

	"eilieili/models"
	"eilieili/services"
)

func (api *LuckyApi) checkBlackUser(uid int) (bool, *models.LtUser) {
	info := services.NewUserService().Get(uid)
	if info != nil && info.Blacktime > int(time.Now().Unix()) {
		return false, info
	}

	return true, info
}
