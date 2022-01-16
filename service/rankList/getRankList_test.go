package serviceranklist

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/smartystreets/goconvey/convey"
)

func TestGetRankListHandler(t *testing.T) {
	type args struct {
		w   *httptest.ResponseRecorder
		req *http.Request
	}
	tests := []struct {
		name     string
		args     args
		wantBody string
	}{
		{
			name: "case1 参数错误",
			args: args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest("GET", "http://test.com/getRankList", nil),
			},
			wantBody: `{"code":400,"msg":"params error","data":{}}`,
		},
		{
			name: "case2",
			args: args{
				w:   httptest.NewRecorder(),
				req: httptest.NewRequest("GET", "http://test.com/getRankList?rankId=1", nil),
			},
			wantBody: `{"code":400,"msg":"not such rank","data":{}}`,
		},
	}
	for _, tt := range tests {
		convey.Convey("test handler", t, func() {
			GetRankListHandler(tt.args.w, tt.args.req)
			resp := tt.args.w.Result()
			body, _ := ioutil.ReadAll(resp.Body)
			convey.So(string(body), convey.ShouldEqual, tt.wantBody)
		})
	}
}
