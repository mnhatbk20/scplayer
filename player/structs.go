package player

import(
	"time"
)

type PlayList struct {
	Duration     int         `json:"duration"`
	PermalinkURL string      `json:"permalink_url"`
	RepostsCount int         `json:"reposts_count"`
	Genre        string      `json:"genre"`
	Permalink    string      `json:"permalink"`
	PurchaseURL  interface{} `json:"purchase_url"`
	Description  interface{} `json:"description"`
	URI          string      `json:"uri"`
	LabelName    interface{} `json:"label_name"`
	TagList      string      `json:"tag_list"`
	SetType      string      `json:"set_type"`
	Public       bool        `json:"public"`
	TrackCount   int         `json:"track_count"`
	UserID       int         `json:"user_id"`
	LastModified time.Time   `json:"last_modified"`
	License      string      `json:"license"`
	Tracks       []struct {
		CommentCount int       `json:"comment_count,omitempty"`
		FullDuration int       `json:"full_duration,omitempty"`
		Downloadable bool      `json:"downloadable,omitempty"`
		CreatedAt    time.Time `json:"created_at,omitempty"`
		Description  string    `json:"description,omitempty"`
		Media        struct {
			Transcodings []interface{} `json:"transcodings"`
		} `json:"media,omitempty"`
		Title             string `json:"title,omitempty"`
		PublisherMetadata struct {
		} `json:"publisher_metadata,omitempty"`
		Duration          int         `json:"duration,omitempty"`
		HasDownloadsLeft  bool        `json:"has_downloads_left,omitempty"`
		ArtworkURL        interface{} `json:"artwork_url,omitempty"`
		Public            bool        `json:"public,omitempty"`
		Streamable        bool        `json:"streamable,omitempty"`
		TagList           string      `json:"tag_list,omitempty"`
		Genre             string      `json:"genre,omitempty"`
		ID                int         `json:"id,omitempty"`
		RepostsCount      int         `json:"reposts_count,omitempty"`
		State             string      `json:"state,omitempty"`
		LabelName         interface{} `json:"label_name,omitempty"`
		LastModified      time.Time   `json:"last_modified,omitempty"`
		Commentable       bool        `json:"commentable,omitempty"`
		Policy            string      `json:"policy,omitempty"`
		Visuals           interface{} `json:"visuals,omitempty"`
		Kind              string      `json:"kind,omitempty"`
		PurchaseURL       interface{} `json:"purchase_url,omitempty"`
		Sharing           string      `json:"sharing,omitempty"`
		URI               string      `json:"uri,omitempty"`
		SecretToken       interface{} `json:"secret_token,omitempty"`
		DownloadCount     int         `json:"download_count,omitempty"`
		LikesCount        int         `json:"likes_count,omitempty"`
		Urn               string      `json:"urn,omitempty"`
		License           string      `json:"license,omitempty"`
		PurchaseTitle     interface{} `json:"purchase_title,omitempty"`
		DisplayDate       time.Time   `json:"display_date,omitempty"`
		EmbeddableBy      string      `json:"embeddable_by,omitempty"`
		ReleaseDate       interface{} `json:"release_date,omitempty"`
		UserID            int         `json:"user_id,omitempty"`
		MonetizationModel string      `json:"monetization_model,omitempty"`
		WaveformURL       string      `json:"waveform_url,omitempty"`
		Permalink         string      `json:"permalink,omitempty"`
		PermalinkURL      string      `json:"permalink_url,omitempty"`
		User              struct {
			AvatarURL    string    `json:"avatar_url"`
			FirstName    string    `json:"first_name"`
			FullName     string    `json:"full_name"`
			ID           int       `json:"id"`
			Kind         string    `json:"kind"`
			LastModified time.Time `json:"last_modified"`
			LastName     string    `json:"last_name"`
			Permalink    string    `json:"permalink"`
			PermalinkURL string    `json:"permalink_url"`
			URI          string    `json:"uri"`
			Urn          string    `json:"urn"`
			Username     string    `json:"username"`
			Verified     bool      `json:"verified"`
			City         string    `json:"city"`
			CountryCode  string    `json:"country_code"`
		} `json:"user,omitempty"`
		PlaybackCount int `json:"playback_count,omitempty"`
	} `json:"tracks"`
	ID             int         `json:"id"`
	ReleaseDate    interface{} `json:"release_date"`
	DisplayDate    time.Time   `json:"display_date"`
	Sharing        string      `json:"sharing"`
	SecretToken    interface{} `json:"secret_token"`
	CreatedAt      time.Time   `json:"created_at"`
	LikesCount     int         `json:"likes_count"`
	Kind           string      `json:"kind"`
	Title          string      `json:"title"`
	PurchaseTitle  interface{} `json:"purchase_title"`
	ManagedByFeeds bool        `json:"managed_by_feeds"`
	ArtworkURL     string      `json:"artwork_url"`
	IsAlbum        bool        `json:"is_album"`
	User           struct {
		AvatarURL    string      `json:"avatar_url"`
		FirstName    string      `json:"first_name"`
		FullName     string      `json:"full_name"`
		ID           int         `json:"id"`
		Kind         string      `json:"kind"`
		LastModified time.Time   `json:"last_modified"`
		LastName     string      `json:"last_name"`
		Permalink    string      `json:"permalink"`
		PermalinkURL string      `json:"permalink_url"`
		URI          string      `json:"uri"`
		Urn          string      `json:"urn"`
		Username     string      `json:"username"`
		Verified     bool        `json:"verified"`
		City         interface{} `json:"city"`
		CountryCode  interface{} `json:"country_code"`
	} `json:"user"`
	PublishedAt  time.Time `json:"published_at"`
	EmbeddableBy string    `json:"embeddable_by"`
}
type Track struct {
	CommentCount int       `json:"comment_count"`
	FullDuration int       `json:"full_duration"`
	Downloadable bool      `json:"downloadable"`
	CreatedAt    time.Time `json:"created_at"`
	Description  string    `json:"description"`
	Media        struct {
		Transcodings []struct {
			URL      string `json:"url"`
			Preset   string `json:"preset"`
			Duration int    `json:"duration"`
			Snipped  bool   `json:"snipped"`
			Format   struct {
				Protocol string `json:"protocol"`
				MimeType string `json:"mime_type"`
			} `json:"format"`
			Quality string `json:"quality"`
		} `json:"transcodings"`
	} `json:"media"`
	Title             string `json:"title"`
	PublisherMetadata struct {
		Urn            string `json:"urn"`
		ContainsMusic  bool   `json:"contains_music"`
		Artist         string `json:"artist"`
		WriterComposer string `json:"writer_composer"`
		Publisher      string `json:"publisher"`
		Isrc           string `json:"isrc"`
		ID             int    `json:"id"`
	} `json:"publisher_metadata"`
	Duration          int         `json:"duration"`
	HasDownloadsLeft  bool        `json:"has_downloads_left"`
	ArtworkURL        interface{} `json:"artwork_url"`
	Public            bool        `json:"public"`
	Streamable        bool        `json:"streamable"`
	TagList           string      `json:"tag_list"`
	Genre             string      `json:"genre"`
	ID                int         `json:"id"`
	RepostsCount      int         `json:"reposts_count"`
	State             string      `json:"state"`
	LabelName         interface{} `json:"label_name"`
	LastModified      time.Time   `json:"last_modified"`
	Commentable       bool        `json:"commentable"`
	Policy            string      `json:"policy"`
	Visuals           interface{} `json:"visuals"`
	Kind              string      `json:"kind"`
	PurchaseURL       interface{} `json:"purchase_url"`
	Sharing           string      `json:"sharing"`
	URI               string      `json:"uri"`
	SecretToken       interface{} `json:"secret_token"`
	DownloadCount     int         `json:"download_count"`
	LikesCount        int         `json:"likes_count"`
	Urn               string      `json:"urn"`
	License           string      `json:"license"`
	PurchaseTitle     interface{} `json:"purchase_title"`
	DisplayDate       time.Time   `json:"display_date"`
	EmbeddableBy      string      `json:"embeddable_by"`
	ReleaseDate       interface{} `json:"release_date"`
	UserID            int         `json:"user_id"`
	MonetizationModel string      `json:"monetization_model"`
	WaveformURL       string      `json:"waveform_url"`
	Permalink         string      `json:"permalink"`
	PermalinkURL      string      `json:"permalink_url"`
	User              struct {
		AvatarURL    string    `json:"avatar_url"`
		FirstName    string    `json:"first_name"`
		FullName     string    `json:"full_name"`
		ID           int       `json:"id"`
		Kind         string    `json:"kind"`
		LastModified time.Time `json:"last_modified"`
		LastName     string    `json:"last_name"`
		Permalink    string    `json:"permalink"`
		PermalinkURL string    `json:"permalink_url"`
		URI          string    `json:"uri"`
		Urn          string    `json:"urn"`
		Username     string    `json:"username"`
		Verified     bool      `json:"verified"`
		City         string    `json:"city"`
		CountryCode  string    `json:"country_code"`
	} `json:"user"`
	PlaybackCount int `json:"playback_count"`
}
type Stream struct {
	URL string `json:"url"`
}
