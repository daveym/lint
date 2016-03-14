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
		t.Fatal("TestValidConsumerKey failed")
	}
}

func TestAuthInvalidConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("INVALIDKEY")

	expectedmsg := "Please check your consumer key, it does not appear to be valid."
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Fatal("TestAuthInvalidConsumerKey failed")
	}
}

func TestAuthAccessTokenExists(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("12345")
	mc.SetAccessToken("12345")

	expectedmsg := "Already authenticated - Access Token already present in lint.yaml"
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Fatal("TestAuthAccessTokenExists failed")
	}
}

/*
func TestBrowserAuthFail(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("INVALIDBROWSER")

	expectedmsg := "Error whilst approving Lint access to Pocket data. Please check your connectivity/default browser."
	actualmsg, err := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Fatal("TestAuthAccessTokenExists failed")
	}
}



func TestAuthGetAccessTokenSuccess(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := GetAccessToken(mc)

	t.Fatal("TestGetAccessTokenSuccess failed")
}

func TestAuthGetAccessTokenFail(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := GetAccessToken(mc)
	t.Fatal("TestGetAccessTokenFail failed")
}
*/
