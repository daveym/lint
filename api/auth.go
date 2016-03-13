package api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/daveym/lint/pocket"
	"github.com/spf13/viper"
)

// Authenticate agaimst Pockets API
func Authenticate(pc pocket.API) {

	if viper.GetBool("AccessToken") {
		println("Already authenticated - Access Token present in lint.yaml")
		return
	}

	AuthNResp := &pocket.AuthenticationResponse{}
	err := pc.Authenticate(pocket.ConsumerKey, AuthNResp)

	if err != nil {
		fmt.Println("Please check your consumer key, it does not appear to be valid.")
		return
	}

	browser := exec.Command("open", pocket.UserAuthorisationURL+
		"request_token="+AuthNResp.Code+
		"&redirect_uri="+pocket.RedirectURI)

	_, err = browser.Output()

	if err != nil {
		println("Error opening authentication page. Please check your connectivity/default browser.")
		return
	}

	fmt.Println("Press ENTER when you have authorised Lint to access to Pocket.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	AuthRResp := &pocket.AuthorisationResponse{}
	err = pc.Authorise(pocket.ConsumerKey, AuthNResp.Code, AuthRResp)

	if err != nil {
		fmt.Println("Error authorising your consumer key and request token. Have you granted permission to Lint?")
	}

	cfgval := fmt.Sprintf("AccessToken: %v\nUsername: %v", AuthRResp.AccessToken, AuthRResp.Username)
	err = ioutil.WriteFile(pocket.CfgFile, []byte(cfgval), 0644)

	viper.Set("AccessToken", AuthRResp.AccessToken)
	viper.Set("Username", AuthRResp.Username)

	println("Authentication Successful - Access Token is persisted to lint.yaml")

}
