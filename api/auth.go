package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/daveym/lint/pocket"
	"github.com/spf13/viper"
)

// Authenticate against Pockets API
func Authenticate(pc pocket.API) string {

	msg := ""

	if len(viper.GetString("AccessToken")) > 0 {
		msg = "Already authenticated - Access Token present in lint.yaml"
		return msg
	}

	AuthNResp := &pocket.AuthenticationResponse{}
	err := pc.Authenticate(pocket.ConsumerKey, AuthNResp)

	if err != nil {
		msg = "Please check your consumer key, it does not appear to be valid."
		return msg
	}

	err = pc.AuthoriseUse(pocket.AuthorisationURL, AuthNResp.Code, pocket.RedirectURI)
	if err != nil {
		msg = "Error opening authentication page. Please check your connectivity/default browser."
		return msg
	}

	fmt.Print("Press ENTER when you have authorised Lint to access to Pocket.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	AuthRResp := &pocket.AuthorisationResponse{}
	err = pc.GetAccessToken(pocket.ConsumerKey, AuthNResp.Code, AuthRResp)

	if err != nil {
		msg = "Error authorising your consumer key and request token. Have you granted permission to Lint?"
		return msg
	}

	cfgval := fmt.Sprintf("AccessToken: %v\nUsername: %v", AuthRResp.AccessToken, AuthRResp.Username)
	err = ioutil.WriteFile(pocket.CfgFile, []byte(cfgval), 0644)

	viper.Set("AccessToken", AuthRResp.AccessToken)
	viper.Set("Username", AuthRResp.Username)

	msg = "Authentication Successful - Access Token is persisted to lint.yaml"
	return msg
}
