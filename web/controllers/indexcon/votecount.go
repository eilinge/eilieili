package indexcon

import (
	"log"
	"time"

	"eilieili/models"
	"eilieili/services"
	"eilieili/web/utils"
)

// StoreVotecount 服务器重启, redis缓存数据全部清空
func StoreVotecount(tokenId int) {
	log.Println("token id: ", tokenId)
	if utils.Existkey(int64(tokenId)) {
		utils.IncrVoteCountNum(int(tokenId))
	} else {
		utils.InitVoteCountNum(tokenId, 1)
	}
}

// StoreVotecountSQL 先读取mysql数据库数据, 操作redis, 插入mysql, 保证数据不丢失, 但操作mysql数据太频繁, 操作时间长
func StoreVotecountSQL(tokenId int) {
	votecount := services.NewvotecountService()
	VotecountInfo := votecount.GetByTokenid(tokenId)
	if VotecountInfo != nil && VotecountInfo.TokenId == tokenId {
		// 该资产已经被投过票
		VotecountInfo.Amount++
		utils.InitVoteCountNum(tokenId, int64(VotecountInfo.Amount))
		err := votecount.Update(VotecountInfo, nil)
		if err != nil {
			log.Println("failed to votecount.Update err: ", err)
		}
		utils.IncrVoteCountNum(int(tokenId))
	} else {
		// 创建该资产的投票信息
		VotecountInfo = &models.Votecount{
			TokenId:  tokenId,
			Amount:   1,
			VoteTime: int(time.Now().Unix()),
		}
		err := votecount.Create(VotecountInfo)
		if err != nil {
			log.Println("failed to ServiceVotecount.Update err: ", err)
		}
		utils.InitVoteCountNum(tokenId, 1)
	}
}
