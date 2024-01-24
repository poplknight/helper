package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/poplknight/helper/gin/binding"
)

type Req struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func main() {
	engine := gin.New()
	// 注册请求结构体 也可以不注册
	// 如果不注册会在运行时自动注册
	binding.Register(new(Req))
	// 该函数的curl命令为：curl http://127.0.0.1:8080/user -X POST -d '{"username":"poplknight","password":"123456"}'
	// 将以上的curl复制到命令行执行即可看到输出，证明绑定是成功的
	engine.POST("/user", func(ctx *gin.Context) {
		req, err := binding.GetRequest[Req](ctx)
		if err != nil {
			ctx.JSON(400, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(200, req)
	})
	fmt.Println(engine.Run(":8080"))
}
