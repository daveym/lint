package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, args []string) string {

	msg := ""

	itemreq := pocket.RetrieveRequest{}
	itemreq.ConsumerKey = pc.GetConsumerKey()
	itemreq.AccessToken = pc.GetAccessToken()
	itemreq.Count = 1

	itemresp := &pocket.RetrieveResponse{}
	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = "Error retrieving from Pocket: " + err.Error()
	}

	for _, item := range itemresp.List {
		fmt.Println(item.ItemID)
	}

	return msg
}
