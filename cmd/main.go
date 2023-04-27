package main

import (
	"log"

	"github.com/dsauerbrun/cragcast/pkg/http"
	"github.com/spf13/cobra"
)

var bindHost string
var port int

func init() {
	serveCmd.Flags().StringVar(&bindHost, "bind-host", "127.0.0.1", "The address to bind to. To accept all connections, use 0.0.0.0.")
	serveCmd.Flags().IntVar(&port, "port", api.DefaultPort, "The port to listen on.")

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(serveCmd)
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the server",
	Run: func(cmd *cobra.Command, args []string) {
		api.Start(api.StartOptions{
			BindHost: bindHost,
			Port:     port,
		})
	},
}

var rootCmd = &cobra.Command{
	Use:   "cragcast",
	Short: "Weather-related information for the world's climbing areas.",
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.Help()
		// err should never occur displaying help text. If it does,
		// panic is acceptable.
		if err != nil {
			panic(err)
		}
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Println(err)
	}
}
