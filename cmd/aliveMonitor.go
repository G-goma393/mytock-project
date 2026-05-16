package cmd

//mytock-projectアプリケーションの死活情報を内部的に確認するコマンド
// アプリケーションは応答しているか
// dorandネットワークは初期化できたか
// といった情報を確認できる
import (
	"fmt"

	"github.com/spf13/cobra"
)

var aliveMonitorCmd = &cobra.Command{
	Use:   "aliveMonitor",
	Short: "Life and death monitor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("mytock-project called")
		fmt.Println("initializing network")
		fmt.Println("application ready")

	},
}

func init() {
	rootCmd.AddCommand(aliveMonitorCmd)
}
