package cmd

import (
	"fmt"
	"github.com/drand/tlock"
	tlockHttp "github.com/drand/tlock/networks/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var cfgFile string
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
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "configuration file path : $HOME/.mytock.yaml")
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".mytock")
	}
	//環境変数まわり
	viper.SetEnvPrefix("mytock") //例:MYTOCK_HOSTがあれば読み込むように
	viper.AutomaticEnv()
	//デフォルト値
	viper.SetDefault("host", "https://api.drand.sh/")
	viper.SetDefault("chainHash", "52db9ba70e0cc0f6eaf7803dd07447a1f5477735fd3f661792ba94600c84e971")

	if err := viper.ReadInConfig(); err == nil {
		//usedでいいのか楽だなぁ
		fmt.Println("configuration file loading successful : ", viper.ConfigFileUsed())
	}
}

func setupInfrastructure() error {
	host := viper.GetString("host")
	chainHash := viper.GetString("chainHash")

	net, err := tlockHttp.NewNetwork(host, chainHash)
	if err != nil {
		return fmt.Errorf("drant network setup failed: %w", err)
	}

	SharedNetwork = net
	return nil
}
