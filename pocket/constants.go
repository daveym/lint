package pocket

// PocketHome - Pocket base url
const PocketHome string = "https://getpocket.com"

// CfgFile - Name of viper config file. Lint will store its values in the working directory.
const CfgFile string = "lint.yaml"

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
	// ModifyURL - API Address to modify Pocket Items
	ModifyURL = "https://getpocket.com/v3/send"
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
	ItemMediaAttachmentHasMedia ItemMediaAttachment = 1
	// ItemMediaAttachmentIsMedia - The Pocket Item is media only
	ItemMediaAttachmentIsMedia ItemMediaAttachment = 2
)

// ARTICLE TYPE
const (
	// ContentTypeArticle -  article item
	ContentTypeArticle ContentType = "article"
	// ContentTypeVideo - Video item
	ContentTypeVideo ContentType = "video"
	// ContentTypeImage - Image item
	ContentTypeImage ContentType = "image"
)

// DETAIL TYPE
const (
	// DetailTypeSimple - only return the titles and urls of each item
	DetailTypeSimple DetailType = "simple"
	// DetailTypeComplete - return all data about each item, including tags, images, authors, videos and more
	DetailTypeComplete DetailType = "complete"
)

// FAVOURITE FILTER
const (
	FavoriteFilterUnspecified FavoriteFilter = ""
	FavoriteFilterUnfavorited FavoriteFilter = "0"
	FavoriteFilterFavorited   FavoriteFilter = "1"
)

// SORT FILTER
const (
	SortNewest Sort = "newest"
	SortOldest Sort = "oldest"
	SortTitle  Sort = "title"
	SortSite   Sort = "site"
)
