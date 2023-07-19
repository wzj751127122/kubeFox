package middle

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseData struct {
	Code ResCode     `json:"code"`
	Msg  interface{} `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ResponseError(c *gin.Context, code ResCode) {

	rd := &ResponseData{
		Code: code,
		Msg:  code.Msg(),
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}

func ResponseErrorWithMsg(c *gin.Context, code ResCode,msg interface{}) {

	rd := &ResponseData{
		Code: code,
		Msg:  msg,
		Data: nil,
	}
	c.JSON(http.StatusOK, rd)

}



func ResponseSuccess(c *gin.Context, data interface{}) {

	rd := &ResponseData{
		Code: CodeSuccess,
		Msg:  CodeSuccess.Msg(),
		Data: data,
	}
	c.JSON(http.StatusOK, rd)

}
