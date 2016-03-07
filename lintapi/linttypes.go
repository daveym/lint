package Lint

import "time"

// CONSTANTS
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

const (
	// ItemStatusUnread - Pocket Item has not been read
	ItemStatusUnread ItemStatus = 0
	// ItemStatusArchived - Pocket Item has been archived
	ItemStatusArchived = 1
	// ItemStatusDeleted - Pocket Item has been deleted
	ItemStatusDeleted = 2
)

const (
	// ItemMediaAttachmentNoMedia - No media attached to the Pocket Item
	ItemMediaAttachmentNoMedia ItemMediaAttachment = 0
	// ItemMediaAttachmentHasMedia - Media is attached to the Pocket Item
	ItemMediaAttachmentHasMedia = 1
	// ItemMediaAttachmentIsMedia - The Pocket Item is media only
	ItemMediaAttachmentIsMedia = 2
)

// TYPES

// Time - alias for time.Time
type Time time.Time

// ItemStatus - Status of the Pocket Item, defined by ItemStatus constants
type ItemStatus int

// ItemMediaAttachment - Is Media attached to the Pocket Item. Defined by ItemMedia constants
type ItemMediaAttachment int

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
	Since       Time
	Count       int
	Offset      int
}

type listItem struct {
	ItemID        int        `json:"item_id,string"`
	ResolvedID    int        `json:"resolved_id,string"`
	GivenURL      string     `json:"given_url"`
	ResolvedURL   string     `json:"resolved_url"`
	GivenTitle    string     `json:"given_title"`
	ResolvedTitle string     `json:"resolved_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	Excerpt       string
	IsArticle     int                 `json:"is_article,string"`
	HasImage      ItemMediaAttachment `json:"has_image,string"`
	HasVideo      ItemMediaAttachment `json:"has_video,string"`
	WordCount     int                 `json:"word_count,string"`
	Tags          map[string]map[string]interface{}
	Authors       map[string]map[string]interface{}
	Images        map[string]map[string]interface{}
	Videos        map[string]map[string]interface{}
	SortID        int  `json:"sort_id"`
	TimeAdded     Time `json:"time_added"`
	TimeUpdated   Time `json:"time_updated"`
	TimeRead      Time `json:"time_read"`
	TimeFavorited Time `json:"time_favorited"`
}
