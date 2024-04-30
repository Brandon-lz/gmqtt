package routers

import (
	publishrouter "github.com/Brandon-lz/gmqtt/http_service/routers/publish_router"
	subscriberouter "github.com/Brandon-lz/gmqtt/http_service/routers/subscribe_router"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	router.GET("/", Root)
	router.GET("/ping", Ping)

	publishrouter.RegisterRoutes(router)
	subscriberouter.RegisterRoutes(router)
}

// 根路由
// @Summary  根路由
// @Description  根路由
// @Tags     default
// @Accept   json
// @Produce  json
// @Success  200  {object}  ApiResponse  "欢迎使用OPC-UA OpenAPI"
// @Router   /api/v1/ [get]
func Root(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "欢迎使用OPC-UA OpenAPI",
	})
}

type ApiResponse struct {
	Message string `json:"message" example:"欢迎使用OPC-UA OpenAPI"`
}

// Ping 路由
// @Summary  ping 路由
// @Description  ping 路由
// @Tags     default
// @Security BearerAuth
// @Accept   json
// @Produce  json
// @Success  200  {string}  pong  "pong"
// @Router   /api/v1/ping [get]
func Ping(c *gin.Context) {
	// c.Header("Content-Type", "charset=utf-8")
	c.String(200, "pong")
}
