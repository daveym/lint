package cmd

import (
	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete items stored witin your Pocket data store.",
	Long:  `Delete permanently removes items stored within Pocket. Be careful, once you do this, they are gone forever!`,
}

func init() {

	deleteCmd.Run = delete
	RootCmd.AddCommand(deleteCmd)
}

func delete(cmd *cobra.Command, args []string) {

	pc := &pocket.Client{}
	pc.SetConsumerKey(viper.GetString("ConsumerKey"))
	pc.SetAccessToken(viper.GetString("AccessToken"))

	//msg := api.Retrieve(pc, searchVal, domainVal, tagVal, countVal)
	//fmt.Println(msg)
}
