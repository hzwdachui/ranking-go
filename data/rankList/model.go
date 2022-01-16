// 基于 SortedSet 数据结构实现排行榜
package ranklist

import (
	"sync"

	"github.com/hdt3213/godis/datastruct/sortedset"
)

type ranklist struct {
	mu  sync.Mutex
	id  int                 // 排行榜 id
	set sortedset.SortedSet // 存储排行榜的数据接口
}

type ranklistMap struct {
	mu          sync.Mutex
	ranklistMap map[int]RankList // 排行榜id和排行榜的映射
}

type RankList interface {
	Id() int

	Range(int, int, bool) []*sortedset.Element
	Vote(int, int)
	Len() int64
	Get(starId int) int
}

type RankListMap interface {
	Get(int) (RankList, bool)
	Set(int, RankList)
	RanklistMap() map[int]RankList
	SetNX(rankListId int, rl RankList) RankList
}
