// cmd/dataset.go
package cmd

import (
	"godeep/internal/client"

	"github.com/spf13/cobra"
)

var createDatasetCmd = &cobra.Command{
	Use:   "create-dataset",
	Short: "Create a new dataset",
	Long: `Create a new dataset in DeepLake with tensor_db enabled.
Example: godeep create-dataset --path="hub://org/dataset"`,
	RunE: func(cmd *cobra.Command, args []string) error {
		c, err := client.NewDeepLakeClient(cfg)
		if err != nil {
			return err
		}
		return c.CreateDataset()
	},
}

func init() {
	rootCmd.AddCommand(createDatasetCmd)

	createDatasetCmd.Flags().String("path", "", "Dataset path")
	createDatasetCmd.Flags().Bool("tensor-db", true, "Enable tensor_db")
}
