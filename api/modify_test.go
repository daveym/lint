package api

import (
	"testing"

	"github.com/daveym/lint/pocket"
)

func TestModifyNoConsumerKey(t *testing.T) {

	t.Log("Executing: TestModifyNoConsumerKey")

	var action string
	var itemVal string
	var args []string

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("")

	action = ""

	expectedmsg := "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestModifyNoConsumerKey failed")
	}
}

func TestModifyNoCriteria(t *testing.T) {

	t.Log("Executing: TestModifyNoCriteria")

	var action string
	var itemVal string
	var args []string

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	action = ""
	itemVal = ""

	expectedmsg := "Please specify a modify parameter (-a : add)"
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestRetrieveNoCriteria failed")
	}
}
