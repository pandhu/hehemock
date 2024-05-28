package routes

import (
	"github.com/gin-gonic/gin"
	usecaseprovider "github.com/pandhu/hehemock/app/providers/usecase"
	"github.com/pandhu/hehemock/config"
	v1Routes "github.com/pandhu/hehemock/routes/api/v1"
)

// InitRouter initialize router
func InitRouter(uc *usecaseprovider.Usecase) *gin.Engine {
	engine := gin.New()

	engine.Use(gin.LoggerWithWriter(gin.DefaultWriter))

	engine.Use(gin.Recovery())

	engine.MaxMultipartMemory = 8 << 20 // 8 MiB memory for multipart forms

	router := engine.Group(config.All().App.ApiPrefix)

	v1 := router.Group("api/v1")
	{
		generalV1Route := v1Routes.NewGeneralRoute(v1)
		generalV1Route.Routes(uc)

	}
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not Found"})
	})

	return engine
}
