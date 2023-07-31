package routes

import (
	"AlittleRequire/api"
	"github.com/gin-gonic/gin"
)

func NewRoute() *gin.Engine {
	ginRouter := gin.Default()

	ginRouter.GET("ping", func(c *gin.Context) {
		c.JSON(200, "success")
	})
	v1 := ginRouter.Group("/api/v1")
	{
		v1.POST("topic/create", api.CreateTopicHandle())
		v1.POST("post/create", api.CreatePostHandle())
		v1.GET("topic/info", api.GetTopicInfoHandle())
	}

	return ginRouter
}
