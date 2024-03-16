package request

import (
	"strings"

	"encoding/json"
	"github.com/Edward-Alphonse/saywo_pkg/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Request[T any] struct {
	ctx *gin.Context

	headers *Headers
	params  *T
}

func NewRequest[T any](context *gin.Context) *Request[T] {
	r := &Request[T]{
		ctx: context,
	}
	r.bindHeaders()
	r.bindParams()
	return r
}

func (r *Request[T]) bindParams() {
	params := new(T)

	// 绑定body 中json参数
	method := strings.ToLower(r.ctx.Request.Method)
	if method == "post" {
		err := r.ctx.ShouldBindBodyWith(params, binding.JSON)
		if err != nil {
			logs.Error("bind json params failed", map[string]any{
				"error": err.Error(),
			})
			return
		}
	}

	// 绑定url中query参数，会覆盖body中解出来的参数
	err := r.ctx.ShouldBindQuery(params)
	if err != nil {
		logs.Error("bind query params failed", map[string]any{
			"error": err.Error(),
		})
		return
	}

	r.params = params
	b, _ := json.Marshal(params)
	logs.Info("请求参数", map[string]any{
		"params": string(b),
	})
}

func (r *Request[T]) bindHeaders() {
	headers := &Headers{}
	err := r.ctx.ShouldBindHeader(headers)
	if err != nil {
		logs.Error("bind request headers failed", map[string]any{
			"error": err.Error(),
		})
		return
	}
	r.headers = headers
}

func (r *Request[T]) GetParams() *T {
	return r.params
}

func (r *Request[T]) GetHeaders() *Headers {
	return r.headers
}

func (r *Request[T]) GetVersion() string {
	return r.ctx.Param("version")
}
