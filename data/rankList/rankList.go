package ranklist

import (
	"strconv"

	"github.com/hdt3213/godis/datastruct/sortedset"
)

// NewRankList
func NewRankList(id int) RankList {
	return &ranklist{
		id:  id,
		set: *sortedset.Make(),
	}
}

// NewRankListMap
func NewRankListMap() RankListMap {
	return &ranklistMap{
		ranklistMap: map[int]RankList{},
	}
}

// Id
func (rl *ranklist) Id() int {
	return rl.id
}

// Vote 给明星设置票数
func (rl *ranklist) Vote(starId int, ticketNum int) {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	strStarId := strconv.FormatInt(int64(starId), 10)

	node, has := rl.set.Get(strStarId)
	if !has {
		rl.set.Add(strStarId, float64(ticketNum))
	} else {
		newTicketNum := node.Score + float64(ticketNum)
		rl.set.Add(strStarId, newTicketNum)
	}
}

// Range 排行榜信息 取值范围 [start, end)
func (rl *ranklist) Range(start int, end int, desc bool) []*sortedset.Element {
	len := rl.set.Len()
	// 计算合法的上下区间
	if end < start {
		return make([]*sortedset.Element, 0)
	}

	if int64(end) >= len {
		return rl.set.Range(int64(start), len, desc)
	}

	return rl.set.Range(int64(start), int64(end), desc)
}

// Len 排行榜大小
func (rl *ranklist) Len() int64 {
	return rl.set.Len()
}

// Get 获取明星票数
func (rl *ranklist) Get(starId int) int {
	strStarId := strconv.FormatInt(int64(starId), 10)
	node, has := rl.set.Get(strStarId)
	if !has {
		return 0
	}
	return int(node.Score)
}

// Get 通过排行榜id获取排行榜
func (rlm *ranklistMap) Get(rankListId int) (RankList, bool) {
	rlm.mu.Lock()
	defer rlm.mu.Unlock()
	rl, has := rlm.ranklistMap[rankListId]
	return rl, has
}

// Set 设置排行榜id和排行榜的映射
func (rlm *ranklistMap) Set(rankListId int, rl RankList) {
	rlm.mu.Lock()
	defer rlm.mu.Unlock()
	rlm.ranklistMap[rankListId] = rl
}

// RanklistMap
func (rlm *ranklistMap) RanklistMap() map[int]RankList {
	return rlm.ranklistMap
}

// SetNX 存在排行榜直接返回排行榜，不存在则设置排行榜，并且返回新排行榜
func (rlm *ranklistMap) SetNX(rankListId int, rl RankList) RankList {
	rlm.mu.Lock()
	defer rlm.mu.Unlock()
	rlOld, has := rlm.ranklistMap[rankListId]
	if !has {
		rlm.ranklistMap[rankListId] = rl
		return rl
	}
	return rlOld
}
