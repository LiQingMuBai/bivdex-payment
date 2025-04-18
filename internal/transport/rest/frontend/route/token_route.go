package route

import (
	"github.com/1stpay/1stpay/internal/config"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NewTokenRouter(env *config.Env, db *gorm.DB, group *gin.RouterGroup, deps *config.Dependencies) {
	rates := group.Group("/token")
	c := deps.Controllers.MerchantTokenController
	{
		rates.GET("/list/", c.ListActive)
	}
}
