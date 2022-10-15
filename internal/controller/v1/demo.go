package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/qauzy/mynahs/internal/controller"
)

type DemoController struct {
	controller.Api
}

func NewDemoController() *DemoController {
	return &DemoController{}
}

// HelloWorld hello world
func (api *DemoController) HelloWorld(c *gin.Context) {
	str, ok := c.GetQuery("name")
	if !ok {
		str = "mynahs"
	}

	api.Success(c, fmt.Sprintf("hello %s", str))
}
