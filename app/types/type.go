package types

type Header struct {
	StatusCode  int64   `json:"status_code"`
	ExecuteTime float64 `json:"execute_time"`
}

type LyricsResponse struct {
	Message LyricsMessage `json:"message"`
}

type LyricsMessage struct {
	Header Header     `json:"header"`
	Body   LyricsBody `json:"body"`
}

type LyricsBody struct {
	Lyrics Lyrics `json:"lyrics"`
}

type Lyrics struct {
	LyricsId          int64  `json:"lyrics_id"`
	Restricted        int64  `json:"restricted"`
	Instrumental      int64  `json:"instrumental"`
	LyricsBody        string `json:"lyrics_body"`
	LyricsLanguage    string `json:"lyrics_language"`
	ScriptTrackingUrl string `json:"script_tracking_url"`
	PixelTrackingUrl  string `json:"pixel_tracking_url"`
	LyricsCopyright   string `json:"lyrics_copyright"`
	BacklinkUrl       string `json:"backlink_url"`
	UpdatedTime       string `json:"updated_time"`
}

type ArtistResponse struct {
	Message ArtistMessage `json:"message"`
}

type ArtistMessage struct {
	Header Header     `json:"header"`
	Body   ArtistBody `json:"body"`
}

type ArtistBody struct {
	Artist Artist `json:"artist"`
}

type Artist struct {
	ArtistId         int64         `json:"artist_id"`
	ArtistMbid       string        `json:"artist_mbid"`
	ArtistName       string        `json:"artist_name"`
	ArtistCountry    string        `json:"artist_country"`
	ArtistRating     int64         `json:"artist_rating"`
	ArtistTwitterUrl string        `json:"artist_twitter_url"`
	ArtistAliasList  []ArtistAlias `json:"artist_alias_list"`
	UpdatedTime      string        `json:"updated_time"`
}

type ArtistAlias struct {
	ArtistAlias string `json:"artist_alias"`
}

type TrackRegionResponse struct {
	Message TrackRegionMessage `json:"message"`
}

type TrackRegionMessage struct {
	Header Header          `json:"header"`
	Body   TrackRegionBody `json:"body"`
}

type TrackRegionBody struct {
	TrackList []TrackRegionObj `json:"track_list"`
}

type TrackRegionObj struct {
	Track Track `json:"track"`
}

type Track struct {
	AlbumId       int64         `json:"album_id"`
	AlbumName     string        `json:"album_name"`
	ArtistId      int64         `json:"artist_id"`
	ArtistName    string        `json:"artist_name"`
	CommonTrackId int64         `json:"commontrack_id"`
	Explicit      int64         `json:"explicit"`
	HasLyrics     int64         `json:"has_lyrics"`
	HasRichSync   int64         `json:"has_richsync"`
	HasSubtitles  int64         `json:"has_subtitles"`
	Instrumental  int64         `json:"instrumental"`
	NumFavourite  int64         `json:"num_favourite"`
	Restricted    int64         `json:"restricted"`
	TrackEditUrl  string        `json:"track_edit_url"`
	TrackId       int64         `json:"track_id"`
	TrackName     string        `json:"track_name"`
	TrackRating   int64         `json:"track_rating"`
	TrackShareUrl string        `json:"track_share_url"`
	UpdatedTime   string        `json:"updated_time"`
	PrimaryGenres PrimaryGenres `json:"primary_genres"`
	Artist        Artist        `json:"artist"`
	Lyrics        Lyrics        `json:"lyrics"`
}

type PrimaryGenres struct {
	MusicGenreList []MusicGenresList `json:"music_genre_list"`
}

type MusicGenresList struct {
	MusicGenre MusicGenre `json:"music_genre"`
}

type MusicGenre struct {
	MusicGenreId       int64  `json:"music_genre_id"`
	MusicGenreName     string `json:"music_genre_name"`
	MusicGenreExtended string `json:"music_genre_name_extended"`
	MusicGenreParentId int64  `json:"music_genre_parent_id"`
	MusicGenreVanity   string `json:"music_genre_vanity"`
}
