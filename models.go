package subsonic

import "time"

// Type subsonicResponse is the main target for unmarshalling JSON data from the API - everything within the "subsonic-response" key
type subsonicResponse struct {
	Status        string `json:"status"`        // standard
	Version       string `json:"version"`       // standard
	Type          string `json:"type"`          // navidrome
	ServerVersion string `json:"serverVersion"` // navidrome
	Error         *errorResponse
	License       *License              // getLicense
	MusicFolders  *musicFolderContainer // getMusicFolders
	Indexes       *Index                // getIndexes
	Directory     *Directory            // getMusicDirectory
	Genres        *genreContainer       // getGenres
	Artists       *ArtistsContainer     // getArtists
	Artist        *Artist               // getArtist
	Album         *Album                // getAlbum
	Song          *Song                 // getSong
	ArtistInfo    *ArtistInfo           // getArtistInfo
	ArtistInfo2   *ArtistInfo           // getArtistInfo2
	AlbumInfo     *AlbumInfo            // getAlbumInfo
	SimilarSongs  *songList             // getSimilarSongs
	SimilarSongs2 *songList             // getSimilarSongs2
	TopSongs      *songList             // getTopSongs
	AlbumList     *albumList            // getAlbumList
	AlbumList2    *albumList            // getAlbumList2
	RandomSongs   *songList             // getRandomSongs
	SongsByGenre  *songList             // getSongsByGenre
	NowPlaying    *nowPlayingList       // getNowPlaying
	Starred       *Starred              // getStarred
	Starred2      *Starred              // getStarred2
	SearchResult2 *SearchResult         // search2
	SearchResult3 *SearchResult         // search3
	Playlists     *playlistList         // getPlaylists
	Playlist      *Playlist             // getPlaylist
}

type apiResponse struct {
	Response *subsonicResponse `json:"subsonic-response"`
}

type errorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// License contains information about the Subsonic server's license validity and contact information in the case of a trial subscription.
type License struct {
	Valid        bool   `json:"valid"`        // standard
	Email        string `json:"email"`        // subsonic
	TrialExpires string `json:"trialExpires"` // subsonic
}

// MusicFolder is a representation of a source of music files added to the server. It is identified primarily by the numeric ID.
type MusicFolder struct {
	Id   int    `json:"id"` // subsonic returns an int, navidrome a string
	Name string `json:"name"`
}

type musicFolderContainer struct {
	Folders []*MusicFolder `json:"musicFolder"`
}

// Song is all metadata about a single song from the server.
type Song struct {
	ID            string    `json:"id"`
	AlbumID       string    `json:"albumId"`
	Album         string    `json:"album"`
	ArtistID      string    `json:"artistId"`
	Artist        string    `json:"artist"`
	BitRate       int       `json:"bitRate"`
	ContentType   string    `json:"contentType"`
	Created       time.Time `json:"created"`
	Starred       time.Time `json:"starred,omitempty"` // getStarred only
	Duration      int       `json:"duration"`
	Genre         string    `json:"genre"`
	IsDir         bool      `json:"isDir"`
	Parent        string    `json:"parent"`
	Path          string    `json:"path"`
	PlayCount     int       `json:"playCount"`
	Size          int       `json:"size"`
	Suffix        string    `json:"suffix"`
	Title         string    `json:"title"`
	Track         int       `json:"track"`
	Type          string    `json:"type"`
	Year          int       `json:"year"`
	AverageRating float32   `json:"averageRating,omitempty"` // subsonic only
	CoverArt      string    `json:"coverArt"`                // subsonic only
}

// Album is all metadata about an album from the server, including songs if fetched from getAlbum.
type Album struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"` // getAlbumList & getStarred returns each album with a "title" key rather than a "name" key
	Artist    string    `json:"artist"`
	ArtistID  string    `json:"artistId,omitempty"`
	SongCount int       `json:"songCount"`
	Duration  int       `json:"duration"`
	Created   time.Time `json:"created"`
	Starred   time.Time `json:"starred,omitempty"` // getStarred only
	Year      int       `json:"year"`
	Genre     string    `json:"genre"`
	PlayCount int       `json:"playCount"`
	CoverArt  string    `json:"coverArt"`
	IsDir     bool      `json:"isDir"`
	Songs     []*Song   `json:"song,omitempty"`
	IsVideo   bool      `json:"isVideo,omitempty"` // navidrome only
	Size      string    `json:"size,omitempty"`    // navidrome only
}

// Artist is a representation of one artist from the server.
// Many calls return Artists with few fields, but getArtist will give more data.
type Artist struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	AlbumCount     int      `json:"albumCount"`
	ArtistImageURL string   `json:"artistImageUrl"` // subsonic only
	CoverArt       string   `json:"coverArt"`       // subsonic only
	Albums         []*Album `json:"album"`          // only filled by getArtist
}

// ArtistIndex is a by-letter representation of every artist on the server.
type ArtistIndex struct {
	Name    string    `json:"name"`
	Artists []*Artist `json:"artist"`
}

// Index holds every artist or single track in the database alphabetically sorted.
type Index struct {
	LastModified    int64          `json:"lastModified"` // subsonic returns an int64, navidrome a string
	IgnoredArticles string         `json:"ignoredArticles"`
	Indexes         []*ArtistIndex `json:"index"`
	Children        []*Child       `json:"child"`
}

type Child struct {
	ID            string    `json:"id"`
	Album         string    `json:"album"`
	BitRate       int       `json:"bitRate"`
	ContentType   string    `json:"contentType"`
	CoverArt      string    `json:"coverArt"`
	Created       time.Time `json:"created"`
	Duration      int       `json:"duration"`
	IsDir         bool      `json:"isDir"`
	IsVideo       bool      `json:"isVideo"`
	Parent        string    `json:"parent"`
	Path          string    `json:"path"`
	Size          int64     `json:"size"` // string in navidrome
	Suffix        string    `json:"suffix"`
	Title         string    `json:"title"`
	Type          string    `json:"type"`
	PlayCount     int       `json:"playCount"`               // subsonic / airsonic
	UserRating    int       `json:"userRating"`              // subsonic only
	AverageRating float32   `json:"averageRating,omitempty"` // subsonic only
	Artist        string    `json:"artist"`                  // this and all following fields are navidrome only
	Track         int       `json:"track"`
	Year          int       `json:"year"`
	Genre         string    `json:"genre"`
	DiscNumber    int       `json:"discNumber"`
	AlbumID       string    `json:"albumId"`
	ArtistID      string    `json:"artistId"`
}

type Directory struct {
	Children   []*Child `json:"child"`
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	PlayCount  int      `json:"playCount"`  // airsonic only
	AlbumCount int      `json:"albumCount"` // navidrome only
	Parent     string   `json:"parent"`     // navidrome only
}

// Genre is a generic tag describing the style of a song or album. The Value property of this struct is the name of the genre.
type Genre struct {
	SongCount  int    `json:"songCount"`
	AlbumCount int    `json:"albumCount"`
	Value      string `json:"value"`
}

type genreContainer struct {
	Genre []*Genre `json:"genre"`
}

type ArtistsContainer struct {
	IgnoredArticles string         `json:"ignoredArticles"`
	Indexes         []*ArtistIndex `json:"index"`
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo/GetArtistInfo2
type ArtistInfo struct {
	Biography      string    `json:"biography"`
	MusicBrainzID  string    `json:"musicBrainzId"`
	LastFmURL      string    `json:"lastFmUrl"`
	SmallImageURL  string    `json:"smallImageUrl"`
	MediumImageURL string    `json:"mediumImageUrl"`
	LargeImageURL  string    `json:"largeImageUrl"`
	SimilarArtist  []*Artist `json:"similarArtist"`
}

// AlbumInfo is a collection of notes and links describing an album.
// Fetch one by ID with GetAlbumInfo/GetAlbumInfo2.
type AlbumInfo struct {
	Notes          string `json:"notes"`
	MusicBrainzID  string `json:"musicBrainzId"`
	LastFmURL      string `json:"lastFmUrl"`
	SmallImageURL  string `json:"smallImageUrl"`
	MediumImageURL string `json:"mediumImageUrl"`
	LargeImageURL  string `json:"largeImageUrl"`
}

type songList struct {
	Songs []*Song `json:"song"`
}

type albumList struct {
	Albums []*Album `json:"album"`
}

// NowPlaying is a data about a recently played song from a user, including recent plays (MinutesAgo).
type NowPlaying struct {
	ID          string `json:"id"`
	Album       string `json:"album"`
	BitRate     int    `json:"bitRate"`
	ContentType string `json:"contentType"`
	CoverArt    string `json:"coverArt"`
	Created     string `json:"created"`
	Duration    int64  `json:"duration"`
	IsDir       bool   `json:"isDir"`
	IsVideo     bool   `json:"isVideo"`
	MinutesAgo  int    `json:"minutesAgo"`
	Parent      string `json:"parent"`
	Path        string `json:"path"`
	PlayCount   int    `json:"playCount"`
	PlayerID    int64  `json:"playerId"`
	Size        int64  `json:"size"`
	Suffix      string `json:"suffix"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	Username    string `json:"username"`
}

type nowPlayingList struct {
	Entries []*NowPlaying `json:"entry"`
}

// Starred is a collection of songs, albums, and artists flagged by a user as starred.
type Starred struct {
	Songs   []*Song   `json:"song"`
	Albums  []*Album  `json:"album"`
	Artists []*Artist `json:"artist"`
}

// SearchResult is a collection of songs, albums, and artists returned by a call to Search2 or Search3.
type SearchResult struct {
	Songs   []*Song   `json:"song"`
	Albums  []*Album  `json:"album"`
	Artists []*Artist `json:"artist"`
}

type Playlist struct {
	Changed   string  `json:"changed"`
	Comment   string  `json:"comment"`
	CoverArt  string  `json:"coverArt"`
	Created   string  `json:"created"`
	Duration  int64   `json:"duration"`
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Owner     string  `json:"owner"`
	Public    bool    `json:"public"`
	SongCount int     `json:"songCount"`
	Entries   []*Song `json:"entry"` // getPlaylist only
}

type playlistList struct {
	Entries []*Playlist `json:"playlist"`
}
