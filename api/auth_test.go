package api

import (
	"testing"

	"github.com/daveym/lint/api"
	"github.com/daveym/lint/pocket"
)

func TestAuthValidConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}

	expectedmsg := "Please check your consumer key, it does not appear to be valid."
	actualmsg := api.Authenticate(mc)

	t.Fatal("TestValidConsumerKey failed")
}

func TestAuthInvalidConsumerKey(t *testing.T) {

	mc := &pocket.MockClient{}

	expectedmsg := "Please check your consumer key, it does not appear to be valid."
	actualmsg := api.Authenticate(mc)

	t.Fatal("TestInvalidConsumerKey failed")
}

func TestAuthBrowserPocketSuccess(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := api.AuthoriseUse(mc)

	t.Fatal("TestBrowserPocketAuthSuccess failed")
}

func TestAuthBrowserPocketFail(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := api.AuthoriseUse(mc)

	t.Fatal("TestBrowserPocketAuthFail failed")
}

func TestAuthGetAccessTokenSuccess(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := api.GetAccessToken(mc)

	t.Fatal("TestGetAccessTokenSuccess failed")
}

func TestAuthGetAccessTokenFail(t *testing.T) {

	mc := &pocket.MockClient{}
	actualmsg := GetAccessToken(mc)
	t.Fatal("TestGetAccessTokenFail failed")
}
