package cmd

import (
	"../internal/pkg/api"
	"errors"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"
)

// shareCmd represents the share command
var shareCmd = &cobra.Command{
	Use:   "share",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("share command requires a filepath")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		client := api.FilesClient{BaseUri: "http://localhost/api/v1/files", Client: http.Client{}}

		filePath := args[0]

		resp, err := client.Upload(filePath)

		if err != nil {
			fmt.Printf("Error: %s", err)
		} else {
			fmt.Printf("Download (one-time use): %s\nDelete: %s", resp.Data.DownloadUri, resp.Data.DeleteUri)
		}
	},
}

func init() {
	rootCmd.AddCommand(shareCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// shareCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// shareCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
