package Lint

import "time"

const (
	// AuthenticationURL - API address to Authenticate Pocket Consumer Key
	AuthenticationURL string = "https://getpocket.com/v3/oauth/request"

	// AuthorisationURL - API address to Authorise and recience a Request Key
	AuthorisationURL string = "https://getpocket.com/v3/oauth/authorize"

	//UserAuthorisationURL - Address that a user must enter into their browser to Authorise Lint to access Pocket
	UserAuthorisationURL string = "https://getpocket.com/auth/authorize?"

	//RedirectURI - Link back location after Authorisation has been granted
	RedirectURI string = "https://github.com/daveym/lint/blob/master/AUTHCOMPLETE.md"
)

// Initial Authentication response from Pocket
type authenticationResponse struct {
	Code  string
	State string
}

// Initial Authorisation response from Pocket
type authorisationResponse struct {
	AccessToken string
	Username    string
}

type getListRequest struct {
	State       string
	Favorite    int
	Tag         string
	ContentType string
	Sort        string
	DetailType  string
	Search      string
	Domain      string
	Since       time.Time
	Count       int
	Offset      int
}
