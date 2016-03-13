package pocket

// MockClient - Used for mocking
type MockClient struct{}

// Authenticate - Mock instance
func (p *MockClient) Authenticate(consumerKey string, resp interface{}) error {

	var err error
	return err
}

// AuthoriseUse - Mock instance
func (p *MockClient) AuthoriseUse(url string, code string, uri string) error {

	var err error
	return err
}

// GetAccessToken -  Mock instance
func (p *MockClient) GetAccessToken(consumerKey string, code string, resp interface{}) error {
	var err error
	return err
}

// Retrieve -  Mock instance
func (p *MockClient) Retrieve(itemRequest ItemRequest, resp interface{}) error {
	var err error
	return err
}
