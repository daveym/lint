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
)

func main() {

	consumerKey := flag.String("auth", "No key set", "")
	flag.Parse()

	fmt.Println(*consumerKey)
	request_token := getRequestToken(*consumerKey)

	fmt.Println("go to https://getpocket.com/auth/authorize?request_token=" + request_token)
	fmt.Println("type ENTER when the application is authorized")

}
