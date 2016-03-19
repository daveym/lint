package api

import "github.com/daveym/lint/pocket"

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, args []string) string {

	msg := ""

	itemreq := pocket.RetrieveRequest{}
	itemreq.ConsumerKey = pc.GetConsumerKey()
	itemreq.AccessToken = pc.GetAccessToken()
	itemreq.Tag = "docker"
	itemreq.Count = 1

	itemresp := &pocket.RetrieveResponse{}
	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = "Error retrieving from Pocket: " + err.Error()
	}

	/*	items := []api.Item{}
		for _, item := range res.List {
			items = append(items, item)
		}*/

	return msg
}
