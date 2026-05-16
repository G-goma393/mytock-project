package cmd

import (
	"bytes"
	"fmt"
	"github.com/drand/tlock"
	"github.com/spf13/cobra"
	"os"
)

var decryptCmd = &cobra.Command{
	Use:   "decrypt [fileName].tle",
	Short: "(ショート説明)Decrypt the time lock cipher",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start decrypt")
		targetFile := args[0]

		task := CryptTask{
			EncryptedName: targetFile,
			Network:       SharedNetwork,
		}

		encIn, _ := os.Open(task.EncryptedName)
		defer encIn.Close()

		var plainData bytes.Buffer
		if err := tlock.New(task.Network).Decrypt(&plainData, encIn); err != nil {
			fmt.Printf("Decryption failed ... as expected: %v\n", err)
		} else {
			fmt.Printf("SUCCESSFUL \n source: %s\n", plainData.String())
		}
	},
}

func init() {
	rootCmd.AddCommand(decryptCmd)
}
