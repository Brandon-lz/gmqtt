package subscriberouter

import (
	"net/http"

	pubsub "github.com/Brandon-lz/gmqtt/pub_sub"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin:     func(r *http.Request) bool { return true },
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Subscribe(c *gin.Context, topic string) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	suber := pubsub.Agent.Subscribe(topic)
	defer pubsub.Agent.Unsubscribe(suber)

	// Send data to front-end
	for {
		// data := waitForData()
		err := conn.WriteMessage(websocket.TextMessage, []byte(<-suber.Msg))
		if err != nil {
			return
		}
	}

}
