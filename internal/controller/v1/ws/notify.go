package ws

import (
	"github.com/gin-gonic/gin"
	"github.com/zzpu/mynahs/internal/controller"
)

type NotifyController struct {
	controller.Api
}

func NewNotifyController() *NotifyController {
	return &NotifyController{}
}

// HelloWorld hello world
func (api *NotifyController) Notify(c *gin.Context) {
	str, ok := c.GetQuery("msg")
	if !ok {
		WebsocketManager.SendAll([]byte("收到非法消息"))
	}
	WebsocketManager.SendAll([]byte(str))
	api.Success(c, "success")
}
