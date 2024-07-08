package client

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/daveaugustus/wscli/utils"

	"github.com/gorilla/websocket"
)

var conn *websocket.Conn

func ConnectAndRun(endpoint string, headers map[string]string) error {
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(endpoint, utils.ParseHeaders(headers))
	if err != nil {
		return fmt.Errorf("dial error: %w", err)
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established successfully")
	fmt.Print("\nEnter message (or 'exit' to quit): ")
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		cancel()
		conn.Close()
	}()

	var wg sync.WaitGroup
	wg.Add(2)

	canSend := make(chan bool, 1)
	canSend <- true
	promptChan := make(chan struct{})

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Send routine stopped due to context cancellation")
				return
			case <-canSend:
				message := utils.ReadUserInput()
				if message == "exit" {
					cancel()
					return
				}
				err := conn.WriteMessage(websocket.TextMessage, []byte(message))
				if err != nil {
					fmt.Println("write error:", err)
					cancel()
					return
				}
				canSend <- false
			}
		}
	}()

	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Receive routine stopped due to context cancellation")
				return
			default:
				_, message, err := conn.ReadMessage()
				if err != nil {
					if websocket.IsCloseError(err, websocket.CloseNormalClosure) {
						fmt.Println("connection closed")
					} else {
						fmt.Println("read error:", err)
					}
					cancel()
					return
				}
				fmt.Println("Received from server:", string(message))
				promptChan <- struct{}{}
			}
		}
	}()

	go func() {
		for range promptChan {
			fmt.Print("\nEnter message (or 'exit' to quit): ")
			canSend <- true
		}
	}()

	wg.Wait()
	return nil
}
