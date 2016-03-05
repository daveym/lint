/*
Usage:
    lint auth
    lint count
    lint list
    lint archive [comma seoperated ids]
    lint delete  [comma seoperated ids]
    lint --version

Commands:
    auth		      Get an access token for batch work.
    count             Count the number of items in your Pocket list.
	list			  List the articles held in pocket
*/

package main

import (
	"flag"
	"fmt"

	"github.com/daveym/lint/lint"
)

func main() {

	consumerKey := flag.String("auth", "No key set", "Please enter your pocket consumer_key.")
	flag.Parse()

	// TEST
	*consumerKey = "52204-09dc66720db59e4959bf47f5"

	if *consumerKey != "" {
		fmt.Println(*consumerKey)
		requestToken := Pocket.Authenticate(*consumerKey)
		fmt.Println("go to https://getpocket.com/auth/authorize?request_token=" + requestToken)
		fmt.Println("type ENTER when the application is authorized")
	}
}
