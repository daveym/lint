package api

import (
	"testing"

	"github.com/daveym/lint/pocket"
)

var err error

func TestAuthNoConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("")

	expectedmsg := "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestValidConsumerKey failed")
	}
}

func TestAuthAccessTokenExists(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("12345")
	mc.SetAccessToken("12345")

	expectedmsg := "Already authenticated - Access Token already present in lint.yaml"
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthAccessTokenExists failed")
	}
}

func TestAuthInvalidConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("INVALIDKEY")

	expectedmsg := "Please check your consumer key, it does not appear to be valid."
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthInvalidConsumerKey failed")
	}
}

func TestBrowserAuthFail(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("INVALIDBROWSER")

	expectedmsg := "Error whilst approving Lint access to Pocket data. Please check your connectivity/default browser."
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestBrowserAuthFail failed")
	}
}

func TestAuthGetAccessTokenFail(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("12345")

	expectedmsg := "Error authorising your consumer key and request token. Have you granted permission to Lint?"
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthGetAccessTokenFail failed")
	}
}

/*
func TestAuthGetAccessTokenSuccess(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("12345")
	mc.SetAccessToken("SUCCESS")

	expectedmsg := "Authentication Successful - Access Token is persisted to lint.yaml"
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthGetAccessTokenSuccess failed")
	}
}*/
