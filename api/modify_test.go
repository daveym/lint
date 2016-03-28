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

func TestModifySuccessfulUpdate(t *testing.T) {

	t.Log("Executing: TestModifySuccessfulUpdate")

	var action string
	var itemVal int
	var args []string

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("SUCCESS")

	action = "archive"
	itemVal = 12345

	expectedmsg := "Update applied successfully."
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

	expectedmsg := "Error executing " + action + " against the pocket API."
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

	expectedmsg := "Update applied successfully."
	actualmsg := Modify(mc, action, itemVal, args)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestModifySuccessfulUpdate failed")
	}
}
