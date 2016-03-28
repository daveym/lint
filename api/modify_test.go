package api

import (
	"testing"

	"github.com/daveym/lint/pocket"
)

func TestModifyNoConsumerKey(t *testing.T) {

	t.Log("Executing: TestModifyNoConsumerKey")

	var action string
	var itemVal int
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
