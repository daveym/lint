/*
Usage:
    lint auth
    lint count
    lint list
    lint archive [comma separated ids]
    lint delete  [comma separated ids]
    lint --version

Commands:
    auth		      Get an access token for batch work.
    count             Count the number of items in your Pocket list.
	list			  List the articles held in pocket
*/

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"github.com/daveym/lint/lintapi"
)

func main() {

	consumerKey := flag.String("auth", "No key set", "Please enter your pocket consumer_key.")
	flag.Parse()

	// TEST KEY
	*consumerKey = "52204-09dc66720db59e4959bf47f5"

	if *consumerKey != "" {
		fmt.Println(*consumerKey)
		requestToken := Lint.Authenticate(*consumerKey)
		fmt.Println("go to https://getpocket.com/auth/authorize?request_token=" + requestToken)
		fmt.Println("Please go to:" + Lint.UserAuthorisationURL + "request_token=" + requestToken + "&redirect_uri=" + Lint.RedirectURI)
		fmt.Println("type ENTER when the application is authorized")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		accessToken, username := Lint.Authorise(*consumerKey, requestToken)
		fmt.Println(accessToken, username)
	}
}
