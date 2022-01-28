package serviceranklist

import (
	"encoding/json"
	"net/http"
	"strconv"

	ranklist "example.com/rankingSystem/data/rankList"
	"example.com/rankingSystem/data/user"
	"example.com/rankingSystem/global"
	"example.com/rankingSystem/utils"
	"github.com/hdt3213/godis/datastruct/sortedset"
)

type getRankListParams struct {
	RankId *int `form:"rankId" binding:"required"`
}

type getRankListResponse struct {
	Code int                     `json:"code"`
	Msg  string                  `json:"msg"`
	Data getRankListResponseData `json:"data"`
}

type getRankListResponseData struct {
	RankId   string         `json:"rankId"`
	RankList []*elementResp `json:"ranklist"`
	// RankList []*sortedset.Element `json:"ranklist"`
}

type elementResp struct {
	Member string  `json:"starId"`
	Score  float64 `json:"votes"`
}

// GetRankListHandler 处理获取排行榜的接口
func GetRankListHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	params, err := checkParamsGetRankList(req)
	if err != nil {
		getRankListResponseErr(w, "params error")
		return
	}
	rl, has := global.RL_MAP.Get(*params.RankId)
	if !has {
		getRankListResponseErr(w, "not such rank")
		return
	}
	// 获取排行榜并排序
	elements := rl.Range(0, user.VOTE_LIMIT, true)
	getRankListSucc(w, rl, elements)
}

// checkParamsGetRankList 参数校验
func checkParamsGetRankList(req *http.Request) (*getRankListParams, error) {
	params := &getRankListParams{}
	err := utils.BindQuery(params, req)
	return params, err
}

// getRankListResponseErr 返回失败
func getRankListResponseErr(w http.ResponseWriter, errMsg string) {
	var resp voteResponse
	resp.Code = 400
	resp.Msg = errMsg
	resp.Data = &voteResponseData{}
	respByte, _ := json.Marshal(resp)
	w.Write(respByte)
}

// getRankListSucc 返回成功
func getRankListSucc(w http.ResponseWriter, rl ranklist.RankList, elements []*sortedset.Element) {
	var resp getRankListResponse
	l := []*elementResp{}
	for _, ele := range elements {
		l = append(l, &elementResp{ele.Member, ele.Score})
	}
	resp.Code = 200
	resp.Data = getRankListResponseData{
		RankId:   strconv.Itoa(rl.Id()),
		RankList: l,
	}
	respJson, _ := json.Marshal(resp)
	w.Write(respJson)
}
