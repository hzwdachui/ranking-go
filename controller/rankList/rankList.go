package ranklist

import (
	"net/http"

	serviceranklist "example.com/rankingSystem/service/rankList"
)

func Register() {
	http.HandleFunc("/vote", serviceranklist.VoteHandler)
	http.HandleFunc("/getRankList", serviceranklist.GetRankListHandler)
}
