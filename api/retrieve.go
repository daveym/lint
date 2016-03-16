package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, itemreq pocket.ItemRequest, itemresp pocket.ItemResponse) string {

	msg := ""

	fmt.Println(pc.GetConsumerKey())
	fmt.Println(pc.GetAccessToken())
	fmt.Println(itemreq.ConsumerKey)
	fmt.Println(itemreq.AccessToken)
	fmt.Println(itemreq.Count)
	fmt.Println(itemreq.DetailType)

	err := pc.Retrieve(itemreq, itemresp)

	fmt.Println(itemresp)

	if err != nil {
		fmt.Println(err.Error())
		msg = "Error retrieving from Pocket"
	}

	return msg
}
