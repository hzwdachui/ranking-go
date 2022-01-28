package user

const (
	VOTE_LIMIT = 10
)

// NewUser
func NewUser(id int, voteCnt int) User {
	return &user{
		id:        id,
		voteCnt:   voteCnt,
		voteLimit: VOTE_LIMIT,
	}
}

// NewUserMap
func NewUserMap() UserMap {
	return &userMap{
		userMap: map[int]User{},
	}
}

// Id
func (u *user) Id() int {
	return u.id
}

// VoteCnt
func (u *user) VoteCnt() int {
	return u.voteCnt
}

// VoteLimit
func (u *user) VoteLimit() int {
	return u.voteLimit
}

// Incr 用户投票数+1操作，原子性
func (u *user) Incr() bool {
	u.mu.Lock()
	defer u.mu.Unlock()
	if u.voteCnt >= u.voteLimit {
		return false
	}

	u.voteCnt += 1
	return true
}

// Get 通过用户id获取用户
func (um *userMap) Get(uid int) (User, bool) {
	um.mu.Lock()
	defer um.mu.Unlock()
	u, has := um.userMap[uid]
	return u, has
}

// Set 设置用户id和用户的映射关系
func (um *userMap) Set(uid int, u User) {
	um.mu.Lock()
	defer um.mu.Unlock()
	um.userMap[uid] = u
}

// UserMap
func (um *userMap) UserMap() map[int]User {
	return um.userMap
}

// SetNX 存在用户直接返回用户，不存在则设置用户，并且返回新用户
func (um *userMap) SetNX(uid int, u User) User {
	um.mu.Lock()
	defer um.mu.Unlock()
	uOld, has := um.userMap[uid]
	if !has {
		um.userMap[uid] = u
		return u
	}
	return uOld
}
