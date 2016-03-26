package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daveym/lint/pocket"
	"github.com/spf13/viper"
)

// Authenticate against Pocket API. Interface used to allow mock to be passed in.
func Authenticate(pc pocket.API) string {

	msg := ""

	if len(pc.GetConsumerKey()) == 0 {
		msg = "Consumer Key is not present in lint.yaml. Please add, using the format 'ConsumerKey: value', without quotes."
		return msg
	}

	AuthNResp := &pocket.AuthenticationResponse{}
	err := pc.Authenticate(pc.GetConsumerKey(), AuthNResp)

	if err != nil {
		msg = "Please check your consumer key, it does not appear to be valid."
		return msg
	}

	err = pc.UserAuthorise(pocket.UserAuthorisationURL, AuthNResp.Code, pocket.RedirectURI)
	if err != nil {
		msg = "Error whilst approving Lint access to Pocket data. Please check your connectivity/default browser."
		return msg
	}

	fmt.Print("Press ENTER when you have authorised Lint to access to Pocket.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	AuthRResp := &pocket.AuthorisationResponse{}
	err = pc.RetrieveAccessToken(pc.GetConsumerKey(), AuthNResp.Code, AuthRResp)

	if err != nil {
		msg = "Error authorising your consumer key and request token. Have you granted permission to Lint?"
		return msg
	}

	cfgval := fmt.Sprintf("ConsumerKey: %v\nAccessToken: %v\nUsername: %v", pc.GetConsumerKey(), AuthRResp.AccessToken, AuthRResp.Username)
	err = ioutil.WriteFile(pocket.CfgFile, []byte(cfgval), 0644)

	if err != nil {
		msg = "Error persisting consumer key, access token and username to lint.yaml"
		return msg
	}

	viper.Set("ConsumerKey", pc.GetConsumerKey())
	viper.Set("AccessToken", pc.GetAccessToken())
	viper.Set("Username", AuthRResp.Username)

	msg = "Authentication Successful - Pocket Access Token is persisted to lint.yaml"

	return msg
}
