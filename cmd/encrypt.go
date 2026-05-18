package cmd

import (
	"bytes"
	"fmt"
	"github.com/drand/tlock"
	"github.com/spf13/cobra"
	"log"
	"os"
	"time"
)

type CryptTask struct {
	FileName      string
	EncryptedName string
	Network       tlock.Network
}

var encryptCmd = &cobra.Command{
	Use:   "encrypt [fileName] --[duration]",
	Short: "(ショート説明)Encrypts the file with a time lock",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start encrypt")
		targetFile := args[0]

		task := CryptTask{
			FileName:      targetFile,
			EncryptedName: targetFile + ".tle",
			Network:       SharedNetwork,
		}

		in, _ := os.Open(task.FileName)
		defer in.Close()

		duration := 10 * time.Second
		roundNumber := task.Network.Current(time.Now().Add(duration))

		var cipherData bytes.Buffer
		if err := tlock.New(task.Network).Encrypt(&cipherData, in, roundNumber); err != nil {
			log.Fatalf("ohh, lol: %v", err)
		}
		os.WriteFile(task.EncryptedName, cipherData.Bytes(), 0644)
	},
}

func init() {
	rootCmd.AddCommand(encryptCmd)

}
