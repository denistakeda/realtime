package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	r := gin.Default()

	r.GET("/ws", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		// Handle WebSocket connection
		go handleWebSocketConnection(conn)
	})

	r.Run(":8080")
}

func handleWebSocketConnection(conn *websocket.Conn) {
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			conn.Close()
			break
		}

		fmt.Println(string(message))

		// Send response to the client
		err = conn.WriteMessage(websocket.TextMessage, []byte("Message received"))
		if err != nil {
			conn.Close()
			break
		}
	}
}
