package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/zzpu/mynahs/internal/controller"
	"github.com/zzpu/mynahs/internal/model"
	"github.com/zzpu/mynahs/internal/pkg/errors"
	log "github.com/zzpu/mynahs/internal/pkg/logger"
	"github.com/zzpu/mynahs/internal/service"
	"go.uber.org/zap"
	"strconv"
)

type GoodsController struct {
	controller.Api
}

func NewGoodsController() *GoodsController {
	return &GoodsController{}
}

// HelloWorld hello world
func (api *GoodsController) Check(c *gin.Context) {
	str, ok := c.GetQuery("barcode")
	if !ok {
		api.Success(c, "无效商品码")
		return
	}
	goodsSrv := service.NewGoodsService()

	goods, err := goodsSrv.Check(str)
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c, goods)
}
func (api *GoodsController) UpdatePrice(c *gin.Context) {
	str, ok := c.GetQuery("barcode")
	if !ok {
		api.Fail(c, errors.InvalidParameter, "无效商品码")
		return
	}
	price, ok := c.GetQuery("price")

	goodsSrv := service.NewGoodsService()

	err := goodsSrv.UpdatePrice(str, price)
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c, "成功")
}

func (api *GoodsController) Update(c *gin.Context) {
	var data model.GoodBean
	err := c.BindJSON(&data)
	if err != nil {
		api.Fail(c, errors.InvalidParameter, "无效参数")
		return
	}
	goodsSrv := service.NewGoodsService()
	log.Logger.Info("position:%d,name:%s", zap.Any("Position:", data.Position), zap.Any("GoodsName:", data.GoodsName))
	goods, err := goodsSrv.Update(&data)
	if err != nil {
		api.Err(c, err)
		return
	}
	log.Logger.Info("position:%d,name:%s", zap.Any("Position:", goods.Position), zap.Any("spec:", goods.Spec))
	api.Success(c, goods)
}

func (api *GoodsController) List(c *gin.Context) {
	var p model.Page
	err := c.BindJSON(&p)
	if err != nil || p.PageSize == 0 {
		p.Page, _ = strconv.Atoi(c.Param("page"))
		p.PageSize, _ = strconv.Atoi(c.Param("pageSize"))
	}
	log.Logger.Info("page:%d,pageSize:%d", zap.Any("page:", p.Page), zap.Any("pageSize:", p.PageSize))
	if p.GetOffset() < 0 {
		p.Page = 0
	}

	if p.GetLimit() <= 0 {
		p.PageSize = 10
	}
	goodsSrv := service.NewGoodsService()
	goods, err := goodsSrv.GoodsList(&p)
	if err != nil {
		api.Err(c, err)
		return
	}

	api.Success(c, goods)
}
