package response

import (
	"github.com/gin-gonic/gin"
	"github.com/zzpu/mynahs/config"
	"github.com/zzpu/mynahs/internal/pkg/errors"
	"net/http"
	"time"
)

type result struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Cost    string      `json:"cost"`
}

type Response struct {
	httpCode int
	result   *result
}

func Resp() *Response {
	// 初始化response
	return &Response{
		httpCode: http.StatusOK,
		result: &result{
			Code:    0,
			Message: "",
			Data:    nil,
			Cost:    "",
		},
	}
}

// Fail 错误返回
func (r *Response) Fail(c *gin.Context, code int, msg string, data ...interface{}) {
	r.SetCode(code)
	r.SetMessage(msg)
	if data != nil {
		r.WithData(data[0])
	}
	r.json(c)
}

// FailCode 自定义错误码返回
func (r *Response) FailCode(c *gin.Context, code int, msg ...string) {
	r.SetCode(code)
	if msg != nil {
		r.SetMessage(msg[0])
	}
	r.json(c)
}

// Success 正确返回
func (r *Response) Success(c *gin.Context) {
	r.SetCode(errors.SUCCESS)
	r.json(c)
}

// WithDataSuccess 成功后需要返回值
func (r *Response) WithDataSuccess(c *gin.Context, data interface{}) {
	r.SetCode(errors.SUCCESS)
	r.WithData(data)
	r.json(c)
}

// SetCode 设置返回code码
func (r *Response) SetCode(code int) *Response {
	r.result.Code = code
	return r
}

// SetHttpCode 设置http状态码
func (r *Response) SetHttpCode(code int) *Response {
	r.httpCode = code
	return r
}

type defaultRes struct {
	Result interface{} `json:"result"`
}

// WithData 设置返回data数据
func (r *Response) WithData(data interface{}) *Response {
	switch data.(type) {
	case string, int, bool:
		r.result.Data = &defaultRes{Result: data}
	default:
		r.result.Data = data
	}
	return r
}

// SetMessage 设置返回自定义错误消息
func (r *Response) SetMessage(message string) *Response {
	r.result.Message = message
	return r
}

var ErrorText = errors.NewErrorText(config.Config.Language)

// json 返回 gin 框架的 HandlerFunc
func (r *Response) json(c *gin.Context) {
	if r.result.Message == "" {
		r.result.Message = ErrorText.Text(r.result.Code)
	}

	// if r.Data == nil {
	// 	r.Data = struct{}{}
	// }

	r.result.Cost = time.Since(c.GetTime("requestStartTime")).String()
	c.AbortWithStatusJSON(r.httpCode, r.result)
}

// Success 业务成功响应
func Success(c *gin.Context, data ...interface{}) {
	if data != nil {
		Resp().WithDataSuccess(c, data[0])
		return
	}
	Resp().Success(c)
}

// FailCode 业务失败响应
func FailCode(c *gin.Context, code int, data ...interface{}) {
	if data != nil {
		Resp().WithData(data[0]).FailCode(c, code)
		return
	}
	Resp().FailCode(c, code)
}

// Fail 业务失败响应
func Fail(c *gin.Context, code int, message string, data ...interface{}) {
	if data != nil {
		Resp().WithData(data[0]).FailCode(c, code, message)
		return
	}
	Resp().FailCode(c, code, message)
}
