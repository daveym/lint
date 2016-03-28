package cmd

import (
	"fmt"
	"strconv"

	"github.com/daveym/lint/api"
	"github.com/daveym/lint/pocket"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var itemArg string

var modifyCmd = &cobra.Command{
	Use:   "modify",
	Short: "Modify items stored witin your Pocket data store (archive, favourite, unfavourite, delete)",
	Long: `Modify your pocket items. 
	Archive - Move an item to pockets archive
	Favorite - Mark an item as a favorite
	Unfavorite - Remove an item from your favorites
	Delete - Permanently remove an item from your pocket account
	`,
}

func init() {

	modifyCmd.Run = modify

	// itemID is the first parameter to each of the commands.
	modifyCmd.Flags().StringVarP(&itemArg, "add", "n", "", "Add a new item to pocket.")
	modifyCmd.Flags().StringVarP(&itemArg, "arc", "a", "", "Archive an item in pocket using ./lint modify -a itemID [itemID2 itemID3 itemID3 etc.]")
	modifyCmd.Flags().StringVarP(&itemArg, "del", "d", "", "Delete an  item from pocket using ./lint modify -d itemID [itemID2 itemID3 itemID3 etc.]")
	modifyCmd.Flags().StringVarP(&itemArg, "fav", "f", "", "Favourite an item in pocket using ./lint modify -f itemID [itemID2 itemID3 itemID3 etc.]")
	modifyCmd.Flags().StringVarP(&itemArg, "unfav", "u", "", "(Un)Favourite an item in pocket using ./lint modify -u itemID [itemID2 itemID3 itemID3 etc.]")
	modifyCmd.Flags().StringVarP(&itemArg, "readd", "r", "", "Readd an item that was archived pocket using ./lint modify -r itemID [itemID2 itemID3 itemID3 etc.]")

	RootCmd.AddCommand(modifyCmd)
}

func modify(cmd *cobra.Command, args []string) {

	msg := ""

	pc := &pocket.Client{}
	pc.SetConsumerKey(viper.GetString("ConsumerKey"))
	pc.SetAccessToken(viper.GetString("AccessToken"))

	itemVal, err := strconv.Atoi(itemArg)

	if err != nil {
		fmt.Println("ItemID paramemter must be numeric")
		return
	}

	switch {
	case cmd.Flag("arc").Changed:
		msg = api.Modify(pc, "archive", itemVal, args)
	case cmd.Flag("del").Changed:
		msg = api.Modify(pc, "delete", itemVal, args)
	case cmd.Flag("fav").Changed:
		msg = api.Modify(pc, "favourite", itemVal, args)
	case cmd.Flag("unfav").Changed:
		msg = api.Modify(pc, "unfavourite", itemVal, args)
	case cmd.Flag("readd").Changed:
		msg = api.Modify(pc, "readd", itemVal, args)
	case cmd.Flag("add").Changed:
		msg = api.Modify(pc, "add", itemVal, args)
	}

	fmt.Println(msg)
}
