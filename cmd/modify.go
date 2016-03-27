package cmd

import (
	"fmt"

	"github.com/daveym/lint/api"
	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var itemID string

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify items stored witin your Pocket data store (add, archive, re-add, favourite, unfavourite, delete)",
	Long: `Modify your pocket items. 
	Add - Add a new item to pocket using ./lint modify -a itemID, tags (comma separated), title, url
	Archive - Move an item to pockets archive
	Re-add - (unarchive) an item from archive
	Favorite - Mark an item as a favorite
	Unfavorite - Remove an item from your favorites
	Delete - Permanently remove an item from your pocket account
	`,
}

func init() {

	modifyCmd.Run = modify

	// itemID is the first parameter to each of the commands.
	modifyCmd.Flags().StringVarP(&itemID, "add", "a", "", "Add a new item to pocket using ./lint modify -a itemID, tags(comma separated), title, url")
	modifyCmd.Flags().StringVarP(&itemID, "del", "d", "", "Delete an  item from pocket using ./lint modify -d itemID")

	RootCmd.AddCommand(modifyCmd)
}

func modify(cmd *cobra.Command, args []string) {

	msg := ""

	pc := &pocket.Client{}
	pc.SetConsumerKey(viper.GetString("ConsumerKey"))
	pc.SetAccessToken(viper.GetString("AccessToken"))

	switch {
	case cmd.Flag("add").Changed:
		msg = api.Modify(pc, "add", itemID, args)
	case cmd.Flag("del").Changed:
		msg = api.Modify(pc, "del", itemID, args)
	}

	fmt.Println(msg)
}
