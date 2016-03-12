package LintApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func postJSON(action string, url string, data []byte, res interface{}) (err error) {

	req, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	return json.NewDecoder(resp.Body).Decode(res)
}

// Authenticate takes the the users consumer key and performs a one time authentication with
// the Pocket API to request access. A Request Token is returned that should be used for all
//  subsequent requests to Pocket.
func Authenticate(consumerKey string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": RedirectURI}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthenticationURL, jsonStr, resp)

	return err
}

// Authorise -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func Authorise(consumerKey string, code string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "code": code}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthorisationURL, jsonStr, resp)

	return err
}

// GetItems -  Pull back items from Pocket
func GetItems(itemRequest ItemRequest) (ItemResponse, error) {

	jsonStr, _ := json.Marshal(itemRequest)

	fmt.Println(string(jsonStr))
	req, err := http.NewRequest("GET", RetrieveURL, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("GetItems response Body:", string(body))

	var r = new(ItemResponse)
	err = json.Unmarshal([]byte(body), &r)

	return *r, err
}
