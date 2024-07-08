package cmd

import (
	"github.com/daveaugustus/wscli/client"
	"github.com/daveaugustus/wscli/server"

	"github.com/spf13/cobra"
)

var (
	endpoint string
	headers  map[string]string
	rootCmd  = &cobra.Command{
		Use:   "ws-client",
		Short: "Golang websocket client with Cobra",
	}
	clientCmd = &cobra.Command{
		Use:   "run-client",
		Short: "Run the WebSocket client",
		RunE: func(cmd *cobra.Command, args []string) error {
			return client.ConnectAndRun(endpoint, headers)
		},
	}
	startServerCmd = &cobra.Command{
		Use:   "start-server",
		Short: "Start a test WebSocket server",
		RunE: func(cmd *cobra.Command, args []string) error {
			return server.StartTestServer()
		},
	}
)

func init() {
	clientCmd.Flags().StringVarP(&endpoint, "endpoint", "e", "", "Websocket server endpoint URL")
	clientCmd.Flags().StringToStringVarP(&headers, "headers", "H", nil, "Headers (key=value format)")
	_ = clientCmd.MarkFlagRequired("endpoint") // Mark endpoint flag as required

	rootCmd.AddCommand(clientCmd)
	rootCmd.AddCommand(startServerCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
