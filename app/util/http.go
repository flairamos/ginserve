package util

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Success   bool   `json:"success"`
	Message   string `json:"message"`
	Code      int    `json:"code"`
	Result    any    `json:"result"`
	TimeStamp int64  `json:"timestamp"`
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, Resp{
		Success:   true,
		Message:   "操作成功！",
		Code:      200,
		Result:    data,
		TimeStamp: time.Now().Unix(),
	})
}

func Fail(c *gin.Context, data interface{}) {
	c.JSON(http.StatusInternalServerError, Resp{
		Success:   false,
		Message:   "操作失败！",
		Code:      500,
		Result:    data,
		TimeStamp: time.Now().Unix(),
	})
}
