package api

import (
	"testing"

	"github.com/daveym/lint/pocket"
)

func TestRetrieveNoConsumerKey(t *testing.T) {

	t.Log("Executing: TestRetrieveNoConsumerKey")

	var searchVal string
	var domainVal string
	var tagVal string
	var countVal int

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("")

	expectedmsg := "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
	actualmsg := Retrieve(mc, searchVal, domainVal, tagVal, countVal)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestValidConsumerKey failed")
	}
}
