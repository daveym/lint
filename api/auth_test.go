package api

import (
	"testing"

	"github.com/daveym/lint/pocket"
)

var err error

func TestAuthNoConsumerKey(t *testing.T) {

	t.Log("Executing: TestAuthNoConsumerKey")

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("")

	expectedmsg := pocket.CONSUMERKEYNOTFOUNDEn
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestValidConsumerKey failed")
	}
}

func TestAuthInvalidConsumerKey(t *testing.T) {

	t.Log("Executing: TestAuthInvalidConsumerKey")

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("INVALIDKEY")

	expectedmsg := pocket.CONSUMERKEYNOTVALIDEn
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthInvalidConsumerKey failed")
	}
}

func TestBrowserAuthFail(t *testing.T) {

	t.Log("Executing: TestBrowserAuthFail")

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("INVALIDBROWSER")

	expectedmsg := pocket.ERRORAPPROVINGLINTEn
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestBrowserAuthFail failed")
	}
}

func TestAuthGetAccessTokenFail(t *testing.T) {

	t.Log("Executing: TestAuthGetAccessTokenFail")

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("FAIL")
	mc.SetAccessToken("FAIL")

	expectedmsg := pocket.ERRORAUTHen
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthGetAccessTokenFail failed")
	}
}

func TestAuthGetAccessTokenSuccess(t *testing.T) {

	t.Log("Executing: TestAuthGetAccessTokenSuccess")

	mc := &pocket.MockClient{}
	mc.SetConsumerKey("45678")
	mc.SetAccessToken("SUCCESS")

	expectedmsg := pocket.AUTHSUCCESSen
	actualmsg := Authenticate(mc)

	if actualmsg != expectedmsg {
		t.Log("Expected: " + expectedmsg)
		t.Log("Actual: " + actualmsg)
		t.Fatal("TestAuthGetAccessTokenSuccess failed")
	}
}
