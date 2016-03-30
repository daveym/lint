package api

import (
	"strconv"

	"github.com/daveym/lint/pocket"
)

// Modify against Pocket API. Interface used to allow mock to be passed in.
func Modify(pc pocket.API, action string, itemVal int, args []string) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = pocket.CONSUMERKEYNOTVALIDEn
		return msg
	}

	switch action {
	case "archive", "delete", "favourite", "unfavorite", "readd":
		msg = applyAction(pc, action, itemVal, args)
	case "add":
		msg = "Add not yet implemented."
	}

	return msg
}

func applyAction(pc pocket.API, action string, itemVal int, args []string) string {

	msg := ""
	modreq := pocket.ModifyRequest{}
	modreq.ConsumerKey = pc.GetConsumerKey()
	modreq.AccessToken = pc.GetAccessToken()

	modact := &pocket.Action{}
	modact.Action = action
	modact.ItemID = itemVal
	modreq.Actions = append(modreq.Actions, modact)

	// Bulk update, with itemIds in Args parameter
	if len(args) > 0 {
		for i := 0; i < len(args); i++ {
			modact := &pocket.Action{}
			modact.Action = action
			modact.ItemID, _ = strconv.Atoi(args[i])
			modreq.Actions = append(modreq.Actions, modact)
		}
	}

	modresp := &pocket.ModifyResponse{}
	err := pc.Modify(modreq, modresp)

	if err != nil {
		msg = pocket.ERRORRETRIEVINGen + err.Error()
		return msg
	}

	if modresp.Status == 0 {
		msg = pocket.ERROREXECUTINGen + action
		return msg
	}

	return pocket.UPDATEAPPLIEDen
}
