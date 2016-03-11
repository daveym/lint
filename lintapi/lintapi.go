// Copyright © 2016 davey mcglade  <davey.mcglade@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package LintApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// New creates a new Pocket client.
func New(consumerKey string, accessToken string) *PocketClient {

	return &PocketClient{ConsumerKey: consumerKey, AccessToken: accessToken}
}

// Authenticate takes the the users consumer key and performs a one time authentication with the Pocket API to request access.
// A Request Token is returned that should be used for all subsequent requests to Pocket.
func Authenticate(consumerKey string) (string, error) {

	request := map[string]string{"consumer_key": consumerKey, "redirect_uri": RedirectURI}
	jsonStr, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", AuthenticationURL, bytes.NewBuffer(jsonStr))
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

	var r = new(AuthenticationResponse)
	err = json.Unmarshal([]byte(body), &r)

	return r.Code, err

}

// Authorise -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func Authorise(consumerKey string, code string) (string, string, error) {

	request := map[string]string{"consumer_key": consumerKey, "code": code}
	jsonStr, _ := json.Marshal(request)

	req, err := http.NewRequest("POST", AuthorisationURL, bytes.NewBuffer(jsonStr))
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

	var r = new(AuthorisationResponse)
	err = json.Unmarshal([]byte(body), &r)

	return r.AccessToken, r.Username, err
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
