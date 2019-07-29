package indexcon

import (
	"time"

	"lottery/models"
	"lottery/services"
)

func (api *LuckyApi) checkBlackip(ip string) (bool, *models.LtBlackip) {
	blackipInfo := services.NewBlackipService()
	info := blackipInfo.GetByIp(ip)
	if info == nil || info.Ip == "" {
		return true, nil
	}
	if info.Blacktime > int(time.Now().Unix()) {
		return false, info
	}

	return true, info
}
