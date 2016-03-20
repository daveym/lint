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
	itemreq.Count = 15

	itemresp := &pocket.RetrieveResponse{}
	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = "Error retrieving from Pocket: " + err.Error()
	}

	items := itemresp.List
	for _, item := range items {
		msg = fmt.Sprintf("%v %v %v\n", item.ItemID, item.GivenTitle, item.GivenURL)
	}

	return msg
}
