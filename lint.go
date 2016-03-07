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
	"fmt"
	"os"

	"github.com/daveym/lint/lintapi"
)

func main() {

	consumerKey := ""

	if consumerKey != "" {
		fmt.Println(consumerKey)
		requestToken := Lint.Authenticate(consumerKey)
		fmt.Println("Please go to:" + Lint.UserAuthorisationURL + "request_token=" + requestToken + "&redirect_uri=" + Lint.RedirectURI)
		fmt.Println("")
		fmt.Println("Press ENTER when you have authorised the application to use Lint.")
		bufio.NewReader(os.Stdin).ReadBytes('\n')

		accessToken, username := Lint.Authorise(consumerKey, requestToken)
		fmt.Println("ACCESS TOKEN: " + accessToken + " USERNAME:" + username)

		// Test Retrieve
		//var ItemRequest Lint.ItemRequest
		//ItemRequest.Count = 10

		//items := Lint.GetItems(consumerKey, accessToken, ItemRequest)
		//fmt.Println(items)

	} else {
		fmt.Println(consumerKey)
		fmt.Println("lint.json missing. Please add this file, with the following entries:")
		fmt.Println("{'consumer_key:', 'Your Consumer Key'}")
	}
}
