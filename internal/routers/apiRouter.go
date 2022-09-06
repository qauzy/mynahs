package routers

import (
	"github.com/gin-gonic/gin"
	controllerV1 "github.com/zzpu/mynahs/internal/controller/v1"
	"github.com/zzpu/mynahs/internal/controller/v1/ws"
)

func setApiRoute(r *gin.Engine) {
	go ws.WebsocketManager.Start()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendService()
	go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendGroupService()
	go ws.WebsocketManager.SendAllService()
	go ws.WebsocketManager.SendAllService()
	//go ws.TestSendGroup()
	//go ws.TestSendAll()

	// version 1
	v1 := r.Group("/api/v1")
	{
		auth := controllerV1.NewAuthController()
		v1.POST("/login", auth.Login)
		demo := controllerV1.NewDemoController()
		v1.GET("/hello-world", demo.HelloWorld)

		wsGroup := v1.Group("/ws")
		{
			wsGroup.GET("/:channel", ws.WebsocketManager.WsClient)
			wsGroup.GET("/notify", ws.NewNotifyController().Notify)
		}
	}

}
