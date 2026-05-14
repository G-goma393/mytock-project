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

var encryptCmd = &cobra.Command{
	Use:   "encrypt [fileName]",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start encrypt")
		targetName := args[0]

		task := CryptTask{
			FileName: targetFile,
			EncryptName: targetFile + ".tle"
			Network: SharedNetwork,
		}
		
		in, _ := os.Open(targetName)
		defer in.Close()

		//514am:durationはコマンド引数から受け取るよう
		// それからより詳細な設計書を作ろう

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

}
