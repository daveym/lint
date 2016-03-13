package pocket

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os/exec"
)

// API - Base interface type for Pocket API. Allows us to mock/test.
type API interface {
	Authenticate(string, interface{}) error
	AuthoriseUse(string, string, string) error
	GetAccessToken(string, string, interface{}) error
	Retrieve(ItemRequest, interface{}) error
}

// Client - Provide access the Pocket API
type Client struct{}

// Authenticate takes the the users consumer key and performs a one time authentication with
// the Pocket API to request access. A Request Token is returned that should be used for all
// subsequent requests to Pocket.
func (p *Client) Authenticate(consumerKey string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": RedirectURI}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthenticationURL, jsonStr, resp)

	return err
}

// AuthoriseUse -  Redirect the user to the Pocket Authorise screen
func (p *Client) AuthoriseUse(url string, code string, uri string) error {

	browser := exec.Command("open", url+
		"request_token="+code+
		"&redirect_uri="+uri)

	_, err := browser.Output()

	return err
}

// GetAccessToken -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func (p *Client) GetAccessToken(consumerKey string, code string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "code": code}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthorisationURL, jsonStr, resp)

	return err
}

// Retrieve -  Pull back items from Pocket
func (p *Client) Retrieve(itemRequest ItemRequest, resp interface{}) error {

	jsonStr, _ := json.Marshal(itemRequest)
	err := postJSON("GET", RetrieveURL, jsonStr, resp)

	return err
}

// Generic post method, url and data are incoming. Response is a  base interface
// that we can use to that we can use to return many reponses types.
func postJSON(action string, url string, data []byte, resp interface{}) (err error) {

	req, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	jsonResp, err := client.Do(req)

	return json.NewDecoder(jsonResp.Body).Decode(resp)
}
