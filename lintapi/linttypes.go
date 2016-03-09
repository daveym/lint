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
	// RetrieveURL - API Address to query Pocket Items
	RetrieveURL = "https://getpocket.com/v3/get"
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
	// ItemStateUnread - only return unread items (default)
	ItemStateUnread string = "unread"
	// ItemStateArchive - only return archived items
	ItemStateArchive string = "archive"
	// ItemStateAll - return both unread and archived items
	ItemStateAll string = "all"
)

const (
	// ItemMediaAttachmentNoMedia - No media attached to the Pocket Item
	ItemMediaAttachmentNoMedia ItemMediaAttachment = 0
	// ItemMediaAttachmentHasMedia - Media is attached to the Pocket Item
	ItemMediaAttachmentHasMedia = 1
	// ItemMediaAttachmentIsMedia - The Pocket Item is media only
	ItemMediaAttachmentIsMedia = 2
)

const (
	// ContentTypeArticle -  article item
	ContentTypeArticle ContentType = "article"
	// ContentTypeVideo - Video item
	ContentTypeVideo = "video"
	// ContentTypeImage - Image item
	ContentTypeImage = "image"
)

const (
	// DetailTypeSimple - only return the titles and urls of each item
	DetailTypeSimple DetailType = "simple"
	// DetailTypeComplete - return all data about each item, including tags, images, authors, videos and more
	DetailTypeComplete = "complete"
)

// TYPES

// Time - alias for time.Time
type Time time.Time

// ItemStatus - Status of retrieved item
type ItemStatus int

// ContentType - Article video or image
type ContentType string

// DetailType - Simple or complete detail
type DetailType string

// ItemMediaAttachment - Is Media attached to the Pocket Item. Defined by ItemMedia constants
type ItemMediaAttachment int

// Configuration - holds Lint configuration in JSON format
type Configuration struct {
	// ConsumerKey - Private Key to authenticate against Pocket API
	ConsumerKey string
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
	Since       Time
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