package rest

import (
	"github.com/ValkyrieKia/golang-deals-test-project/internal/app/usecase"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/config"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/common"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/controller/dto"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/infrastructure/repository/item"
	"github.com/ValkyrieKia/golang-deals-test-project/internal/util"
	"github.com/gin-gonic/gin"
)

type itemController struct {
	rg     *gin.RouterGroup
	config *config.Config

	itemUc *usecase.ItemUsecase
}

func NewItemHttpController(
	rg *gin.RouterGroup,
	providers *common.HttpServerProviders,
) *itemController {
	requiredItemRepository := item.NewitemMysqlRepository(providers.MainDbConnection)

	ctrl := &itemController{
		rg:     rg,
		config: providers.Config,
		itemUc: usecase.NewItemUsecase(
			providers.Config,
			requiredItemRepository,
		),
	}

	rg.GET("", ctrl.GetAll)
	rg.POST("", ctrl.Create)

	return ctrl
}

func (ac *itemController) GetAll(c *gin.Context) {

	data, err := ac.itemUc.GetAll(c)
	if err != nil {
		_ = c.Error(err)
		return
	}
	c.JSON(200, data)
}

func (ac *itemController) Create(c *gin.Context) {
	var body *dto.ItemRequestDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		_ = c.Error(util.NewCommonError(err, util.ErrBadRequest, err.Error()))
		return
	}

	transformToEntity := body.ToEntity()

	err := ac.itemUc.Create(c, *transformToEntity)
	if err != nil {
		_ = c.Error(err)
		return
	}

	c.JSON(200, &gin.H{})
}
