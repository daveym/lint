package cmd

import (
	"fmt"

	"github.com/daveym/lint/api"
	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve items stored witin your Pocket data store",
	Long: `Retrieve provides you with the means to query Pocket, pull back items by specific order 
	and then use the results to perform follow up actions such as deleting or tagging`,
	Run: func(cmd *cobra.Command, args []string) {

		pc := &pocket.Client{}

		pc.SetConsumerKey(viper.GetString("ConsumerKey"))
		pc.SetAccessToken(viper.GetString("AccessToken"))

		msg := api.Retrieve(pc, args)
		fmt.Println(msg)
	},
}

func init() {
	RootCmd.AddCommand(retrieveCmd)
}
