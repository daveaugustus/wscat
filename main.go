package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	// Replace with the actual websocket server URL
	u := "wss://localhost:8080/api/v1/blog/draft/6564321"

	headers := http.Header{}
	headers.Set("Authorization", "Bearer ajd123bnsdjbcbbdbfrb2e89we89cksdbcdsbcbd")

	// Connect to the websocket server
	conn, _, err := websocket.DefaultDialer.Dial(u, headers)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer conn.Close()

	fmt.Println("Connected to the websocket server")

	// Define the message to send
	message := `
	{
		"time": 1278687357,
		"blocks": [
			{
				"id": "mhTl6ghSkV",
				"type": "paragraph",
				"data": {
					"text": "Hey. Meet the new Editor. On this picture you can see it in action. Then, try a demo 🤓"
				},
				"author": "John",
				"time": 17107371588
			}
		]
	}
	`

	// Send the message
	err = conn.WriteMessage(websocket.TextMessage, []byte(message))
	if err != nil {
		log.Println("write:", err)
		return
	}

	// Read any incoming messages (optional)
	_, msg, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}

	fmt.Printf("Received from server: %s\n", msg)
}
