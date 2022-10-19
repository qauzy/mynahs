package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qauzy/mynahs/internal/controller"
	"github.com/qauzy/mynahs/internal/dto"
	"github.com/qauzy/mynahs/internal/pkg/errors"
	"github.com/qauzy/mynahs/internal/pkg/logger"
	"github.com/qauzy/mynahs/internal/service"
)

type OrdersController struct {
	controller.Api
}

func NewOrdersController() *OrdersController {
	return &OrdersController{}
}

/*
*
 */
func (api *OrdersController) Create(c *gin.Context) {
	var req dto.OrdersDto
	err := c.BindJSON(&req)
	if err != nil {
		api.Fail(c, errors.InvalidParameter, "参数无法解析"+err.Error())
		return
	}
	if req.Detail == nil || len(req.Detail) == 0 {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	logger.Info(req)
	OrdersSrv := service.NewOrdersService()

	err = OrdersSrv.Create(&req)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c)
}

/*
*获取订单列表
 */
func (api *OrdersController) List(c *gin.Context) {
	var req dto.OrdersDto
	err := c.BindJSON(&req)
	if err != nil {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	logger.Info("OrdersController:", req)
	OrdersSrv := service.NewOrdersService()
	var data *service.OrdersBo
	data, err = OrdersSrv.List(&req)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c, data)
}

/*
*获取订单详情
 */
func (api *OrdersController) Detail(c *gin.Context) {
	var data dto.OrdersDto
	err := c.BindJSON(&data)
	if err != nil {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	if data.Detail == nil || len(data.Detail) == 0 {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	logger.Info(data)
	OrdersSrv := service.NewOrdersService()

	err = OrdersSrv.Create(&data)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c)
}

func (api *OrdersController) Delete(c *gin.Context) {
	var data dto.OrdersDto
	err := c.BindJSON(&data)
	if err != nil {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	if data.Detail == nil || len(data.Detail) == 0 {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	logger.Info(data)
	OrdersSrv := service.NewOrdersService()

	err = OrdersSrv.Create(&data)
	if err != nil {
		api.Err(c, err)
		return
	}
	api.Success(c)
}
