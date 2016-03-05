package Pocket

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	url string = "https://getpocket.com/v3/oauth/request"
)

// Authenticate takes the the users consumer key and performs a one time authentication with the Pocket API to request access.
// A Request Token is returned that should be used for all subsequent requests to Pocket.
func Authenticate(consumerKey string) string {

	var jsonStr = []byte("{consumer_key:" + consumerKey + "redirect_uri:pocketapp1234:authorizationFinished}")
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return string(body)

}

// Used by Authenticate - take a consumer key and return a request token.
func getRequestToken(consumerKey *string) {

}
