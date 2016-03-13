package linttests

type MockClient struct{}

// Authenticate takes the the users consumer key and performs a one time authentication with
// the Pocket API to request access. A Request Token is returned that should be used for all
// subsequent requests to Pocket.
func (p *MockClient) Authenticate(consumerKey string, resp interface{}) error {

	var err error
	return err
}

// Authorise -  Using the consumerKey and request code, obtain an Access token and Pocket Username
func (p *MockClient) Authorise(consumerKey string, code string, resp interface{}) error {
	var err error
	return err
}

// Retrieve -  Pull back items from Pocket
func (p *MockClient) Retrieve(itemRequest ItemRequest, resp interface{}) error {
	var err error
	return err
}
