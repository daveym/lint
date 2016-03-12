package cmd

import (
	"github.com/daveym/lint/api"
	"github.com/spf13/cobra"
)

var authenticateCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Pocket  to get an Access Token. One off activity.",
	Long: `To access the Pocket API, you need to authenticate with your consumer key. Your consumer key 
	can be found under the development area within the pocket website`,
	Run: func(cmd *cobra.Command, args []string) {

		Api.Authenticate()

	}}

func init() {
	RootCmd.AddCommand(authenticateCmd)
}
