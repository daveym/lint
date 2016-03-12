package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List items stored witin your Pocket data store",
	Long: `List provides you with the means to query Pocket, pull back items by specific order 
	and then use the results to perform follow up actions such as deleting or tagging`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(viper.GetString("AccessToken"))

		/*var ItemRequest LintApi.ItemRequest
		ItemRequest.ConsumerKey = LintApi.ConsumerKey
		ItemRequest.AccessToken = AccessToken
		ItemRequest.State = LintApi.ItemStateAll
		ItemRequest.ContentType = string(LintApi.ContentTypeArticle)
		ItemRequest.DetailType = string(LintApi.DetailTypeSimple)
		ItemRequest.Count = 10

		items, err := LintApi.Retrieve(ItemRequest)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(items.GivenTitle) */
	},
}

func init() {
	RootCmd.AddCommand(listCmd)
}
