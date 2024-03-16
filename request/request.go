package request

import (
	"strings"

	"encoding/json"
	"github.com/Edward-Alphonse/saywo_pkg/logs"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type BaseParameter Mapable
type BaseHeader Mapable

type Request[P BaseParameter, H BaseHeader] struct {
	ctx *gin.Context

	headers *H
	params  *P
}

func NewRequest[P BaseParameter, H BaseHeader](context *gin.Context) *Request[P, H] {
	r := &Request[P, H]{
		ctx: context,
	}
	r.bindHeaders()
	r.bindParams()
	return r
}

func (r *Request[P, H]) bindParams() {
	params := new(P)

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

func (r *Request[P, H]) bindHeaders() {
	headers := new(H)
	err := r.ctx.ShouldBindHeader(headers)
	if err != nil {
		logs.Error("bind request headers failed", map[string]any{
			"error": err.Error(),
		})
		return
	}
	r.headers = headers
}

func (r *Request[P, H]) GetParams() *P {
	return r.params
}

func (r *Request[T, H]) GetHeaders() *H {
	return r.headers
}
