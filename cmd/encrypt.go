/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"github.com/drand/tlock"
	"github.com/drand/tlock/http"
	"github.com/spf13/cobra"
)

type CryptTask struct{
	FileName string
	EncryptedName string
	Network tlock.Network
}

// encryptCmd represents the encrypt command
var encryptCmd = &cobra.Command{
	Use:   "encrypt",
	Short: "short:Let's encrypt the file",
	Long:  `"long:Let's encrypt the file"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start encrypt")
		fileName := args[0]
		in, _ := os.Open(fileName)
		defer in.Close()

		network, _:= http.NewNetwork(host, chainHash)
		duration := 10 * time.second
		roundNumber := network.RoundNumber(time.Now().Add(duration))

		var cipherData bytes.Buffer
		if err := tlock.New(network).Encrypt(&cipherData, in, roundNumber); err != mil{
			log.Fatalf("ohh, lol: %v" err)
		}
		os.WriteFile(encryptedName, cipherData.Bytes(), 0644)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// encryptCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// encryptCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
