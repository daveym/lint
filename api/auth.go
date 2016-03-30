package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daveym/lint/pocket"
)

// Authenticate against Pocket API. Interface used to allow mock to be passed in.
func Authenticate(pc pocket.API) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = pocket.CONSUMERKEYNOTFOUNDEn
		return msg
	}

	AuthNResp := &pocket.AuthenticationResponse{}
	err := pc.Authenticate(pc.GetConsumerKey(), AuthNResp)

	if err != nil {
		msg = pocket.CONSUMERKEYNOTVALIDEn
		return msg
	}

	err = pc.UserAuthorise(pocket.UserAuthorisationURL, AuthNResp.Code, pocket.RedirectURI)
	if err != nil {
		msg = pocket.ERRORAPPROVINGLINTEn
		return msg
	}

	fmt.Print(pocket.ACKNOWLEDGEAUTHen)
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	AuthRResp := &pocket.AuthorisationResponse{}
	err = pc.RetrieveAccessToken(pc.GetConsumerKey(), AuthNResp.Code, AuthRResp)

	if err != nil {
		msg = pocket.ERRORAUTHen
		return msg
	}

	cfgval := fmt.Sprintf("ConsumerKey: %v\nAccessToken: %v\nUsername: %v",
		pc.GetConsumerKey(), AuthRResp.AccessToken, AuthRResp.Username)

	err = ioutil.WriteFile(pocket.CfgFile, []byte(cfgval), 0644)

	if err != nil {
		msg = pocket.ERRORSAVINGCONSUMERKEYen
		return msg
	}

	msg = pocket.AUTHSUCCESSen

	return msg
}
