package request

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/pkg/errors"
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
	return r
}

func ParseParams[P BaseParameter, H BaseHeader](context *gin.Context) (*Request[P, H], error) {
	req := NewRequest[P, H](context)
	err := req.bindHeaders()
	if err != nil {
		return nil, errors.Wrap(err, "request.ParseParams.bindHeaders failed")
	}
	err = req.bindParams()
	if err != nil {
		return nil, errors.Wrap(err, "request.ParseParams.bindParams failed")
	}
	return req, nil
}

func (r *Request[P, H]) bindParams() error {
	params := new(P)

	// 绑定body 中json参数
	method := strings.ToLower(r.ctx.Request.Method)
	if method == "post" {
		err := r.ctx.ShouldBindBodyWith(params, binding.JSON)
		if err != nil {
			return errors.Wrap(err, "bind json params failed")
		}
	}

	// 绑定url中query参数，会覆盖body中解出来的参数
	err := r.ctx.ShouldBindQuery(params)
	if err != nil {
		return errors.Wrap(err, "bind query params failed")
	}

	r.params = params
	return nil
}

func (r *Request[P, H]) bindHeaders() error {
	headers := new(H)
	err := r.ctx.ShouldBindHeader(headers)
	if err != nil {
		return errors.Wrap(err, "bind request headers failed")
	}
	r.headers = headers
	return nil
}

func (r *Request[P, H]) GetParams() *P {
	return r.params
}

func (r *Request[T, H]) GetHeaders() *H {
	return r.headers
}
