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

		itemreq := pocket.ItemRequest{}
		itemreq.ConsumerKey = pc.GetConsumerKey()
		itemreq.AccessToken = pc.GetAccessToken()
		itemreq.State = pocket.ItemStateAll
		itemreq.ContentType = string(pocket.ContentTypeArticle)
		itemreq.DetailType = string(pocket.DetailTypeSimple)
		itemreq.Count = 10

		itemresp := pocket.ItemResponse{}
		msg := api.Retrieve(pc, itemreq, itemresp)

		fmt.Println(msg)
		fmt.Println(itemresp.GivenTitle)
	},
}

func init() {
	RootCmd.AddCommand(retrieveCmd)
}
