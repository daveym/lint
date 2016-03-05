package Pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Initial Auth response from Pocket
type authResponse struct {
	Code  string
	State string
}

const (
	url         string = "https://getpocket.com/v3/oauth/request"
	redirectURI string = "https://github.com/daveym/pocket/blob/master/AUTHCOMPLETE.md"
)

// Authenticate takes the the users consumer key and performs a one time authentication with the Pocket API to request access.
// A Request Token is returned that should be used for all subsequent requests to Pocket.
func Authenticate(consumerKey string) string {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": redirectURI}
	jsonStr, _ := json.Marshal(request)

	fmt.Println(string(jsonStr))
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

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	var s = new(authResponse)
	err = json.Unmarshal([]byte(body), &s)

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	return string(s.Code)

}

// Used by Authenticate - take a consumer key and return a request token.
func getRequestToken(consumerKey *string) {

}
