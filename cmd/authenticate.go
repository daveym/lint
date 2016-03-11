// Copyright © 2016 davey mcglade  <davey.mcglade@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"

	"github.com/daveym/lint/lintapi"
	"github.com/spf13/cobra"
)

// ConsumerKey - Private key used to access Pocket API.
var ConsumerKey = ""

// AccessToken - Access token  used to access Pocket API (per request)
var AccessToken = ""

var authenticateCmd = &cobra.Command{
	Use:   "auth",
	Short: "Authenticate against Pocket using your consumer key. One off activity.",
	Long: `To access the Pocket API, you need to authenticate with your consumer key. Your consumer key 
	can be found under the development area within the pocket website`,
	Run: func(cmd *cobra.Command, args []string) {

		if len(args) > 0 {

			ConsumerKey = args[0]

			client := LintApi.New(ConsumerKey, AccessToken)
			requestToken, err := LintApi.Authenticate(client.ConsumerKey)

			if err != nil {
				fmt.Println("Please check your consumer key, it does not appear to be valid.")
			} else {

				cmd := exec.Command("open", LintApi.UserAuthorisationURL+"request_token="+requestToken+"&redirect_uri="+LintApi.RedirectURI)
				_, err := cmd.Output()

				if err != nil {
					println(err.Error())
					return
				}

				fmt.Println("")
				fmt.Println("Please authorise Lint from within your browser, or alternatively copy and paste the following link if the authentication page has not been displayed.")
				fmt.Println(LintApi.UserAuthorisationURL + "request_token=" + requestToken + "&redirect_uri=" + LintApi.RedirectURI)
				fmt.Println("")
				fmt.Println("and press ENTER when you have authorised the application to use Lint.")
				bufio.NewReader(os.Stdin).ReadBytes('\n')

				AccessToken, _, err := LintApi.Authorise(ConsumerKey, requestToken)

				if err != nil {
					fmt.Println("Error authorising your consumer key and request token")
				} else {
					println(AccessToken)
				}

			}

		} else {
			fmt.Println("Consumer Key missing - please obtain this from the Pocket Developer site.")
		}
	}}

func init() {
	RootCmd.AddCommand(authenticateCmd)
}
