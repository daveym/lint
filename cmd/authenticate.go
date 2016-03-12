package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/daveym/lint/lintapi"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var authenticateCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Pocket using your consumer key. One off activity.",
	Long: `To access the Pocket API, you need to authenticate with your consumer key. Your consumer key 
	can be found under the development area within the pocket website`,
	Run: func(cmd *cobra.Command, args []string) {

		fmt.Println(viper.Get("AccessToken"))

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

		println(AuthRResp.AccessToken, AuthRResp.Username)
	}}

func init() {
	RootCmd.AddCommand(authenticateCmd)
}
