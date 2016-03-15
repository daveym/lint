package api

import "github.com/daveym/lint/pocket"

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, ir pocket.ItemRequest, irsp pocket.ItemResponse) string {

	msg := ""
	err := pc.Retrieve(ir, &irsp)

	if err != nil {
		msg = "Error retrieving from Pocket"
	}

	return msg
}
