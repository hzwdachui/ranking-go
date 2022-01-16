package user

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewUser(t *testing.T) {
	type args struct {
		id      int
		voteCnt int
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		{
			name: "case1 正常创建",
			args: args{
				id:      0,
				voteCnt: 1,
			},
			want: &user{
				id:        0,
				voteCnt:   1,
				voteLimit: 10,
			},
		},
	}
	for _, tt := range tests {
		convey.Convey("test NewUser", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := NewUser(tt.args.id, tt.args.voteCnt)
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}

func TestNewUserMap(t *testing.T) {
	tests := []struct {
		name string
		want UserMap
	}{
		{
			name: "case1 正常创建",
			want: &userMap{
				userMap: map[int]User{},
			},
		},
	}
	for _, tt := range tests {
		convey.Convey("test NewUserMap", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := NewUserMap()
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_user_Id(t *testing.T) {
	tests := []struct {
		name string
		u    *user
		want int
	}{
		{
			name: "case1 正常获取",
			u: &user{
				id: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Id", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.u.Id()
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_user_VoteLimit(t *testing.T) {
	tests := []struct {
		name string
		u    *user
		want int
	}{
		{
			name: "case1 正常获取",
			u: &user{
				id:        1,
				voteLimit: 10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Id", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.u.VoteLimit()
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_user_VoteCnt(t *testing.T) {
	tests := []struct {
		name string
		u    *user
		want int
	}{
		{
			name: "case1 正常增加",
			u: &user{
				id:        1,
				voteLimit: 10,
				voteCnt:   3,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Id", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.u.VoteCnt()
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_user_Incr(t *testing.T) {
	tests := []struct {
		name string
		u    *user
		want int
	}{
		{
			name: "case1 正常获取",
			u: &user{
				id:        1,
				voteLimit: 10,
				voteCnt:   3,
			},
			want: 4,
		},
		{
			name: "case2 超过限制",
			u: &user{
				id:        1,
				voteLimit: 10,
				voteCnt:   10,
			},
			want: 10,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Id", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				tt.u.Incr()
				convey.So(tt.u.voteCnt, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_userMap_Get(t *testing.T) {
	type args struct {
		uid int
	}
	tests := []struct {
		name  string
		um    func() UserMap
		args  args
		want  User
		want1 bool
	}{
		{
			name: "case1 获取正常",
			um: func() UserMap {
				um := &userMap{
					userMap: map[int]User{},
				}
				um.userMap[1] = NewUser(1, 0)
				return um
			},
			args: args{
				uid: 1,
			},
			want:  NewUser(1, 0),
			want1: true,
		},
		{
			name: "case1 获取失败",
			um: func() UserMap {
				um := &userMap{
					userMap: map[int]User{},
				}
				return um
			},
			args: args{
				uid: 1,
			},
			want:  nil,
			want1: false,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Get", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got, got1 := tt.um().Get(tt.args.uid)
				convey.So(got, convey.ShouldResemble, tt.want)
				convey.So(got1, convey.ShouldEqual, tt.want1)
			})
		})
	}
}

func Test_userMap_Set(t *testing.T) {
	type args struct {
		uid int
		u   User
	}
	tests := []struct {
		name string
		um   *userMap
		args args
		want User
	}{
		{
			name: "case1 正常设置",
			um: &userMap{
				userMap: map[int]User{},
			},
			args: args{
				uid: 1,
				u:   &user{},
			},
			want: &user{},
		},
	}
	for _, tt := range tests {
		convey.Convey("test Set", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				tt.um.Set(tt.args.uid, tt.args.u)
				convey.So(tt.um.userMap[tt.args.uid], convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_userMap_UserMap(t *testing.T) {
	tests := []struct {
		name string
		um   *userMap
		want map[int]User
	}{
		{
			name: "case1 正常获取",
			um: &userMap{
				userMap: map[int]User{},
			},
			want: map[int]User{},
		},
	}
	for _, tt := range tests {
		convey.Convey("test UserMap", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.um.UserMap()
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_userMap_SetNX(t *testing.T) {
	type args struct {
		uid int
		u   User
	}
	tests := []struct {
		name string
		um   *userMap
		args args
		want User
	}{
		{
			name: "case1 已经存在",
			um: &userMap{
				userMap: map[int]User{
					1: &user{},
				},
			},
			args: args{
				uid: 1,
				u:   &user{id: 2},
			},
			want: &user{},
		},
		{
			name: "case1 不存在",
			um: &userMap{
				userMap: map[int]User{},
			},
			args: args{
				uid: 1,
				u:   &user{id: 2},
			},
			want: &user{id: 2},
		},
	}
	for _, tt := range tests {
		convey.Convey("test SetNX", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.um.SetNX(tt.args.uid, tt.args.u)
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}
