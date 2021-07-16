package player

import(
	"time"
)
type User struct {
	AvatarURL            string      `json:"avatar_url"`
	City                 interface{} `json:"city"`
	CommentsCount        int         `json:"comments_count"`
	CountryCode          interface{} `json:"country_code"`
	CreatedAt            time.Time   `json:"created_at"`
	CreatorSubscriptions []struct {
		Product struct {
			ID string `json:"id"`
		} `json:"product"`
	} `json:"creator_subscriptions"`
	CreatorSubscription struct {
		Product struct {
			ID string `json:"id"`
		} `json:"product"`
	} `json:"creator_subscription"`
	Description        interface{} `json:"description"`
	FollowersCount     int         `json:"followers_count"`
	FollowingsCount    int         `json:"followings_count"`
	FirstName          string      `json:"first_name"`
	FullName           string      `json:"full_name"`
	GroupsCount        int         `json:"groups_count"`
	ID                 int         `json:"id"`
	Kind               string      `json:"kind"`
	LastModified       time.Time   `json:"last_modified"`
	LastName           string      `json:"last_name"`
	LikesCount         int         `json:"likes_count"`
	PlaylistLikesCount int         `json:"playlist_likes_count"`
	Permalink          string      `json:"permalink"`
	PermalinkURL       string      `json:"permalink_url"`
	PlaylistCount      int         `json:"playlist_count"`
	RepostsCount       interface{} `json:"reposts_count"`
	TrackCount         int         `json:"track_count"`
	URI                string      `json:"uri"`
	Urn                string      `json:"urn"`
	Username           string      `json:"username"`
	Verified           bool        `json:"verified"`
	Visuals            interface{} `json:"visuals"`
	Badges             struct {
		Pro          bool `json:"pro"`
		ProUnlimited bool `json:"pro_unlimited"`
		Verified     bool `json:"verified"`
	} `json:"badges"`
	StationUrn       string `json:"station_urn"`
	StationPermalink string `json:"station_permalink"`
}

type PlayLists struct {
	Collection []struct {
		ArtworkURL     interface{} `json:"artwork_url"`
		CreatedAt      time.Time   `json:"created_at"`
		Description    string      `json:"description"`
		Duration       int         `json:"duration"`
		EmbeddableBy   string      `json:"embeddable_by"`
		Genre          string      `json:"genre"`
		ID             int         `json:"id"`
		Kind           string      `json:"kind"`
		LabelName      interface{} `json:"label_name"`
		LastModified   time.Time   `json:"last_modified"`
		License        string      `json:"license"`
		LikesCount     int         `json:"likes_count"`
		ManagedByFeeds bool        `json:"managed_by_feeds"`
		Permalink      string      `json:"permalink"`
		PermalinkURL   string      `json:"permalink_url"`
		Public         bool        `json:"public"`
		PurchaseTitle  interface{} `json:"purchase_title"`
		PurchaseURL    interface{} `json:"purchase_url"`
		ReleaseDate    time.Time   `json:"release_date"`
		RepostsCount   int         `json:"reposts_count"`
		SecretToken    interface{} `json:"secret_token"`
		Sharing        string      `json:"sharing"`
		TagList        string      `json:"tag_list"`
		Title          string      `json:"title"`
		URI            string      `json:"uri"`
		UserID         int         `json:"user_id"`
		SetType        string      `json:"set_type"`
		IsAlbum        bool        `json:"is_album"`
		PublishedAt    time.Time   `json:"published_at"`
		DisplayDate    time.Time   `json:"display_date"`
		User           struct {
			AvatarURL      string      `json:"avatar_url"`
			FirstName      string      `json:"first_name"`
			FollowersCount int         `json:"followers_count"`
			FullName       string      `json:"full_name"`
			ID             int         `json:"id"`
			Kind           string      `json:"kind"`
			LastModified   time.Time   `json:"last_modified"`
			LastName       string      `json:"last_name"`
			Permalink      string      `json:"permalink"`
			PermalinkURL   string      `json:"permalink_url"`
			URI            string      `json:"uri"`
			Urn            string      `json:"urn"`
			Username       string      `json:"username"`
			Verified       bool        `json:"verified"`
			City           interface{} `json:"city"`
			CountryCode    interface{} `json:"country_code"`
			Badges         struct {
				Pro          bool `json:"pro"`
				ProUnlimited bool `json:"pro_unlimited"`
				Verified     bool `json:"verified"`
			} `json:"badges"`
			StationUrn       string `json:"station_urn"`
			StationPermalink string `json:"station_permalink"`
		} `json:"user"`
		Tracks []struct {
			ArtworkURL        interface{} `json:"artwork_url,omitempty"`
			Caption           interface{} `json:"caption,omitempty"`
			Commentable       bool        `json:"commentable,omitempty"`
			CommentCount      int         `json:"comment_count,omitempty"`
			CreatedAt         time.Time   `json:"created_at,omitempty"`
			Description       string      `json:"description,omitempty"`
			Downloadable      bool        `json:"downloadable,omitempty"`
			DownloadCount     int         `json:"download_count,omitempty"`
			Duration          int         `json:"duration,omitempty"`
			FullDuration      int         `json:"full_duration,omitempty"`
			EmbeddableBy      string      `json:"embeddable_by,omitempty"`
			Genre             string      `json:"genre,omitempty"`
			HasDownloadsLeft  bool        `json:"has_downloads_left,omitempty"`
			ID                int         `json:"id"`
			Kind              string      `json:"kind"`
			LabelName         interface{} `json:"label_name,omitempty"`
			LastModified      time.Time   `json:"last_modified,omitempty"`
			License           string      `json:"license,omitempty"`
			LikesCount        int         `json:"likes_count,omitempty"`
			Permalink         string      `json:"permalink,omitempty"`
			PermalinkURL      string      `json:"permalink_url,omitempty"`
			PlaybackCount     int         `json:"playback_count,omitempty"`
			Public            bool        `json:"public,omitempty"`
			PublisherMetadata struct {
				ID            int    `json:"id"`
				Urn           string `json:"urn"`
				ContainsMusic bool   `json:"contains_music"`
			} `json:"publisher_metadata,omitempty"`
			PurchaseTitle interface{} `json:"purchase_title,omitempty"`
			PurchaseURL   interface{} `json:"purchase_url,omitempty"`
			ReleaseDate   time.Time   `json:"release_date,omitempty"`
			RepostsCount  int         `json:"reposts_count,omitempty"`
			SecretToken   interface{} `json:"secret_token,omitempty"`
			Sharing       string      `json:"sharing,omitempty"`
			State         string      `json:"state,omitempty"`
			Streamable    bool        `json:"streamable,omitempty"`
			TagList       string      `json:"tag_list,omitempty"`
			Title         string      `json:"title,omitempty"`
			TrackFormat   string      `json:"track_format,omitempty"`
			URI           string      `json:"uri,omitempty"`
			Urn           string      `json:"urn,omitempty"`
			UserID        int         `json:"user_id,omitempty"`
			Visuals       interface{} `json:"visuals,omitempty"`
			WaveformURL   string      `json:"waveform_url,omitempty"`
			DisplayDate   time.Time   `json:"display_date,omitempty"`
			Media         struct {
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
			} `json:"media,omitempty"`
			StationUrn         string `json:"station_urn,omitempty"`
			StationPermalink   string `json:"station_permalink,omitempty"`
			TrackAuthorization string `json:"track_authorization,omitempty"`
			MonetizationModel  string `json:"monetization_model"`
			Policy             string `json:"policy"`
			User               struct {
				AvatarURL      string      `json:"avatar_url"`
				FirstName      string      `json:"first_name"`
				FollowersCount int         `json:"followers_count"`
				FullName       string      `json:"full_name"`
				ID             int         `json:"id"`
				Kind           string      `json:"kind"`
				LastModified   time.Time   `json:"last_modified"`
				LastName       string      `json:"last_name"`
				Permalink      string      `json:"permalink"`
				PermalinkURL   string      `json:"permalink_url"`
				URI            string      `json:"uri"`
				Urn            string      `json:"urn"`
				Username       string      `json:"username"`
				Verified       bool        `json:"verified"`
				City           interface{} `json:"city"`
				CountryCode    interface{} `json:"country_code"`
				Badges         struct {
					Pro          bool `json:"pro"`
					ProUnlimited bool `json:"pro_unlimited"`
					Verified     bool `json:"verified"`
				} `json:"badges"`
				StationUrn       string `json:"station_urn"`
				StationPermalink string `json:"station_permalink"`
			} `json:"user,omitempty"`
		} `json:"tracks"`
		TrackCount int `json:"track_count"`
	} `json:"collection"`
	NextHref interface{} `json:"next_href"`
	QueryUrn interface{} `json:"query_urn"`
}

type PlayList struct {
	ArtworkURL     interface{} `json:"artwork_url"`
	CreatedAt      time.Time   `json:"created_at"`
	Description    string      `json:"description"`
	Duration       int         `json:"duration"`
	EmbeddableBy   string      `json:"embeddable_by"`
	Genre          string      `json:"genre"`
	ID             int         `json:"id"`
	Kind           string      `json:"kind"`
	LabelName      interface{} `json:"label_name"`
	LastModified   time.Time   `json:"last_modified"`
	License        string      `json:"license"`
	LikesCount     int         `json:"likes_count"`
	ManagedByFeeds bool        `json:"managed_by_feeds"`
	Permalink      string      `json:"permalink"`
	PermalinkURL   string      `json:"permalink_url"`
	Public         bool        `json:"public"`
	PurchaseTitle  interface{} `json:"purchase_title"`
	PurchaseURL    interface{} `json:"purchase_url"`
	ReleaseDate    time.Time   `json:"release_date"`
	RepostsCount   int         `json:"reposts_count"`
	SecretToken    interface{} `json:"secret_token"`
	Sharing        string      `json:"sharing"`
	TagList        string      `json:"tag_list"`
	Title          string      `json:"title"`
	URI            string      `json:"uri"`
	UserID         int         `json:"user_id"`
	SetType        string      `json:"set_type"`
	IsAlbum        bool        `json:"is_album"`
	PublishedAt    time.Time   `json:"published_at"`
	DisplayDate    time.Time   `json:"display_date"`
	User           struct {
		AvatarURL      string      `json:"avatar_url"`
		FirstName      string      `json:"first_name"`
		FollowersCount int         `json:"followers_count"`
		FullName       string      `json:"full_name"`
		ID             int         `json:"id"`
		Kind           string      `json:"kind"`
		LastModified   time.Time   `json:"last_modified"`
		LastName       string      `json:"last_name"`
		Permalink      string      `json:"permalink"`
		PermalinkURL   string      `json:"permalink_url"`
		URI            string      `json:"uri"`
		Urn            string      `json:"urn"`
		Username       string      `json:"username"`
		Verified       bool        `json:"verified"`
		City           interface{} `json:"city"`
		CountryCode    interface{} `json:"country_code"`
		Badges         struct {
			Pro          bool `json:"pro"`
			ProUnlimited bool `json:"pro_unlimited"`
			Verified     bool `json:"verified"`
		} `json:"badges"`
		StationUrn       string `json:"station_urn"`
		StationPermalink string `json:"station_permalink"`
	} `json:"user"`
	Tracks []struct {
		ArtworkURL        interface{} `json:"artwork_url,omitempty"`
		Caption           interface{} `json:"caption,omitempty"`
		Commentable       bool        `json:"commentable,omitempty"`
		CommentCount      int         `json:"comment_count,omitempty"`
		CreatedAt         time.Time   `json:"created_at,omitempty"`
		Description       string      `json:"description,omitempty"`
		Downloadable      bool        `json:"downloadable,omitempty"`
		DownloadCount     int         `json:"download_count,omitempty"`
		Duration          int         `json:"duration,omitempty"`
		FullDuration      int         `json:"full_duration,omitempty"`
		EmbeddableBy      string      `json:"embeddable_by,omitempty"`
		Genre             string      `json:"genre,omitempty"`
		HasDownloadsLeft  bool        `json:"has_downloads_left,omitempty"`
		ID                int         `json:"id"`
		Kind              string      `json:"kind"`
		LabelName         interface{} `json:"label_name,omitempty"`
		LastModified      time.Time   `json:"last_modified,omitempty"`
		License           string      `json:"license,omitempty"`
		LikesCount        int         `json:"likes_count,omitempty"`
		Permalink         string      `json:"permalink,omitempty"`
		PermalinkURL      string      `json:"permalink_url,omitempty"`
		PlaybackCount     int         `json:"playback_count,omitempty"`
		Public            bool        `json:"public,omitempty"`
		PublisherMetadata struct {
			ID            int    `json:"id"`
			Urn           string `json:"urn"`
			ContainsMusic bool   `json:"contains_music"`
		} `json:"publisher_metadata,omitempty"`
		PurchaseTitle interface{} `json:"purchase_title,omitempty"`
		PurchaseURL   interface{} `json:"purchase_url,omitempty"`
		ReleaseDate   time.Time   `json:"release_date,omitempty"`
		RepostsCount  int         `json:"reposts_count,omitempty"`
		SecretToken   interface{} `json:"secret_token,omitempty"`
		Sharing       string      `json:"sharing,omitempty"`
		State         string      `json:"state,omitempty"`
		Streamable    bool        `json:"streamable,omitempty"`
		TagList       string      `json:"tag_list,omitempty"`
		Title         string      `json:"title,omitempty"`
		TrackFormat   string      `json:"track_format,omitempty"`
		URI           string      `json:"uri,omitempty"`
		Urn           string      `json:"urn,omitempty"`
		UserID        int         `json:"user_id,omitempty"`
		Visuals       interface{} `json:"visuals,omitempty"`
		WaveformURL   string      `json:"waveform_url,omitempty"`
		DisplayDate   time.Time   `json:"display_date,omitempty"`
		Media         struct {
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
		} `json:"media,omitempty"`
		StationUrn         string `json:"station_urn,omitempty"`
		StationPermalink   string `json:"station_permalink,omitempty"`
		TrackAuthorization string `json:"track_authorization,omitempty"`
		MonetizationModel  string `json:"monetization_model"`
		Policy             string `json:"policy"`
		User               struct {
			AvatarURL      string      `json:"avatar_url"`
			FirstName      string      `json:"first_name"`
			FollowersCount int         `json:"followers_count"`
			FullName       string      `json:"full_name"`
			ID             int         `json:"id"`
			Kind           string      `json:"kind"`
			LastModified   time.Time   `json:"last_modified"`
			LastName       string      `json:"last_name"`
			Permalink      string      `json:"permalink"`
			PermalinkURL   string      `json:"permalink_url"`
			URI            string      `json:"uri"`
			Urn            string      `json:"urn"`
			Username       string      `json:"username"`
			Verified       bool        `json:"verified"`
			City           interface{} `json:"city"`
			CountryCode    interface{} `json:"country_code"`
			Badges         struct {
				Pro          bool `json:"pro"`
				ProUnlimited bool `json:"pro_unlimited"`
				Verified     bool `json:"verified"`
			} `json:"badges"`
			StationUrn       string `json:"station_urn"`
			StationPermalink string `json:"station_permalink"`
		} `json:"user,omitempty"`
	} `json:"tracks"`
	TrackCount int `json:"track_count"`
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
