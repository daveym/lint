package Api

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/daveym/lint/Pocket"
	"github.com/spf13/viper"
)

// Authenticate agaimst Pockets API
func Authenticate() {

	if viper.GetBool("AccessToken") {
		println("Already authenticated - Access Token present in lint.yaml")
		return
	}

	pc := Pocket.Client{}

	AuthNResp := &Pocket.AuthenticationResponse{}
	err := pc.Authenticate(Pocket.ConsumerKey, AuthNResp)

	if err != nil {
		fmt.Println("Please check your consumer key, it does not appear to be valid.")
		return
	}

	browser := exec.Command("open", Pocket.UserAuthorisationURL+
		"request_token="+AuthNResp.Code+
		"&redirect_uri="+Pocket.RedirectURI)

	_, err = browser.Output()

	if err != nil {
		println("Error opening authentication page. Please check your connectivity/default browser.")
		return
	}

	fmt.Println("Press ENTER when you have authorised Lint to access to Pocket.")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	AuthRResp := &Pocket.AuthorisationResponse{}
	err = pc.Authorise(Pocket.ConsumerKey, AuthNResp.Code, AuthRResp)

	if err != nil {
		fmt.Println("Error authorising your consumer key and request token. Have you granted permission to Lint?")
	}

	cfgval := fmt.Sprintf("AccessToken: %v\nUsername: %v", AuthRResp.AccessToken, AuthRResp.Username)
	err = ioutil.WriteFile(Pocket.CfgFile, []byte(cfgval), 0644)

	viper.Set("AccessToken", AuthRResp.AccessToken)
	viper.Set("Username", AuthRResp.Username)

	println("Authentication Successful - Access Token is persisted to lint.yaml")

}
