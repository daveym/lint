package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Modify against Pocket API. Interface used to allow mock to be passed in.
func Modify(pc pocket.API, action string, itemVal string, args []string) string {

	msg := ""

	fmt.Println(action, itemVal)
	for i := 0; i < len(args); i++ {
		fmt.Println("Echo: ", i, " ", args[i])
	}

	if len(pc.GetConsumerKey()) == 0 {
		msg = "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
		return msg
	}

	if len(itemVal) == 0 {
		msg = "Please specify a modify parameter (-a : add)"
		return msg
	}

	return msg
}
