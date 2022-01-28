package ranklist

import (
	"strconv"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/hdt3213/godis/datastruct/sortedset"
	"github.com/smartystreets/goconvey/convey"
)

func TestNewRankList(t *testing.T) {
	type args struct {
		id int
	}
	tests := []struct {
		name string
		want *ranklist
		args args
	}{
		{
			name: "case1 正常创建",
			want: &ranklist{
				id:  1,
				set: *sortedset.Make(),
			},
			args: args{
				id: 1,
			},
		},
	}
	for _, tt := range tests {
		convey.Convey("test NewRankList", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := NewRankList(tt.args.id)
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}
func TestNewRankListMap(t *testing.T) {
	tests := []struct {
		name string
		want RankListMap
	}{
		{
			name: "case1 正常获取",
			want: &ranklistMap{
				ranklistMap: map[int]RankList{},
			},
		},
	}
	for _, tt := range tests {
		convey.Convey("test NewRankListMap", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := NewRankListMap()
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_ranklist_Id(t *testing.T) {
	tests := []struct {
		name string
		rl   *ranklist
		want int
	}{
		{
			name: "case1 正常获取",
			rl: &ranklist{
				id: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		convey.Convey("test RanklistId", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.rl.Id()
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_ranklist_Vote(t *testing.T) {
	type args struct {
		starId    int
		ticketNum int
	}
	tests := []struct {
		name string
		rl   func() RankList
		args args
		want int
	}{
		{
			name: "case1 明星被第一次投票",
			rl: func() RankList {
				return NewRankList(1)
			},
			args: args{
				starId:    1,
				ticketNum: 2,
			},
			want: 2,
		},
		{
			name: "case2 明星被第n次投票",
			rl: func() RankList {
				rl := &ranklist{
					set: *sortedset.Make(),
				}
				rl.set.Add("1", 2)
				return rl
			},
			args: args{
				starId:    1,
				ticketNum: 2,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Vote", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				rl := tt.rl()
				rl.Vote(tt.args.starId, tt.args.ticketNum)
				convey.So(rl.Get(tt.args.starId), convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_ranklist_Range(t *testing.T) {
	type args struct {
		start int
		end   int
		desc  bool
	}
	tests := []struct {
		name string
		rl   func() RankList
		args args
		want []*sortedset.Element
	}{
		{
			name: "case1 正常获取",
			rl: func() RankList {
				rl := &ranklist{
					set: *sortedset.Make(),
				}
				rl.set.Add("1", 2)
				rl.set.Add("1", 4)
				rl.set.Add("2", 3)
				rl.set.Add("3", 1)
				return rl
			},
			args: args{
				start: 0,
				end:   10,
				desc:  false,
			},
			want: []*sortedset.Element{
				{
					Member: "3",
					Score:  1,
				},

				{
					Member: "2",
					Score:  3,
				},
				{
					Member: "1",
					Score:  4,
				},
			},
		},
		{
			name: "case2 反向获取",
			rl: func() RankList {
				rl := &ranklist{
					set: *sortedset.Make(),
				}
				rl.set.Add("1", 2)
				rl.set.Add("2", 3)
				rl.set.Add("3", 1)
				return rl
			},
			args: args{
				start: 0,
				end:   2,
				desc:  true,
			},
			want: []*sortedset.Element{
				{
					Member: "2",
					Score:  3,
				},
				{
					Member: "1",
					Score:  2,
				},
			},
		},
		{
			name: "case3 使用接口错误",
			rl: func() RankList {
				rl := &ranklist{
					set: *sortedset.Make(),
				}
				rl.set.Add("1", 2)
				rl.set.Add("2", 3)
				rl.set.Add("3", 1)
				return rl
			},
			args: args{
				start: 2,
				end:   0,
				desc:  true,
			},
			want: []*sortedset.Element{},
		},
		{
			name: "case4 正常获取2",
			rl: func() RankList {
				rl := &ranklist{
					set: *sortedset.Make(),
				}
				rl.set.Add("1", 2)
				rl.set.Add("1", 4)
				rl.set.Add("2", 3)
				rl.set.Add("3", 1)
				return rl
			},
			args: args{
				start: 1,
				end:   10,
				desc:  false,
			},
			want: []*sortedset.Element{
				{
					Member: "2",
					Score:  3,
				},
				{
					Member: "1",
					Score:  4,
				},
			},
		},
	}
	for _, tt := range tests {
		convey.Convey("test Range", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.rl().Range(tt.args.start, tt.args.end, tt.args.desc)
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_ranklist_Len(t *testing.T) {
	tests := []struct {
		name string
		rl   func() RankList
		want int
	}{
		{
			name: "case1 正常获取",
			rl: func() RankList {
				rl := &ranklist{
					set: *sortedset.Make(),
				}
				rl.set.Add(strconv.FormatInt(int64(1), 10), 2)
				return rl
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Len", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.rl().Len()
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_ranklist_Get(t *testing.T) {
	type args struct {
		starId int
	}
	tests := []struct {
		name string
		rl   RankList
		args args
		want int
	}{
		{
			name: "case1 不存在此明星",
			rl:   NewRankList(1),
			args: args{
				starId: 1,
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Get", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				ranklist := tt.rl
				got := ranklist.Get(tt.args.starId)
				convey.So(got, convey.ShouldEqual, tt.want)
			})
		})
	}
}

func Test_ranklistMap_Get(t *testing.T) {
	type args struct {
		rankListId int
	}
	tests := []struct {
		name  string
		rlm   *ranklistMap
		args  args
		want  RankList
		want1 bool
	}{
		{
			name: "case1 不存在",
			rlm:  &ranklistMap{},
			args: args{
				rankListId: 1,
			},
			want:  nil,
			want1: false,
		},
		{
			name: "case1 存在",
			rlm: &ranklistMap{
				ranklistMap: map[int]RankList{
					1: &ranklist{},
				},
			},
			args: args{
				rankListId: 1,
			},
			want:  &ranklist{},
			want1: true,
		},
	}
	for _, tt := range tests {
		convey.Convey("test Get", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got, got1 := tt.rlm.Get(tt.args.rankListId)
				convey.So(got, convey.ShouldResemble, tt.want)
				convey.So(got1, convey.ShouldEqual, tt.want1)
			})
		})
	}
}

func Test_ranklistMap_Set(t *testing.T) {
	type args struct {
		rankListId int
		rl         RankList
	}
	tests := []struct {
		name string
		rlm  *ranklistMap
		args args
		want RankList
	}{
		{
			name: "",
			rlm: &ranklistMap{
				ranklistMap: map[int]RankList{},
			},
			args: args{
				rankListId: 1,
				rl:         &ranklist{id: 1},
			},
			want: &ranklist{id: 1},
		},
	}
	for _, tt := range tests {
		convey.Convey("test Set", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				tt.rlm.Set(tt.args.rankListId, tt.args.rl)
				convey.So(tt.rlm.ranklistMap[tt.args.rankListId], convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_ranklistMap_RanklistMap(t *testing.T) {
	tests := []struct {
		name string
		rlm  *ranklistMap
		want map[int]RankList
	}{
		{
			name: "case1 正常获取",
			rlm: &ranklistMap{
				ranklistMap: map[int]RankList{},
			},
			want: map[int]RankList{},
		},
	}
	for _, tt := range tests {
		convey.Convey("test RanklistMap", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.rlm.RanklistMap()
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}

func Test_ranklistMap_SetNX(t *testing.T) {
	type args struct {
		rankListId int
		rl         RankList
	}
	tests := []struct {
		name string
		rlm  *ranklistMap
		args args
		want RankList
	}{
		{
			name: "case1 key存在",
			rlm: &ranklistMap{
				ranklistMap: map[int]RankList{
					1: &ranklist{id: 2},
				},
			},
			args: args{
				rankListId: 1,
				rl:         &ranklist{id: 1},
			},
			want: &ranklist{id: 2},
		},
		{
			name: "case2 key不存在",
			rlm: &ranklistMap{
				ranklistMap: map[int]RankList{},
			},
			args: args{
				rankListId: 1,
				rl:         &ranklist{id: 1},
			},
			want: &ranklist{id: 1},
		},
	}
	for _, tt := range tests {
		convey.Convey("test SetNX", t, func() {
			convey.Convey(tt.name, func() {
				ctl := gomock.NewController(t)
				defer ctl.Finish()
				got := tt.rlm.SetNX(tt.args.rankListId, tt.args.rl)
				convey.So(got, convey.ShouldResemble, tt.want)
			})
		})
	}
}
