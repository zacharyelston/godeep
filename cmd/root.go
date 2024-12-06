// cmd/root.go
package cmd

import (
	"fmt"
	"godeep/config"
	"godeep/internal/version"
	"os"

	"github.com/spf13/cobra"
)

var (
	cfgFile string
	cfg     *config.Config
)

var rootCmd = &cobra.Command{
	Use:     "godeep",
	Short:   "DeepLake CLI tool for tensor database operations",
	Version: version.Version,
	Long: `A CLI tool for interacting with DeepLake tensor database.
Complete documentation is available at https://docs.activeloop.ai`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ./config/default.yaml)")

	rootCmd.PersistentFlags().String("token", "", "ActiveLoop API token")
	rootCmd.PersistentFlags().String("org-id", "", "ActiveLoop organization ID")
	rootCmd.PersistentFlags().String("dataset-path", "", "Dataset path")
}

func initConfig() {
	var err error
	cfg, err = config.LoadConfig()
	if err != nil {
		fmt.Printf("Error loading config: %v\n", err)
		os.Exit(1)
	}
}
