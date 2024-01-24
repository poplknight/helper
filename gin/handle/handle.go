package handle

import (
	"github.com/gin-gonic/gin"
	"github.com/poplknight/helper/gin/binding"
	"reflect"
)

type Null struct{}

func MakeGinHandler[Req, Resp any](fc HandlerFunc[Req, Resp]) func(ctx *gin.Context) {
	var r Req
	var a any = r
	if reflect.TypeOf(a).Kind() == reflect.Ptr {
		binding.Register(r)
	} else {
		binding.Register(new(Req))
	}

	return func(ctx *gin.Context) {
		req, err := binding.GetRequest[Req](ctx)
		if err != nil {
			BadRequest(ctx, err)
			return
		}

		resp, err := fc(ctx, req)
		if err != nil {
			InternalServerError(ctx, err)
			return
		}

		Success(ctx, resp)
	}
}
