package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, args []string) string {

	msg := ""

	itemreq := pocket.ItemRequest{}
	itemreq.ConsumerKey = pc.GetConsumerKey()
	itemreq.AccessToken = pc.GetAccessToken()
	itemreq.State = pocket.ItemStateAll
	itemreq.ContentType = string(pocket.ContentTypeArticle)
	itemreq.DetailType = string(pocket.DetailTypeSimple)
	itemreq.Tag = "Golang"
	itemreq.Count = 1

	itemresp := &pocket.ItemResponse{}
	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = "Error retrieving from Pocket" + err.Error()
	}

	fmt.Println(itemresp.WordCount)

	return msg
}
