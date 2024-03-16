package context

import "github.com/gin-gonic/gin/binding"

const (
	MIMEJSON = "application/json"
	MIMEPB   = "application/x-protobuf"
)

type HttpContext interface {
	Set(key string, value any)
	Param(key string) string
	ShouldBindJSON(obj any) error
	ShouldBindBodyWith(obj any, bb binding.BindingBody) (err error)
	ShouldBindQuery(obj any) error
	ShouldBindHeader(obj any) error

	Header(key, value string)
	JSON(code int, obj any)
	ProtoBuf(code int, obj any)
}
