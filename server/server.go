package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

func StartTestServer() error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Upgrade error:", err)
			return
		}
		defer conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
					fmt.Println("connection closed")
				} else {
					fmt.Println("read error:", err)
				}
				return
			}
			fmt.Println("Received from client:", string(message))
			err = conn.WriteMessage(mt, message)
			if err != nil {
				fmt.Println("write error:", err)
				return
			}
		}
	})

	server := &http.Server{Addr: ":8080"}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("ListenAndServe(): %v\n", err)
		}
	}()

	fmt.Println("Test server started at ws://localhost:8080")

	// Handle shutdown gracefully
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("Shutting down test server...")
	return server.Shutdown(context.Background())
}
