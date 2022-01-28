package serviceranklist

import (
	"encoding/json"
	"net/http"

	ranklist "example.com/rankingSystem/data/rankList"
	"example.com/rankingSystem/data/user"
	"example.com/rankingSystem/global"
	"example.com/rankingSystem/utils"
)

type voteParams struct {
	Uid    *int `form:"uid" binding:"required"`
	StarId *int `form:"starId" binding:"required"`
	RankId *int `form:"rankId" binding:"required"`
}

type voteResponse struct {
	Code int               `json:"code"`
	Msg  string            `json:"msg"`
	Data *voteResponseData `json:"data"`
}

type voteResponseData struct {
	VoteCnt int `json:"voteCnt,omitempty"`
}

// VoteHandler 处理用户投票的请求
func VoteHandler(w http.ResponseWriter, req *http.Request) {
	var err error
	// var resp voteResponse
	params, err := checkParamsVote(req)
	if err != nil {
		voteResponseErr(w, "params error")
		return
	}

	// 记录用户投票数
	u := global.USER_MAP.SetNX(*params.Uid, user.NewUser(*params.Uid, 0))
	if suc := u.Incr(); !suc {
		voteResponseErr(w, "vote too many times")
		return
	}

	// 获取排行榜
	rl := global.RL_MAP.SetNX(*params.RankId, ranklist.NewRankList(*params.RankId))
	rl.Vote(*params.StarId, 1)

	voteRepsonseSucc(w, u)
}

// checkParamsVote 参数校验
func checkParamsVote(req *http.Request) (*voteParams, error) {
	params := &voteParams{}
	err := utils.BindQuery(params, req)
	return params, err
}

// voteResponseErr 返回错误
func voteResponseErr(w http.ResponseWriter, errMsg string) {
	var resp voteResponse
	resp.Code = 400
	resp.Msg = errMsg
	resp.Data = &voteResponseData{}
	respByte, _ := json.Marshal(resp)
	w.Write(respByte)
}

// voteRepsonseSucc 返回成功
func voteRepsonseSucc(w http.ResponseWriter, u user.User) {
	var resp voteResponse
	resp.Code = 200
	resp.Msg = "vote successfully"
	resp.Data = &voteResponseData{
		VoteCnt: u.VoteCnt(),
	}
	respByte, _ := json.Marshal(resp)
	w.Write(respByte)
}
