package server

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type server struct {
	engine *gin.Engine
	addr   string
}

type Params struct {
	Addr string
}

func New(params Params) *server {
	s := &server{
		engine: gin.New(),
		addr:   params.Addr,
	}

	s.registerHandlers()
	return s
}

func (s *server) Start() error {
	return s.engine.Run(s.addr)
}

func (s *server) registerHandlers() {
	s.registerStaticHandlers()
  s.registerApiHandlers()
}

func (s *server) registerStaticHandlers() {
	s.engine.Static("/static", "./static")
	s.engine.StaticFile("/", "./static/index.html")
}

func (s *server) registerApiHandlers() {
	upgrader := websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	s.engine.GET("/ws", func(c *gin.Context) {
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
