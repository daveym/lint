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

// PocketHome - Pocket base url
const PocketHome string = "https://getpocket.com"

// URLS
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

// ITEM STATUS
const (
	// ItemStatusUnread - Pocket Item has not been read
	ItemStatusUnread ItemStatus = 0
	// ItemStatusArchived - Pocket Item has been archived
	ItemStatusArchived = 1
	// ItemStatusDeleted - Pocket Item has been deleted
	ItemStatusDeleted = 2
)

// ITEM STATE
const (
	// ItemStateUnread - only return unread items (default)
	ItemStateUnread string = "unread"
	// ItemStateArchive - only return archived items
	ItemStateArchive string = "archive"
	// ItemStateAll - return both unread and archived items
	ItemStateAll string = "all"
)

// MEDIA ATTACHMENTS
const (
	// ItemMediaAttachmentNoMedia - No media attached to the Pocket Item
	ItemMediaAttachmentNoMedia ItemMediaAttachment = 0
	// ItemMediaAttachmentHasMedia - Media is attached to the Pocket Item
	ItemMediaAttachmentHasMedia = 1
	// ItemMediaAttachmentIsMedia - The Pocket Item is media only
	ItemMediaAttachmentIsMedia = 2
)

// ARTICLE TYPE
const (
	// ContentTypeArticle -  article item
	ContentTypeArticle ContentType = "article"
	// ContentTypeVideo - Video item
	ContentTypeVideo = "video"
	// ContentTypeImage - Image item
	ContentTypeImage = "image"
)

// DETAIL TYPE
const (
	// DetailTypeSimple - only return the titles and urls of each item
	DetailTypeSimple DetailType = "simple"
	// DetailTypeComplete - return all data about each item, including tags, images, authors, videos and more
	DetailTypeComplete = "complete"
)

// FAVOURITE FILTER
const (
	FavoriteFilterUnspecified FavoriteFilter = ""
	FavoriteFilterUnfavorited                = "0"
	FavoriteFilterFavorited                  = "1"
)

// SORT FILTER
const (
	SortNewest Sort = "newest"
	SortOldest      = "oldest"
	SortTitle       = "title"
	SortSite        = "site"
)
