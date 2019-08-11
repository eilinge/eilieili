package utils

import (
	"fmt"
	"log"
	"math"
	"time"

	"eilieili/comm"
	"eilieili/datasource"
)

const VoteFrameSize = 2

func init() {
	resetGroupVoteList()
}

/*
用户今日抽奖次数, hash中的计数器, utils/Vote_day

incrVoteCountyNum 原子性递增用户今日的抽奖次数
initVoteCountyNum 以数据库数据为准, 从数据库初始化缓存数据(单个用户抽奖次数需要很准确)

resetGroupVoteList 每天凌晨计数器归零

VoteFrameSize 优化, 将hash结构散列为多段数据, 让每个hash小点, 以提高redis的执行效率
*/

func resetGroupVoteList() {
	log.Println("Vote_count.resetGroupVoteList start")
	cacheObj := datasource.InstanceCache()
	for i := 0; i < VoteFrameSize; i++ {
		key := fmt.Sprintf("votecount_%d", i)
		cacheObj.Do("DEL", key)
	}
	log.Println("Vote_count.resetGroupVoteList stop")
	// TODO: Votecount当天的统计数, 投票结束之后的时候归零, 设置定时器
	duration := comm.VoteEnd()
	time.AfterFunc(duration, resetGroupVoteList)
}

// IncrVoteCountNum token_id放入缓存中, votecount进行累加
func IncrVoteCountNum(uid int) int64 {
	i := uid % VoteFrameSize
	key := fmt.Sprintf("votecount_%d", i)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, uid, 1)
	if err != nil {
		log.Println("uid_count redis HINCRBY error=", err)
		return math.MaxInt32
	}
	return rs.(int64)
}

// InitVoteCountNum init vote count of token id
func InitVoteCountNum(uid int, num int64) {
	if num <= 1 {
		return
	}
	i := uid % VoteFrameSize
	key := fmt.Sprintf("votecount_%d", i)
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, uid, num)
	if err != nil {
		log.Println("Vote_day.InitVoteCountNum redis HSET key=", key, ",error=", err)
		return
	}
}
