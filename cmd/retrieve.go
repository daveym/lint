package cmd

import (
	"fmt"

	"github.com/daveym/lint/api"
	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchVal string
var domainVal string
var tagVal string
var countVal int

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve items stored witin your Pocket data store. Use flag -s, -d to specific a search term or domain",
	Long: `Retrieve provides you with the means to query Pocket, pull back items by specific order 
	and then use the results to perform follow up actions such as deleting or tagging`,
}

func init() {

	retrieveCmd.Run = retrieve
	retrieveCmd.Flags().StringVarP(&searchVal, "search", "s", "", "Only return items whose title or url contain the search string")
	retrieveCmd.Flags().StringVarP(&domainVal, "domain", "d", "", "Only return items from a particular domain")
	retrieveCmd.Flags().StringVarP(&tagVal, "tag", "t", "", "only return items tagged with tag name")
	retrieveCmd.Flags().IntVarP(&countVal, "count", "c", 10, "total number of items to return")

	RootCmd.AddCommand(retrieveCmd)
}

func retrieve(cmd *cobra.Command, args []string) {

	pc := &pocket.Client{}
	pc.SetConsumerKey(viper.GetString("ConsumerKey"))
	pc.SetAccessToken(viper.GetString("AccessToken"))

	msg := api.Retrieve(pc, searchVal, domainVal, tagVal, countVal)
	fmt.Println(msg)
}
