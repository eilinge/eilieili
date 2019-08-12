package utils

import (
	"fmt"
	"log"
	"math"
	"time"

	"eilieili/comm"
	"eilieili/datasource"
	"eilieili/services"

	"github.com/gomodule/redigo/redis"
)

const VoteFrameSize = 2

func init() {
	resetGroupVoteList()
}

/*
资产投票次数, token_id列替换amout列, hash中的计数器, utils/Vote_day

incrVoteCountyNum 原子性递增资产数
initVoteCountyNum 以数据库数据为准, 从数据库初始化缓存数据(单个用户抽奖次数需要很准确)

resetGroupVoteList 投票结束时, 计数器归零
*/

// InitvoteCountRedis ...
func InitvoteCountRedis() {
	votelist := services.NewvotecountService().GetAll()
	for _, vote := range votelist {
		InitVoteCountNum(vote.TokenId, int64(vote.Amount))
	}
}

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
func IncrVoteCountNum(tokenid int) int64 {
	key := fmt.Sprintf("votecount_%d", tokenid)
	cacheObj := datasource.InstanceCache()
	rs, err := cacheObj.Do("HINCRBY", key, tokenid, 1)
	if err != nil {
		log.Println("tokenid_count redis HINCRBY error=", err)
		return math.MaxInt32
	}
	return rs.(int64)
}

// InitVoteCountNum init vote count of token id
func InitVoteCountNum(tokenid int, num int64) {
	key := fmt.Sprintf("votecount_%d", tokenid)
	cacheObj := datasource.InstanceCache()
	_, err := cacheObj.Do("HSET", key, tokenid, num)
	if err != nil {
		log.Println("Vote_day.InitVoteCountNum redis HSET key=", key, ",error=", err)
		return
	}
}

// GetVoteCountNum Get vote count of token id
func GetVoteCountNum(tokenid int64) int64 {
	key := fmt.Sprintf("votecount_%d", tokenid)
	cacheObj := datasource.InstanceCache()
	res, err := redis.Int64(cacheObj.Do("HGET", key, tokenid))
	log.Println("GetVoteCountNum res: ", res)
	if err != nil {
		log.Println("Vote_day.GetVoteCountNum redis HSET key=", key, ",error=", err)
		return -1
	}
	return res
}

// Existkey ...
func Existkey(tokenid int64) bool {
	key := fmt.Sprintf("votecount_%d", tokenid)
	cacheObj := datasource.InstanceCache()
	res, err := redis.Bool(cacheObj.Do("EXISTS", key))
	if err != nil {
		log.Println("Vote_day.GetVoteCountNum redis HSET key=", key, ",error=", err)
		return false
	}
	return res
}
