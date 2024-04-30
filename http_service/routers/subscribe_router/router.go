package subscriberouter

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.RouterGroup) {
	group := router.Group("/subscribe")
	group.GET("/:topic", SubscribeHandler)
}

// SubscribeHandler websocket router ----------------------------------
// @Summary Send To Front End by websocket
// @Description # 数据订阅接口
// @Description ## 使用方法
// @Description ```javascript
// @Description var ws = new WebSocket("ws://127.0.0.1:8060/api/v1/subscribe/topic");
// @Description
// @Description ws.onmessage = function(event) {
// @Description   console.log(event.data);
// @Description };
// @Description
// @Description ws.onopen = function() {
// @Description   console.log("Connection established");
// @Description };
// @Description
// @Description ws.onclose = function() {
// @Description   console.log("Connection closed");
// @Description };
// @Description ```
// @Tags Subscribe
// @Accept json
// @Produce json
// @Router /api/v1/subscribe/{topic} [get]
// @Param topic path string true "topic name"
// @Success 200 {string} string "Hello, World!"
func SubscribeHandler(c *gin.Context) {
	topic := c.Param("topic")
	slog.Info(fmt.Sprintf("new subcribe online: topic=%s ", topic))
	Subscribe(c, topic)
}

// javascript code to receive data from server:
/*

var ws = new WebSocket("ws://localhost:8060/api/v1/data-out/machinning-center");

ws.onmessage = function(event) {

  console.log(event.data);
};

ws.onopen = function() {

  console.log("Connection established");
};


ws.onclose = function() {

  console.log("Connection closed");
};


*/
