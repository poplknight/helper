package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poplknight/helper/gin/handle"
	"github.com/poplknight/helper/hash"
)

type HelloReq struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type HelloResp struct {
	UserId string `json:"user_id"`
}

func main() {
	engine := gin.New()
	// curl调用: curl http://127.0.0.1:8080/user -X POST -d '{"username":"poplknight","password":"123456"}'
	// 这样做的好处是不需要处理handler层，只需要关注service层即可
	engine.POST("/login", handle.MakeGinHandler(ServiceHello))
	fmt.Println(engine.Run(":8080"))
}

// ServiceHello service 层代码
func ServiceHello(_ context.Context, req HelloReq) (HelloResp, error) {
	return HelloResp{UserId: hash.Md5[string, string](req.Username + req.Password)}, nil
}
