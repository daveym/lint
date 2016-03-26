package api

import "github.com/daveym/lint/pocket"

// Modify against Pocket API. Interface used to allow mock to be passed in.
func Modify(pc pocket.API, addCmdVal string) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
		return msg
	}

	if len(addCmdVal) == 0 {
		msg = "Please specify a modify parameter (-a : add)"
		return msg
	}

	msg = addCmdVal

	return msg
}
