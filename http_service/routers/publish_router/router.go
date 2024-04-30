package publishrouter

import "github.com/gin-gonic/gin"

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/publish")
	group.POST("/:topic", PublishHandler)
}
