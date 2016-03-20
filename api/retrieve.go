package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, searchVal string, domainVal string, tagVal string, countVal int) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
		return msg
	}

	itemreq := pocket.RetrieveRequest{}
	itemreq.ConsumerKey = pc.GetConsumerKey()
	itemreq.AccessToken = pc.GetAccessToken()
	itemreq.Search = searchVal
	itemreq.Domain = domainVal
	itemreq.Tag = tagVal
	itemreq.Count = countVal
	itemresp := &pocket.RetrieveResponse{}

	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = "Error retrieving from Pocket: " + err.Error()
	}

	items := itemresp.List
	for _, item := range items {
		msg = msg + fmt.Sprintf("%v %v %v\n", item.ItemID, item.GivenTitle, item.GivenURL)
	}

	return msg
}
