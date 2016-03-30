package api

import (
	"fmt"

	"github.com/daveym/lint/pocket"
)

// Retrieve against Pocket API. Interface used to allow mock to be passed in.
func Retrieve(pc pocket.API, searchVal string, domainVal string, tagVal string, countVal int) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = pocket.CONSUMERKEYNOTVALIDEn
		return msg
	}

	if countVal == 0 {
		msg = pocket.COUNTGREATERTHANZEROen
		return msg
	}

	if len(searchVal) == 0 && len(domainVal) == 0 && len(tagVal) == 0 {
		msg = pocket.SPECIFYSEARCHen
		return msg
	}

	itemreq := pocket.RetrieveRequest{}
	itemreq.ConsumerKey = pc.GetConsumerKey()
	itemreq.AccessToken = pc.GetAccessToken()

	if searchVal != "all" {
		itemreq.Search = searchVal
	}

	itemreq.Domain = domainVal
	itemreq.Tag = tagVal
	itemreq.Count = countVal
	itemreq.Sort = pocket.SortNewest

	itemresp := &pocket.RetrieveResponse{}

	err := pc.Retrieve(itemreq, itemresp)

	if err != nil {
		msg = pocket.ERRORRETRIEVINGen + err.Error()
		return msg
	}

	items := itemresp.List

	if len(items) == 0 {
		msg = pocket.NOMATCHINGVALUESen
		return msg
	}

	for _, item := range items {
		msg = msg + fmt.Sprintf("%v,%v,%v\n", item.ItemID, item.GivenTitle, item.GivenURL)
	}

	return msg
}
