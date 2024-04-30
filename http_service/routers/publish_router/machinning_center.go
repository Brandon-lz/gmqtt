package publishrouter


import (

	pubsub "github.com/Brandon-lz/gmqtt/pub_sub"
	"github.com/gin-gonic/gin"
)

// PublishHandler router ---------------------------------------
// @Tags Publish
// @Summary Publish a message to a topic
// @Description Publish a message to a topic
// @Produce json
// @Param topic path string true "Topic name"
// @Param message body string true "Message to publish"
// @Success 200 {object} core.ApiOKResponse
// @Router /api/v1/publish/:topic [post]
func PublishHandler(c *gin.Context) {
	// TODO: Implement Machinning Center Data In Router
	topic := c.Param("topic")

	pubsub.Agent.Publish(topic,"abc")

}


