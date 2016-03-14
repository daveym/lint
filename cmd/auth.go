package cmd

import (
	"fmt"
	"os"

	"github.com/daveym/lint/api"
	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Pocket  to get an Access Token. One off activity.",
	Long: `To access the Pocket API, you need to authenticate with your consumer key. Your consumer key 
	can be found under the development area within the pocket website`,
	Run: func(cmd *cobra.Command, args []string) {

		pc := &pocket.Client{}

		pc.SetConsumerKey(viper.GetString("ConsumerKey"))
		pc.SetAccessToken(viper.GetString("AccessToken"))

		msg, err := api.Authenticate(pc)

		if err != nil {
			fmt.Println(msg)
			os.Exit(-1)
		}
	}}

func init() {
	RootCmd.AddCommand(authCmd)
}
