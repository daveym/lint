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

		ireq := pocket.ItemRequest{}
		ireq.ConsumerKey = pc.GetConsumerKey()
		ireq.AccessToken = pc.GetConsumerKey()
		ireq.State = pocket.ItemStateAll
		ireq.ContentType = string(pocket.ContentTypeArticle)
		ireq.DetailType = string(pocket.DetailTypeSimple)
		ireq.Count = 10

		irsp := pocket.ItemResponse{}
		msg := api.Retrieve(pc, ireq, irsp)

		fmt.Println(msg)

		fmt.Println(irsp.GivenTitle)
	},
}

func init() {
	RootCmd.AddCommand(retrieveCmd)
}
