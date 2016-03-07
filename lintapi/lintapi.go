package Lint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Authenticate takes the the users consumer key and performs a one time authentication with the Pocket API to request access.
// A Request Token is returned that should be used for all subsequent requests to Pocket.
func Authenticate(consumerKey string) string {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": RedirectURI}
	jsonStr, _ := json.Marshal(request)

	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("POST", AuthenticationURL, bytes.NewBuffer(jsonStr))
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

	fmt.Println("BODY:" + string(body))
	var r = new(AuthenticationResponse)
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
	req, err := http.NewRequest("POST", AuthorisationURL, bytes.NewBuffer(jsonStr))
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
	fmt.Println("BODY:" + string(body))

	var r = new(AuthorisationResponse)
	err = json.Unmarshal([]byte(body), &r)

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	return r.AccessToken, r.Username
}

// GetItems -  Pull back items from Pocket
func GetItems(consumerKey string, accessToken string, itemRequest ItemRequest) *ItemResponse {

	request := map[string]string{"consumer_key": consumerKey, "access_token": accessToken}
	jsonStr, _ := json.Marshal(request)

	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("GET", RetrieveURL, bytes.NewBuffer(jsonStr))
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
	fmt.Println("GetItems response Body:", string(body))

	var r = new(ItemResponse)
	err = json.Unmarshal([]byte(body), &r)

	if err != nil {
		fmt.Printf("%T\n%s\n%#v\n", err, err, err)
	}

	return r

}
