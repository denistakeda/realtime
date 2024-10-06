package main

import (
	"fmt"
	"realtime/internal/server"
)

func main() {

  server := server.New(server.Params{
    Addr: ":8080",
  })


  fmt.Printf("Server stopped unexpectedly: %v", server.Start())
}
