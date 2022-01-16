package user

import "sync"

type user struct {
	mu        sync.Mutex
	id        int // 用户 id
	voteCnt   int // 用户投票数
	voteLimit int // 投票限制数
}

type userMap struct {
	mu      sync.Mutex
	userMap map[int]User // 用户ID与用户的映射
}

type User interface {
	Id() int
	VoteCnt() int
	VoteLimit() int

	Incr() bool
}

type UserMap interface {
	Get(int) (User, bool)
	Set(int, User)

	UserMap() map[int]User
	SetNX(uid int, u User) User
}
