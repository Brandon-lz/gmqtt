package publishrouter

import (
	"fmt"
	"log/slog"

	pubsub "github.com/Brandon-lz/gmqtt/pub_sub"
	gopubsub "github.com/Brandon-lz/go-pubsub"
	"github.com/gin-gonic/gin"
)

// PublishHandler router ---------------------------------------
// @Tags Publish
// @Summary Publish a message to a topic
// @Description Publish a message to a topic
// @Produce json
// @Param topic path string true "Topic name"
// @Param message query string true "Message to publish"
// @Success 200 {object} core.ApiOKResponse
// @Router /api/v1/publish/{topic} [post]
func PublishHandler(c *gin.Context) {
	topic := c.Param("topic")
	message := c.Query("message")
	slog.Info(fmt.Sprintf("Received publish request for topic: %s, message: %s", topic, message))
	pubsub.Agent.Publish(topic,gopubsub.MsgT(message))
}


