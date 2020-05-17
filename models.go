package subsonic

import "time"

// Type SubsonicResponse is the main target for unmarshalling JSON data from the API - everything within the "subsonic-response" key
type SubsonicResponse struct {
	Status        string `json:"status"`        // standard
	Version       string `json:"version"`       // standard
	Type          string `json:"type"`          // navidrome
	ServerVersion string `json:"serverVersion"` // navidrome
	Error         *ErrorResponse
	License       *LicenseValidity      // getLicense
	MusicFolders  *MusicFolderContainer // getMusicFolders
	Indexes       *IndexContainer       // getIndexes
	Directory     *Directory            // getMusicDirectory
	Genres        *GenreContainer       // getGenres
	Artists       *ArtistsContainer     // getArtists
	Artist        *Artist               // getArtist
	Album         *Album                // getAlbum
	Song          *Song                 // getSong
	ArtistInfo    *ArtistInfo           // getArtistInfo
	ArtistInfo2   *ArtistInfo           // getArtistInfo2
	AlbumInfo     *AlbumInfo            // getAlbumInfo
	SimilarSongs  *SongList             // getSimilarSongs
	SimilarSongs2 *SongList             // getSimilarSongs2
	TopSongs      *SongList             // getTopSongs
	AlbumList     *AlbumList            // getAlbumList
	AlbumList2    *AlbumList            // getAlbumList2
	RandomSongs   *SongList             // getRandomSongs
	SongsByGenre  *SongList             // getSongsByGenre
}

type APIResponse struct {
	Response *SubsonicResponse `json:"subsonic-response"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type LicenseValidity struct {
	Valid        bool   `json:"valid"`        // standard
	Email        string `json:"email"`        // subsonic
	TrialExpires string `json:"trialExpires"` // subsonic
}

type MusicFolder struct {
	Id   int    `json:"id"` // subsonic returns an int, navidrome a string
	Name string `json:"name"`
}

type MusicFolderContainer struct {
	Folders []*MusicFolder `json:"musicFolder"`
}

type Song struct {
	ID            string    `json:"id"`
	AlbumID       string    `json:"albumId"`
	Album         string    `json:"album"`
	ArtistID      string    `json:"artistId"`
	Artist        string    `json:"artist"`
	BitRate       int       `json:"bitRate"`
	ContentType   string    `json:"contentType"`
	Created       time.Time `json:"created"`
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
	AverageRating float32   `json:"averageRating,omitempty"` // subsonic only
	CoverArt      string    `json:"coverArt"`                // subsonic only
}

type Album struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Title     string    `json:"title"` // getAlbumList returns each album with a "title" key rather than a "name" key
	Artist    string    `json:"artist"`
	ArtistID  string    `json:"artistId"`
	SongCount int       `json:"songCount"`
	Duration  int       `json:"duration"`
	Created   time.Time `json:"created"`
	Year      int       `json:"year"`
	Genre     string    `json:"genre"`
	PlayCount int       `json:"playCount"`
	CoverArt  string    `json:"coverArt"`
	IsDir     bool      `json:"isDir"`
	Songs     []*Song   `json:"song"`    // populated by getAlbum
	IsVideo   bool      `json:"isVideo"` // navidrome only
	Size      string    `json:"size"`    // navidrome only
}

// Artists are obtained by calls to GetIndex (with few fields), and GetArtists/GetArtist with more fields.
type Artist struct {
	ID             string   `json:"id"`
	Name           string   `json:"name"`
	AlbumCount     int      `json:"albumCount"`
	ArtistImageURL string   `json:"artistImageUrl"` // subsonic only
	CoverArt       string   `json:"coverArt"`       // subsonic only
	Albums         []*Album `json:"album"`          // only filled by getArtist
}

// Type Index contains a by-letter representation of every item in the database.
type Index struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artist"`
}

type IndexContainer struct {
	LastModified    int64   `json:"lastModified"` // subsonic returns an int64, navidrome a string
	IgnoredArticles string  `json:"ignoredArticles"`
	Indexes         []Index `json:"index"`
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

type Genre struct {
	SongCount  int    `json:"songCount"`
	AlbumCount int    `json:"albumCount"`
	Value      string `json:"value"`
}

type GenreContainer struct {
	Genre []*Genre `json:"genre"`
}

type ArtistsContainer struct {
	IgnoredArticles string  `json:"ignoredArticles"`
	Indexes         []Index `json:"index"`
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo/GetArtistInfo2
type ArtistInfo struct {
	Biography      string          `json:"biography"`
	MusicBrainzID  string          `json:"musicBrainzId"`
	LastFmURL      string          `json:"lastFmUrl"`
	SmallImageURL  string          `json:"smallImageUrl"`
	MediumImageURL string          `json:"mediumImageUrl"`
	LargeImageURL  string          `json:"largeImageUrl"`
	SimilarArtist  []SimilarArtist `json:"similarArtist"`
}

type SimilarArtist struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	AlbumCount int    `json:"albumCount"`
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

type SongList struct {
	Songs []*Song `json:"song"`
}

type AlbumList struct {
	Albums []*Album `json:"album"`
}
