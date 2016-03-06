package Lint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Initial Authintication response from Pocket
type authenticationResponse struct {
	Code  string
	State string
}

// Initial Authorisation response from Pocket
type authorisationResponse struct {
	AccessToken string
	Username    string
}

const (
	authenticationURL string = "https://getpocket.com/v3/oauth/request"
	authorisationURL  string = "https://getpocket.com/v3/oauth/authorize"
	redirectURI       string = "https://github.com/daveym/pocket/blob/master/AUTHCOMPLETE.md"
)

// Authenticate takes the the users consumer key and performs a one time authentication with the Pocket API to request access.
// A Request Token is returned that should be used for all subsequent requests to Pocket.
func Authenticate(consumerKey string) string {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": redirectURI}
	jsonStr, _ := json.Marshal(request)

	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", authenticationURL, bytes.NewBuffer(jsonStr))
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
	fmt.Println("AuthN response Body:", string(body))

	var r = new(authenticationResponse)
	err = json.Unmarshal([]byte(body), &r)

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	return r.Code

}

// Authorise -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func Authorise(consumerKey string, code string) (string, string) {

	request := map[string]string{"consumer_key": consumerKey, "code": code}
	jsonStr, _ := json.Marshal(request)

	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", authorisationURL, bytes.NewBuffer(jsonStr))
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
	fmt.Println("AuthR response Body:", string(body))

	var r = new(authorisationResponse)
	err = json.Unmarshal([]byte(body), &r)

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	return r.AccessToken, r.Username

}
