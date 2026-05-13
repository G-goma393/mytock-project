package cmd

import (
	"fmt"
	"github.com/drand/tlock"
	tlockHttp "github.com/drand/tlock/networks/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var SharedNetwork tlock.Network

var rootCmd = &cobra.Command{
	Use:   "mytock-project",
	Short: "共通設定やネットワークの初期化",
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		return setupInfrastructure()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
}

func initConfig() {
	viper.SetDefault("host", "https://api.drand.sh/")
	viper.SetDefault("chainHash", "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971")
}

func setupInfrastructure() error {
	host := viper.GetString("host")
	chainHash := viper.GetString("chainHash")

	net, err := tlockHttp.NewNetwork(host, chainHash)
	if err != nil {
		return fmt.Errorf("ネットワークの構築に失敗しました: %w", err)
	}

	SharedNetwork = net
	return nil
}
