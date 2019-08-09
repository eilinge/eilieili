package utils

import (
	"fmt"
	"log"
	"math"

	"eilieili/datasource"
)

// incrTokenID 放入缓存中, 进行累加
func incrTokenID(id int64) int64 {
	key := fmt.Sprintf("day_ips_%d", id)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, ip, 1)
	if err != nil {
		log.Println("ip_day_lucky redis HINCRBY error=", err)
		return math.MaxInt32
	}
	return rs.(int64)
}

// InitTokenID ...
func InitTokenID(uid int, num int64) {
	if num <= 1 {
		return
	}
	i := uid % userFrameSize
	key := fmt.Sprintf("day_users_%d", i)
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, uid, num)
	if err != nil {
		log.Println("user_day.InitUserLuckNum redis HSET key=", key, ",error=", err)
		return
	}
}
