package Pocket

import "time"

// State of returned items
type State string

// Sort - Sort options for returned items
type Sort string

// ItemStatus - Status of retrieved item
type ItemStatus int

// ContentType - Article video or image
type ContentType string

// DetailType - Simple or complete detail
type DetailType string

// FavoriteFilter - Filter by favourite
type FavoriteFilter string

// ItemMediaAttachment - Is Media attached to the Pocket Item. Defined by ItemMedia constants
type ItemMediaAttachment int

// RequestToken - Obtain a code from Pocket
type RequestToken struct {
	Code string `json:"code"`
}

// AuthenticationResponse response from Pocket
type AuthenticationResponse struct {
	Code  string
	State string
}

// AuthorisationResponse response from Pocket
type AuthorisationResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

// RetrieveRequest - make a request to Pocket for specific items
type RetrieveRequest struct {
	ConsumerKey string         `json:"consumer_key"`
	AccessToken string         `json:"access_token"`
	State       State          `json:"state,omitempty"`
	Favorite    FavoriteFilter `json:"favorite,omitempty"`
	Tag         string         `json:"tag,omitempty"`
	ContentType ContentType    `json:"contentType,omitempty"`
	Sort        Sort           `json:"sort,omitempty"`
	DetailType  DetailType     `json:"detailType,omitempty"`
	Search      string         `json:"search,omitempty"`
	Domain      string         `json:"domain,omitempty"`
	Since       int            `json:"since,omitempty"`
	Count       int            `json:"count,omitempty"`
	Offset      int            `json:"offset,omitempty"`
}

// RetrieveResponse - List of retrieved items from Pocket
type RetrieveResponse struct {
	List     map[string]ItemResponse
	Status   int
	Complete int
	Since    int
}

// ItemRequest - Query pocket for a list of items
type ItemRequest struct {
	ConsumerKey string
	AccessToken string `json:"access_token"`
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

// ItemResponse - An individual Pocket entry
type ItemResponse struct {
	ItemID        int        `json:"item_id"`
	ResolvedID    int        `json:"resolved_id"`
	GivenURL      string     `json:"given_url"`
	ResolvedURL   string     `json:"resolved_url"`
	GivenTitle    string     `json:"given_title"`
	ResolvedTitle string     `json:"resolved_title"`
	Favorite      int        `json:",string"`
	Status        ItemStatus `json:",string"`
	Excerpt       string
	IsArticle     int                 `json:"is_article"`
	HasImage      ItemMediaAttachment `json:"has_image"`
	HasVideo      ItemMediaAttachment `json:"has_video"`
	WordCount     int                 `json:"word_count"`
	SortID        int                 `json:"sort_id"`
	TimeAdded     time.Time           `json:"time_added"`
	TimeUpdated   time.Time           `json:"time_updated"`
	TimeRead      time.Time           `json:"time_read"`
	TimeFavorited time.Time           `json:"time_favorited"`
}
