package pocket

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
	Code  string // Code is request Token
	State string
}

// AuthorisationResponse response from Pocket
type AuthorisationResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

// RetrieveRequest - response from Pocket
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

// RetrieveResponse - List of items retrieved from Pocket
type RetrieveResponse struct {
	Status   int
	Complete int
	List     map[string]Item
	Since    int
}

// Item - Individual Pocket Item
type Item struct {
	ItemID        string `json:"item_id"`
	ResolvedID    string `json:"resolved_id"`
	GivenURL      string `json:"given_url"`
	GivenTitle    string `json:"given_title"`
	Favorite      int
	Status        int
	TimeAdded     string `json:"time_added"`
	TimeUpdated   string `json:"time_updated"`
	TimeRead      string `json:"time_read"`
	TimeFavorited string `json:"time_favorited"`
	SortID        string `json:"sort_id"`
	ResolvedTitle string `json:"resolved_title"`
	ResolvedURL   string `json:"resolved_url"`
	Excerpt       string
	IsArticle     int `json:"is_article"`
	IsIndex       int `json:"is_index"`
	HasVideo      int `json:"has_video"`
	HasImage      int `json:"has_image"`
	WordCount     int `json:"word_count"`
}
