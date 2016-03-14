package api

import (
	"testing"

	"github.com/daveym/lint/pocket"
)

func TestAuthNoConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("")

	expectedmsg := "Consumer Key not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
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
		t.Fatal("TestAuthAccessTokenExists failed")
	}
}

/*
func TestAuthInvalidConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("4623546")

	expectedmsg := "Please check your consumer key, it does not appear to be valid."
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Fatal("TestAuthInvalidConsumerKey failed")
	}
}


func TestAuthValidConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}

	expectedmsg := "Please check your consumer key, it does not appear to be valid."
	actualmsg := Authenticate(mc)

	t.Fatal("TestInvalidConsumerKey failed")
}

func TestAuthBrowserPocketSuccess(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := UserAuthorise(mc)

	t.Fatal("TestBrowserPocketAuthSuccess failed")
}

func TestAuthBrowserPocketFail(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := AuthoriseUse(mc)

	t.Fatal("TestBrowserPocketAuthFail failed")
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
