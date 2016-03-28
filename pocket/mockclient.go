package pocket

import "errors"

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
func (p *MockClient) Authenticate(consumerKey string, resp *AuthenticationResponse) error {

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
func (p *MockClient) RetrieveAccessToken(consumerKey string, code string, resp *AuthorisationResponse) error {

	var err error
	if consumerKey == "FAIL" {
		err = errors.New("Invalid Key")
		return err
	}

	return nil
}

// Retrieve -  Mock instance
func (p *MockClient) Retrieve(req RetrieveRequest, resp *RetrieveResponse) error {

	var err error

	resp.Status = 123
	resp.Complete = 123
	resp.List = make(map[string]Item)
	resp.Since = 789

	fakeItem := Item{
		Excerpt:       "Excerpt",
		Favorite:      1,
		GivenTitle:    "Docker",
		GivenURL:      "http://docker.com",
		HasImage:      ItemMediaAttachmentNoMedia,
		HasVideo:      ItemMediaAttachmentNoMedia,
		IsArticle:     1,
		ItemID:        11111,
		ResolvedID:    11111,
		ResolvedTitle: "Docker",
		ResolvedURL:   "http://docker.com",
		SortID:        11111,
		Status:        ItemStatusUnread,
		WordCount:     150}

	if req.Search == "docker" {
		resp.List["11111"] = fakeItem
	}

	if req.Search == "nothing" {
		resp.Status = 0
		resp.Complete = 0
		resp.List = make(map[string]Item)
		resp.Since = 0
	}

	return err
}

// Modify -  Modify items in Pocket
func (p *MockClient) Modify(req ModifyRequest, resp *ModifyResponse) error {

	actions := req.Actions
	itemVal := actions[0].ItemID

	if itemVal == 12345 {
		resp.Status = 1
	}

	if itemVal == 45678 {
		resp.Status = 0
	}

	var err error
	return err
}
