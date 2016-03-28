package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Modify against Pocket API. Interface used to allow mock to be passed in.
func Modify(pc pocket.API, action string, itemVal string, args []string) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
		return msg
	}

	if len(itemVal) == 0 {
		msg = "Please specify a modify parameter (-a : add)"
		return msg
	}

	switch action {
	case "add":
		add(pc, action, itemVal, args)
	case "del":
		fmt.Println("Del")

	}

	fmt.Println(action, itemVal)
	for i := 0; i < len(args); i++ {
		fmt.Println("Echo: ", i, " ", args[i])
	}

	return msg
}

func add(pc pocket.API, action string, itemVal string, args []string) string {

	msg := ""

	modreq := pocket.ModifyRequest{}
	modreq.ConsumerKey = pc.GetConsumerKey()
	modreq.AccessToken = pc.GetAccessToken()

	modact := pocket.Action{}
	modact.Action = "add"
	modact.ItemID = itemVal
	modact.Tags = args[0]
	modact.Title = args[1]
	modact.URL = args[2]

	modreq.Action = modact

	modresp := &pocket.ModifyResponse{}

	err := pc.Modify(modreq, modresp)

	fmt.Println(modresp)

	if err != nil {
		msg = "Error adding to Pocket: " + err.Error()
		return msg
	}

	if modresp.Status != 200 {
		fmt.Println(modresp)
	}

	return msg
}
