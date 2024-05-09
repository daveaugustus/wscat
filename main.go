package main

import (
	"bufio"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/spf13/cobra"
)

var (
	endpoint string
	headers  map[string]string
	conn     *websocket.Conn
	rootCmd  = &cobra.Command{
		Use:   "ws-client",
		Short: "Golang websocket client with Cobra",
		RunE: func(cmd *cobra.Command, args []string) error {
			return connectAndRun()
		},
	}
)

func init() {
	rootCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "Websocket server endpoint URL")
	rootCmd.Flags().StringToStringVarP(&headers, "headers", "H", nil, "Headers (key=value format)")
	_ = rootCmd.MarkFlagRequired("endpoint") // Mark endpoint flag as required
}

func connectAndRun() error {
	// Establish websocket connection
	var err error
	conn, _, err = websocket.DefaultDialer.Dial(endpoint, parseHeaders(headers))
	if err != nil {
		return fmt.Errorf("dial error: %w", err)
	}
	defer conn.Close()

	fmt.Println("WebSocket connection established successfully")

	// Handle CTRL+C to close connection gracefully
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	go func() {
		<-ch
		cancel()
		conn.Close()
	}()

	// Create a WaitGroup to wait for goroutines to finish
	var wg sync.WaitGroup
	wg.Add(2)

	// Goroutine for sending data
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Send routine stopped due to context cancellation")
				return
			default:
				message := readUserInput()
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
			}
		}
	}()

	// Goroutine for receiving data
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
			}
		}
	}()

	// Wait for routines to finish
	wg.Wait()
	return nil
}

func readUserInput() string {
	fmt.Print("Enter message (or 'exit' to quit): ")
	reader := bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')
	return strings.TrimSpace(message)
}

func parseHeaders(headers map[string]string) http.Header {
	headerMap := http.Header{}
	for key, value := range headers {
		headerMap.Add(key, value)
	}
	return headerMap
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
