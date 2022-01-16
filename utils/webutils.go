package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// parseGetRequest 将 Get 请求参数转换成结构体
func BindQuery(param interface{}, req *http.Request) error {
	c := gin.Context{
		Request: req,
	}
	return c.ShouldBindQuery(param)
}
