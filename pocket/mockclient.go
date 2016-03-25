package pocket

import (
	"errors"
	"fmt"
)

// MockClient - Used for mocking
type MockClient struct {
	_consumerKey string
	_accessToken string
}

// SetConsumerKey - Set the new consumer Key
func (p *MockClient) SetConsumerKey(newKey string) {
	p._consumerKey = newKey
}

// GetConsumerKey - Set the new consumer Key
func (p *MockClient) GetConsumerKey() string {
	return p._consumerKey
}

// SetAccessToken - Set the new access token
func (p *MockClient) SetAccessToken(newToken string) {
	p._accessToken = newToken
}

// GetAccessToken - Set the new access token
func (p *MockClient) GetAccessToken() string {
	return p._accessToken
}

// Authenticate - Mock instance
func (p *MockClient) Authenticate(consumerKey string, resp interface{}) error {

	var err error

	if consumerKey == "INVALIDKEY" {
		err = errors.New("Invalid Key")
	}
	return err
}

// UserAuthorise - Mock instance
func (p *MockClient) UserAuthorise(url string, code string, uri string) error {

	var err error

	if p.GetConsumerKey() == "INVALIDBROWSER" {
		err = errors.New("Invalid Key")
	}
	return err
}

// RetrieveAccessToken -  Mock instance
func (p *MockClient) RetrieveAccessToken(consumerKey string, code string, resp interface{}) error {

	var err error
	if consumerKey == "FAIL" {
		err = errors.New("Invalid Key")
		return err
	}

	return nil
}

// Retrieve -  Mock instance
func (p *MockClient) Retrieve(req RetrieveRequest, resp interface{}) error {

	var err error

	fakeResp := RetrieveResponse{
		Status:   123,
		Complete: 456,
		List:     make(map[string]Item),
		Since:    789}

	fakeItem := Item{
		Excerpt:       "A sample test docker article",
		Favorite:      1,
		GivenTitle:    "Docket Test 1",
		GivenURL:      "http://dockettest1.com",
		HasImage:      ItemMediaAttachmentNoMedia,
		HasVideo:      ItemMediaAttachmentNoMedia,
		IsArticle:     1,
		ItemID:        11111,
		ResolvedID:    11111,
		ResolvedTitle: "Docket Test 1",
		ResolvedURL:   "http://dockettest1.com",
		SortID:        11111,
		Status:        ItemStatusUnread,
		WordCount:     150}

	if req.Search == "docker" {

		fakeResp.List["11111"] = fakeItem
		resp = fakeResp
	}

	if req.Search == "nothing" {

		fakeResp := RetrieveResponse{
			Status:   0,
			Complete: 0,
			List:     make(map[string]Item),
			Since:    0}

		resp = fakeResp

		fmt.Println("")
		fmt.Println("MOCK RESPONSE:")
		fmt.Println(resp)
	}

	return err

}
