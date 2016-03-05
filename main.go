/*
Usage:
    pocket auth
    pocket count
    pocket list
    pocket archive [comma seoperated ids]
    pocket delete  [comma seoperated ids]
    pocket --version

Commands:
    auth		      Get an access token for batch work.
    count             Count the number of items in your Pocket list.
	list			  List the articles held in pocket
*/

package main

import (
	"flag"
	"fmt"

	"github.com/daveym/pocket/pocket"
)

func main() {

	consumerKey := flag.String("auth", "No key set", "Please enter your pocket consumer_key.")
	flag.Parse()

	fmt.Println(*consumerKey)
	requestToken := Pocket.Authenticate(consumerKey)

	fmt.Println("go to https://getpocket.com/auth/authorize?request_token=" + requestToken)
	fmt.Println("type ENTER when the application is authorized")

}
