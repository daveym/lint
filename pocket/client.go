package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
)

// API - Base interface type for Pocket API. Allows us to mock/test.
type API interface {
	SetConsumerKey(string)
	GetConsumerKey() string
	SetAccessToken(string)
	GetAccessToken() string
	Authenticate(string, interface{}) error
	UserAuthorise(string, string, string) error
	RetrieveAccessToken(string, string, interface{}) error
	Retrieve(RetrieveRequest, interface{}) error
}

// Client - Provide access the Pocket API
type Client struct {
	_consumerKey string
	_accessToken string
}

// SetConsumerKey - Set the new consumer Key
func (p *Client) SetConsumerKey(newKey string) {
	p._consumerKey = newKey
}

// GetConsumerKey - Set the new consumer Key
func (p *Client) GetConsumerKey() string {
	return p._consumerKey
}

// SetAccessToken - Set the new consumer Key
func (p *Client) SetAccessToken(newToken string) {
	p._accessToken = newToken
}

// GetAccessToken - Set the new consumer Key
func (p *Client) GetAccessToken() string {
	return p._accessToken
}

// Authenticate takes the the users consumer key and performs a one time authentication with
// the Pocket API to request access. A Request Token is returned that should be used for all
// subsequent requests to Pocket.
func (p *Client) Authenticate(consumerKey string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": RedirectURI}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthenticationURL, jsonStr, resp)

	return err
}

// UserAuthorise -  Redirect the user to the Pocket Authorise screen
func (p *Client) UserAuthorise(url string, reqtoken string, uri string) error {

	browser := exec.Command("open", url+
		"request_token="+reqtoken+
		"&redirect_uri="+uri)

	_, err := browser.Output()

	return err
}

// RetrieveAccessToken -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func (p *Client) RetrieveAccessToken(consumerKey string, code string, resp interface{}) error {

	request := map[string]string{"consumer_key": consumerKey, "code": code}
	jsonStr, _ := json.Marshal(request)
	err := postJSON("POST", AuthorisationURL, jsonStr, resp)

	return err
}

// Retrieve -  Pull back items from Pocket
func (p *Client) Retrieve(itemreq RetrieveRequest, resp interface{}) error {

	jsonStr, err := json.Marshal(itemreq)

	err = postJSON("GET", RetrieveURL, jsonStr, resp)

	return err
}

// Generic post method, url and data are incoming. Response is a  base interface
// that we can use to return many response types.
func postJSON(action string, url string, data []byte, resp interface{}) (err error) {

	req, err := http.NewRequest(action, url, bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF8")
	req.Header.Set("X-Accept", "application/json")

	client := &http.Client{}
	jsonResp, err := client.Do(req)

	if err != nil {
		fmt.Println(err.Error())
	}

	err = json.NewDecoder(jsonResp.Body).Decode(resp)
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
