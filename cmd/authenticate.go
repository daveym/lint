package cmd

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/daveym/lint/lintapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authenticateCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Pocket  to get an Access Token. One off activity.",
	Long: `To access the Pocket API, you need to authenticate with your consumer key. Your consumer key 
	can be found under the development area within the pocket website`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(viper.GetString("AccessToken")) == 0 {

			AuthNResp := &LintApi.AuthenticationResponse{}
			err := LintApi.Authenticate(LintApi.ConsumerKey, AuthNResp)

			if err != nil {
				fmt.Println("Please check your consumer key, it does not appear to be valid.")
				return
			}

			browser := exec.Command("open", LintApi.UserAuthorisationURL+
				"request_token="+AuthNResp.Code+
				"&redirect_uri="+LintApi.RedirectURI)

			_, err = browser.Output()

			if err != nil {
				println("Error opening authentication page. Please check your connectivity/default browser.")
				return
			}

			fmt.Println("Press ENTER when you have authorised Lint to access to Pocket.")
			bufio.NewReader(os.Stdin).ReadBytes('\n')

			AuthRResp := &LintApi.AuthorisationResponse{}
			err = LintApi.Authorise(LintApi.ConsumerKey, AuthNResp.Code, AuthRResp)

			if err != nil {
				fmt.Println("Error authorising your consumer key and request token. Have you granted permission to Lint?")
			}

			cfgval := fmt.Sprintf("AccessToken: %v\nUsername: %v", AuthRResp.AccessToken, AuthRResp.Username)
			err = ioutil.WriteFile(LintApi.CfgFile, []byte(cfgval), 0644)

			viper.Set("AccessToken", AuthRResp.AccessToken)
			viper.Set("Username", AuthRResp.Username)

			println("Authentication Successful - Access Token is persisted to lint.yaml")

		} else {
			println("Already authenticated - Access Token present in lint.yaml")
		}
	}}

func init() {
	RootCmd.AddCommand(authenticateCmd)
}
