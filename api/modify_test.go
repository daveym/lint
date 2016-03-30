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

	expectedmsg := pocket.CONSUMERKEYNOTVALIDEn
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestModifyNoConsumerKey failed")
	}
}

func TestModifySuccessfulUpdate(t *testing.T) {

	t.Log("Executing: TestModifySuccessfulUpdate")

	var action string
	var itemVal int
	var args []string

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("SUCCESS")

	action = "archive"
	itemVal = 12345

	expectedmsg := pocket.UPDATEAPPLIEDen
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestModifySuccessfulUpdate failed")
	}
}

func TestModifyFailedUpdate(t *testing.T) {

	t.Log("Executing: TestModifyFailedUpdate")

	var action string
	var itemVal int
	var args []string

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("SUCCESS")

	action = "archive"
	itemVal = 45678

	expectedmsg := pocket.ERROREXECUTINGen + action
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestModifyFailedUpdate failed")
	}
}

func TestModifySuccessfulUpdateMultipleArgs(t *testing.T) {

	t.Log("Executing: TestModifySuccessfulUpdate")

	var action string
	var itemVal int
	var args []string

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("SUCCESS")

	action = "archive"
	itemVal = 12345
	args = append(args, "45678")
	args = append(args, "91011")

	expectedmsg := pocket.UPDATEAPPLIEDen
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestModifySuccessfulUpdate failed")
	}
}
