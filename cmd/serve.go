/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var port int
var host string

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "しょーと説明ちなserbe",
	Long:  `long説明ちなserve`,
	Run: func(cmd *cobra.Command, args []string) {
		if verbose {
			fmt.Printf("starting server on %s:%d\n", host, port)
			fmt.Println("Verbose mode enabled")
		} else {
			fmt.Printf("server starting on %s:%d\n", host, port)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().IntVarP(&port, "port", "p", 8080, "Port to run the server on")
	serveCmd.Flags().StringVarP(&host, "host", "H", "localhost", "Host to bind the server to")

}
