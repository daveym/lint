package cmd

import (
	"fmt"
	"strings"

	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var retrieveCmd = &cobra.Command{
	Use:   "retrieve",
	Short: "Retrieve items stored witin your Pocket data store. Use flag -s, -d to specific a search term or domain",
	Long: `Retrieve provides you with the means to query Pocket, pull back items by specific order 
	and then use the results to perform follow up actions such as deleting or tagging`,
}

func init() {
	// Setup retrieve flags against the 'retrieve' command
	var searchVal string
	var domainVal string

	retrieveCmd.Run = retrieve
	retrieveCmd.Flags().StringVarP(&searchVal, "search", "s", "", "Only return items whose title or url contain the search string")
	retrieveCmd.Flags().StringVarP(&domainVal, "domain", "d", "", "Only return items from a particular domain")

	RootCmd.AddCommand(retrieveCmd)
	RootCmd.Execute()
}

func retrieve(cmd *cobra.Command, args []string) {

	for i := 0; i < len(args); i++ {
		fmt.Println(strings.Join(args, " "))
	}

	pc := &pocket.Client{}

	pc.SetConsumerKey(viper.GetString("ConsumerKey"))
	pc.SetAccessToken(viper.GetString("AccessToken"))

	//msg := api.Retrieve(pc, args)

	//fmt.Print(msg)
}
