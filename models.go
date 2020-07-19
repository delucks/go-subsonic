package subsonic

/* This file was automatically generated from the xsd schema provided by Subsonic, then manually modified.
 *   http://www.subsonic.org/pages/inc/api/schema/subsonic-rest-api-1.16.1.xsd
 *   xsdgen -o xml.go -pkg subsonic -ns "http://subsonic.org/restapi" subsonic-rest-api-1.16.1.xsd
 * Changes from the original include:
 * - Adding missing name (value of xml element) for each genre
 * - Capitalize "ID" in struct names and add missing ID fields.
 * - Merge *With* variants of structs.
 */

import (
	"bytes"
	"encoding/xml"
	"time"
)

// AlbumID3 is an album that's organized by music file tags.
type AlbumID3 struct {
	ID        string    `xml:"id,attr"`                                    // Manually added
	Song      []*Child  `xml:"http://subsonic.org/restapi song,omitempty"` // Merged from AlbumWithSongsID3
	Name      string    `xml:"name,attr"`
	Artist    string    `xml:"artist,attr,omitempty"`
	ArtistID  string    `xml:"artistId,attr,omitempty"`
	CoverArt  string    `xml:"coverArt,attr,omitempty"`
	SongCount int       `xml:"songCount,attr"`
	Duration  int       `xml:"duration,attr"`
	PlayCount int64     `xml:"playCount,attr,omitempty"`
	Created   time.Time `xml:"created,attr"`
	Starred   time.Time `xml:"starred,attr,omitempty"`
	Year      int       `xml:"year,attr,omitempty"`
	Genre     string    `xml:"genre,attr,omitempty"`
}

// AlbumInfo is a collection of notes and links describing an album.
type AlbumInfo struct {
	Notes          string `xml:"http://subsonic.org/restapi notes,omitempty"`
	MusicBrainzID  string `xml:"http://subsonic.org/restapi musicBrainzId,omitempty"`
	LastFmUrl      string `xml:"http://subsonic.org/restapi lastFmUrl,omitempty"`
	SmallImageUrl  string `xml:"http://subsonic.org/restapi smallImageUrl,omitempty"`
	MediumImageUrl string `xml:"http://subsonic.org/restapi mediumImageUrl,omitempty"`
	LargeImageUrl  string `xml:"http://subsonic.org/restapi largeImageUrl,omitempty"`
}

type albumList struct {
	Album []*Child `xml:"http://subsonic.org/restapi album,omitempty"`
}

type albumList2 struct {
	Album []*AlbumID3 `xml:"http://subsonic.org/restapi album,omitempty"`
}

// Artist is an artist from the server, organized in the folders pattern.
type Artist struct {
	ID             string    `xml:"id,attr"`
	Name           string    `xml:"name,attr"`
	ArtistImageUrl string    `xml:"artistImageUrl,attr,omitempty"`
	Starred        time.Time `xml:"starred,attr,omitempty"`
	UserRating     int       `xml:"userRating,attr,omitempty"`
	AverageRating  float64   `xml:"averageRating,attr,omitempty"`
}

// ArtistID3 is an artist from the server, organized by ID3 tag.
type ArtistID3 struct {
	ID             string      `xml:"id,attr"`                                     // Manually added
	Album          []*AlbumID3 `xml:"http://subsonic.org/restapi album,omitempty"` // Merged with ArtistWithAlbumsID3
	Name           string      `xml:"name,attr"`
	CoverArt       string      `xml:"coverArt,attr,omitempty"`
	ArtistImageUrl string      `xml:"artistImageUrl,attr,omitempty"`
	AlbumCount     int         `xml:"albumCount,attr"`
	Starred        time.Time   `xml:"starred,attr,omitempty"`
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo.
type ArtistInfo struct {
	SimilarArtist  []*Artist `xml:"http://subsonic.org/restapi similarArtist,omitempty"`
	Biography      string    `xml:"http://subsonic.org/restapi biography,omitempty"`
	MusicBrainzID  string    `xml:"http://subsonic.org/restapi musicBrainzId,omitempty"`
	LastFmUrl      string    `xml:"http://subsonic.org/restapi lastFmUrl,omitempty"`
	SmallImageUrl  string    `xml:"http://subsonic.org/restapi smallImageUrl,omitempty"`
	MediumImageUrl string    `xml:"http://subsonic.org/restapi mediumImageUrl,omitempty"`
	LargeImageUrl  string    `xml:"http://subsonic.org/restapi largeImageUrl,omitempty"`
}

// ArtistInfo2 is all auxillary information about an artist from GetArtistInfo2, with similar artists organized by ID3 tags.
type ArtistInfo2 struct {
	SimilarArtist  []*ArtistID3 `xml:"http://subsonic.org/restapi similarArtist,omitempty"`
	Biography      string       `xml:"http://subsonic.org/restapi biography,omitempty"`
	MusicBrainzID  string       `xml:"http://subsonic.org/restapi musicBrainzId,omitempty"`
	LastFmUrl      string       `xml:"http://subsonic.org/restapi lastFmUrl,omitempty"`
	SmallImageUrl  string       `xml:"http://subsonic.org/restapi smallImageUrl,omitempty"`
	MediumImageUrl string       `xml:"http://subsonic.org/restapi mediumImageUrl,omitempty"`
	LargeImageUrl  string       `xml:"http://subsonic.org/restapi largeImageUrl,omitempty"`
}

// ArtistsID3 is an index of every artist on the server organized by ID3 tag, from getArtists.
type ArtistsID3 struct {
	Index           []*IndexID3 `xml:"http://subsonic.org/restapi index,omitempty"`
	IgnoredArticles string      `xml:"ignoredArticles,attr"`
}

type Bookmark struct {
	Entry    *Child    `xml:"http://subsonic.org/restapi entry"`
	Position int64     `xml:"position,attr"`
	Username string    `xml:"username,attr"`
	Comment  string    `xml:"comment,attr,omitempty"`
	Created  time.Time `xml:"created,attr"`
	Changed  time.Time `xml:"changed,attr"`
}

func (t *Bookmark) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Bookmark
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *Bookmark) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Bookmark
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}

type bookmarks struct {
	Bookmark []*Bookmark `xml:"http://subsonic.org/restapi bookmark,omitempty"`
}

type ChatMessage struct {
	Username string `xml:"username,attr"`
	Time     int64  `xml:"time,attr"`
	Message  string `xml:"message,attr"`
}

type chatMessages struct {
	ChatMessage []*ChatMessage `xml:"http://subsonic.org/restapi chatMessage,omitempty"`
}

// Child is a song, or a generic entry in the hierarchical directory structure of the database.
// You can tell if Child is used as a song contextually based on what it was returned by, or if the IsDir boolean was set to true.
type Child struct {
	ID                    string    `xml:"id,attr"` // Manually added
	Parent                string    `xml:"parent,attr,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"`
	Title                 string    `xml:"title,attr"`
	Album                 string    `xml:"album,attr,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"` // May be one of music, podcast, audiobook, video
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"`
}

// Directory is an entry in the hierarchical folder structure organization of the server database.
type Directory struct {
	ID            string    `xml:"id,attr"` // Manually added
	Child         []*Child  `xml:"http://subsonic.org/restapi child,omitempty"`
	Parent        string    `xml:"parent,attr,omitempty"`
	Name          string    `xml:"name,attr"`
	Starred       time.Time `xml:"starred,attr,omitempty"`
	UserRating    int       `xml:"userRating,attr,omitempty"`
	AverageRating float64   `xml:"averageRating,attr,omitempty"`
	PlayCount     int64     `xml:"playCount,attr,omitempty"`
}

type Error struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:"message,attr,omitempty"`
}

// Genre is a style tag for a collection of songs and albums.
type Genre struct {
	Name       string `xml:",chardata"` // Added manually
	SongCount  int    `xml:"songCount,attr"`
	AlbumCount int    `xml:"albumCount,attr"`
}

type genres struct {
	Genre []*Genre `xml:"http://subsonic.org/restapi genre,omitempty"`
}

// Index is a collection of artists that begin with the same first letter, along with that letter or category.
type Index struct {
	Artist []*Artist `xml:"http://subsonic.org/restapi artist,omitempty"`
	Name   string    `xml:"name,attr"`
}

// Index is a collection of artists by ID3 tag that begin with the same first letter, along with that letter or category.
type IndexID3 struct {
	Artist []*ArtistID3 `xml:"http://subsonic.org/restapi artist,omitempty"`
	Name   string       `xml:"name,attr"`
}

// Indexes is the full index of the database, returned by getIndex.
// It contains some Index structs for each letter of the DB, plus Child entries for individual tracks.
type Indexes struct {
	Shortcut        []*Artist `xml:"http://subsonic.org/restapi shortcut,omitempty"`
	Index           []*Index  `xml:"http://subsonic.org/restapi index,omitempty"`
	Child           []*Child  `xml:"http://subsonic.org/restapi child,omitempty"`
	LastModified    int64     `xml:"lastModified,attr"`
	IgnoredArticles string    `xml:"ignoredArticles,attr"`
}

type InternetRadioStation struct {
	Name        string `xml:"name,attr"`
	StreamUrl   string `xml:"streamUrl,attr"`
	HomePageUrl string `xml:"homePageUrl,attr,omitempty"`
}

type internetRadioStations struct {
	InternetRadioStation []*InternetRadioStation `xml:"http://subsonic.org/restapi internetRadioStation,omitempty"`
}

type JukeboxPlaylist struct {
	Entry        []*Child `xml:"http://subsonic.org/restapi entry,omitempty"`
	CurrentIndex int      `xml:"currentIndex,attr"`
	Playing      bool     `xml:"playing,attr"`
	Gain         float32  `xml:"gain,attr"`
	Position     int      `xml:"position,attr,omitempty"`
}

type JukeboxStatus struct {
	CurrentIndex int     `xml:"currentIndex,attr"`
	Playing      bool    `xml:"playing,attr"`
	Gain         float32 `xml:"gain,attr"`
	Position     int     `xml:"position,attr,omitempty"`
}

// License contains information about the Subsonic server's license validity and contact information in the case of a trial subscription.
type License struct {
	Valid          bool      `xml:"valid,attr"`
	Email          string    `xml:"email,attr,omitempty"`
	LicenseExpires time.Time `xml:"licenseExpires,attr,omitempty"`
	TrialExpires   time.Time `xml:"trialExpires,attr,omitempty"`
}

type Lyrics struct {
	Artist string `xml:"artist,attr,omitempty"`
	Title  string `xml:"title,attr,omitempty"`
}

// MusicFolder is a representation of a source of music files added to the server. It is identified primarily by the numeric ID.
type MusicFolder struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr,omitempty"`
}

type musicFolders struct {
	MusicFolder []*MusicFolder `xml:"http://subsonic.org/restapi musicFolder,omitempty"`
}

type newestPodcasts struct {
	Episode []*PodcastEpisode `xml:"http://subsonic.org/restapi episode,omitempty"`
}

type nowPlaying struct {
	Entry []*NowPlayingEntry `xml:"http://subsonic.org/restapi entry,omitempty"`
}

// NowPlayingEntry is one individual stream coming from the server along with information about who was streaming it.
type NowPlayingEntry struct {
	Username              string    `xml:"username,attr"`
	MinutesAgo            int       `xml:"minutesAgo,attr"`
	PlayerID              int       `xml:"playerId,attr"`
	PlayerName            string    `xml:"playerName,attr,omitempty"`
	Parent                string    `xml:"parent,attr,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"`
	Title                 string    `xml:"title,attr"`
	Album                 string    `xml:"album,attr,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"` // May be one of music, podcast, audiobook, video
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"`
}

type PlayQueue struct {
	Entry     []*Child  `xml:"http://subsonic.org/restapi entry,omitempty"`
	Current   int       `xml:"current,attr,omitempty"`
	Position  int64     `xml:"position,attr,omitempty"`
	Username  string    `xml:"username,attr"`
	Changed   time.Time `xml:"changed,attr"`
	ChangedBy string    `xml:"changedBy,attr"`
}

func (t *PlayQueue) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PlayQueue
	var layout struct {
		*T
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	layout.T = (*T)(t)
	layout.Changed = (*xsdDateTime)(&layout.T.Changed)
	return e.EncodeElement(layout, start)
}
func (t *PlayQueue) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PlayQueue
	var overlay struct {
		*T
		Changed *xsdDateTime `xml:"changed,attr"`
	}
	overlay.T = (*T)(t)
	overlay.Changed = (*xsdDateTime)(&overlay.T.Changed)
	return d.DecodeElement(&overlay, &start)
}

// Playlist is a collection of songs with metadata like a name, comment, and information about the total duration of the playlist.
type Playlist struct {
	ID          string    `xml:"id,attr"`                                     // Added manually
	Entry       []*Child  `xml:"http://subsonic.org/restapi entry,omitempty"` // Merged from PlaylistWithSongs
	AllowedUser []string  `xml:"http://subsonic.org/restapi allowedUser,omitempty"`
	Name        string    `xml:"name,attr"`
	Comment     string    `xml:"comment,attr,omitempty"`
	Owner       string    `xml:"owner,attr,omitempty"`
	Public      bool      `xml:"public,attr,omitempty"`
	SongCount   int       `xml:"songCount,attr"`
	Duration    int       `xml:"duration,attr"`
	Created     time.Time `xml:"created,attr"`
	Changed     time.Time `xml:"changed,attr"`
	CoverArt    string    `xml:"coverArt,attr,omitempty"`
}

type playlists struct {
	Playlist []*Playlist `xml:"http://subsonic.org/restapi playlist,omitempty"`
}

type PodcastChannel struct {
	Episode          []*PodcastEpisode `xml:"http://subsonic.org/restapi episode,omitempty"`
	Url              string            `xml:"url,attr"`
	Title            string            `xml:"title,attr,omitempty"`
	Description      string            `xml:"description,attr,omitempty"`
	CoverArt         string            `xml:"coverArt,attr,omitempty"`
	OriginalImageUrl string            `xml:"originalImageUrl,attr,omitempty"`
	Status           string            `xml:"status,attr"` // May be one of new, downloading, completed, error, deleted, skipped
	ErrorMessage     string            `xml:"errorMessage,attr,omitempty"`
}

type PodcastEpisode struct {
	StreamID              string    `xml:"streamId,attr,omitempty"`
	ChannelID             string    `xml:"channelId,attr"`
	Description           string    `xml:"description,attr,omitempty"`
	Status                string    `xml:"status,attr"` // May be one of new, downloading, completed, error, deleted, skipped
	PublishDate           time.Time `xml:"publishDate,attr,omitempty"`
	Parent                string    `xml:"parent,attr,omitempty"`
	IsDir                 bool      `xml:"isDir,attr"`
	Title                 string    `xml:"title,attr"`
	Album                 string    `xml:"album,attr,omitempty"`
	Artist                string    `xml:"artist,attr,omitempty"`
	Track                 int       `xml:"track,attr,omitempty"`
	Year                  int       `xml:"year,attr,omitempty"`
	Genre                 string    `xml:"genre,attr,omitempty"`
	CoverArt              string    `xml:"coverArt,attr,omitempty"`
	Size                  int64     `xml:"size,attr,omitempty"`
	ContentType           string    `xml:"contentType,attr,omitempty"`
	Suffix                string    `xml:"suffix,attr,omitempty"`
	TranscodedContentType string    `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string    `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int       `xml:"duration,attr,omitempty"`
	BitRate               int       `xml:"bitRate,attr,omitempty"`
	Path                  string    `xml:"path,attr,omitempty"`
	IsVideo               bool      `xml:"isVideo,attr,omitempty"`
	UserRating            int       `xml:"userRating,attr,omitempty"`
	AverageRating         float64   `xml:"averageRating,attr,omitempty"`
	PlayCount             int64     `xml:"playCount,attr,omitempty"`
	DiscNumber            int       `xml:"discNumber,attr,omitempty"`
	Created               time.Time `xml:"created,attr,omitempty"`
	Starred               time.Time `xml:"starred,attr,omitempty"`
	AlbumID               string    `xml:"albumId,attr,omitempty"`
	ArtistID              string    `xml:"artistId,attr,omitempty"`
	Type                  string    `xml:"type,attr,omitempty"` // May be one of music, podcast, audiobook, video
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"`
}

func (t *PodcastEpisode) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PodcastEpisode
	var layout struct {
		*T
		PublishDate *xsdDateTime `xml:"publishDate,attr,omitempty"`
		Created     *xsdDateTime `xml:"created,attr,omitempty"`
		Starred     *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.PublishDate = (*xsdDateTime)(&layout.T.PublishDate)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *PodcastEpisode) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PodcastEpisode
	var overlay struct {
		*T
		PublishDate *xsdDateTime `xml:"publishDate,attr,omitempty"`
		Created     *xsdDateTime `xml:"created,attr,omitempty"`
		Starred     *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.PublishDate = (*xsdDateTime)(&overlay.T.PublishDate)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

type podcasts struct {
	Channel []*PodcastChannel `xml:"http://subsonic.org/restapi channel,omitempty"`
}

// Response is the main target for unmarshalling data from the API - everything within the "subsonic-response" key
type Response struct {
	License               *License               `xml:"http://subsonic.org/restapi license"`
	MusicFolders          *musicFolders          `xml:"http://subsonic.org/restapi musicFolders"`
	Indexes               *Indexes               `xml:"http://subsonic.org/restapi indexes"`
	Directory             *Directory             `xml:"http://subsonic.org/restapi directory"`
	Genres                *genres                `xml:"http://subsonic.org/restapi genres"`
	Artists               *ArtistsID3            `xml:"http://subsonic.org/restapi artists"`
	Artist                *ArtistID3             `xml:"http://subsonic.org/restapi artist"`
	Album                 *AlbumID3              `xml:"http://subsonic.org/restapi album"`
	Song                  *Child                 `xml:"http://subsonic.org/restapi song"`
	NowPlaying            *nowPlaying            `xml:"http://subsonic.org/restapi nowPlaying"`
	SearchResult2         *SearchResult2         `xml:"http://subsonic.org/restapi searchResult2"`
	SearchResult3         *SearchResult3         `xml:"http://subsonic.org/restapi searchResult3"`
	Playlists             *playlists             `xml:"http://subsonic.org/restapi playlists"`
	Playlist              *Playlist              `xml:"http://subsonic.org/restapi playlist"`
	JukeboxStatus         *JukeboxStatus         `xml:"http://subsonic.org/restapi jukeboxStatus"`
	JukeboxPlaylist       *JukeboxPlaylist       `xml:"http://subsonic.org/restapi jukeboxPlaylist"`
	Users                 *users                 `xml:"http://subsonic.org/restapi users"`
	User                  *User                  `xml:"http://subsonic.org/restapi user"`
	ChatMessages          *chatMessages          `xml:"http://subsonic.org/restapi chatMessages"`
	AlbumList             *albumList             `xml:"http://subsonic.org/restapi albumList"`
	AlbumList2            *albumList2            `xml:"http://subsonic.org/restapi albumList2"`
	RandomSongs           *songs                 `xml:"http://subsonic.org/restapi randomSongs"`
	SongsByGenre          *songs                 `xml:"http://subsonic.org/restapi songsByGenre"`
	Lyrics                *Lyrics                `xml:"http://subsonic.org/restapi lyrics"`
	Podcasts              *podcasts              `xml:"http://subsonic.org/restapi podcasts"`
	NewestPodcasts        *newestPodcasts        `xml:"http://subsonic.org/restapi newestPodcasts"`
	InternetRadioStations *internetRadioStations `xml:"http://subsonic.org/restapi internetRadioStations"`
	Bookmarks             *bookmarks             `xml:"http://subsonic.org/restapi bookmarks"`
	PlayQueue             *PlayQueue             `xml:"http://subsonic.org/restapi playQueue"`
	Shares                *shares                `xml:"http://subsonic.org/restapi shares"`
	Starred               *Starred               `xml:"http://subsonic.org/restapi starred"`
	Starred2              *Starred2              `xml:"http://subsonic.org/restapi starred2"`
	AlbumInfo             *AlbumInfo             `xml:"http://subsonic.org/restapi albumInfo"`
	ArtistInfo            *ArtistInfo            `xml:"http://subsonic.org/restapi artistInfo"`
	ArtistInfo2           *ArtistInfo2           `xml:"http://subsonic.org/restapi artistInfo2"`
	SimilarSongs          *similarSongs          `xml:"http://subsonic.org/restapi similarSongs"`
	SimilarSongs2         *similarSongs2         `xml:"http://subsonic.org/restapi similarSongs2"`
	TopSongs              *topSongs              `xml:"http://subsonic.org/restapi topSongs"`
	ScanStatus            *ScanStatus            `xml:"http://subsonic.org/restapi scanStatus"`
	Error                 *Error                 `xml:"http://subsonic.org/restapi error"`
	Status                string                 `xml:"status,attr"`  // May be one of ok, failed
	Version               string                 `xml:"version,attr"` // Must match the pattern \d+\.\d+\.\d+
}

type ScanStatus struct {
	Scanning bool  `xml:"scanning,attr"`
	Count    int64 `xml:"count,attr,omitempty"`
}

// SearchResult2 is a collection of songs, albums, and artists related to a query.
type SearchResult2 struct {
	Artist []*Artist `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*Child  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child  `xml:"http://subsonic.org/restapi song,omitempty"`
}

// SearchResult3 is a collection of songs, albums, and artists related to a query.
type SearchResult3 struct {
	Artist []*ArtistID3 `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*AlbumID3  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child     `xml:"http://subsonic.org/restapi song,omitempty"`
}

type Share struct {
	Entry       []*Child  `xml:"http://subsonic.org/restapi entry,omitempty"`
	Url         string    `xml:"url,attr"`
	Description string    `xml:"description,attr,omitempty"`
	Username    string    `xml:"username,attr"`
	Created     time.Time `xml:"created,attr"`
	Expires     time.Time `xml:"expires,attr,omitempty"`
	LastVisited time.Time `xml:"lastVisited,attr,omitempty"`
	VisitCount  int       `xml:"visitCount,attr"`
}

func (t *Share) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Share
	var layout struct {
		*T
		Created     *xsdDateTime `xml:"created,attr"`
		Expires     *xsdDateTime `xml:"expires,attr,omitempty"`
		LastVisited *xsdDateTime `xml:"lastVisited,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Expires = (*xsdDateTime)(&layout.T.Expires)
	layout.LastVisited = (*xsdDateTime)(&layout.T.LastVisited)
	return e.EncodeElement(layout, start)
}
func (t *Share) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Share
	var overlay struct {
		*T
		Created     *xsdDateTime `xml:"created,attr"`
		Expires     *xsdDateTime `xml:"expires,attr,omitempty"`
		LastVisited *xsdDateTime `xml:"lastVisited,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Expires = (*xsdDateTime)(&overlay.T.Expires)
	overlay.LastVisited = (*xsdDateTime)(&overlay.T.LastVisited)
	return d.DecodeElement(&overlay, &start)
}

type shares struct {
	Share []*Share `xml:"http://subsonic.org/restapi share,omitempty"`
}

type similarSongs struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

type similarSongs2 struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

type songs struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

// Starred is a collection of songs, albums, and artists annotated by a user as starred.
type Starred struct {
	Artist []*Artist `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*Child  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child  `xml:"http://subsonic.org/restapi song,omitempty"`
}

// Starred2 is a collection of songs, albums, and artists organized by ID3 tags annotated by a user as starred.
type Starred2 struct {
	Artist []*ArtistID3 `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*AlbumID3  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child     `xml:"http://subsonic.org/restapi song,omitempty"`
}

type topSongs struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

type User struct {
	Folder              []int     `xml:"http://subsonic.org/restapi folder,omitempty"`
	Username            string    `xml:"username,attr"`
	Email               string    `xml:"email,attr,omitempty"`
	ScrobblingEnabled   bool      `xml:"scrobblingEnabled,attr"`
	MaxBitRate          int       `xml:"maxBitRate,attr,omitempty"`
	AdminRole           bool      `xml:"adminRole,attr"`
	SettingsRole        bool      `xml:"settingsRole,attr"`
	DownloadRole        bool      `xml:"downloadRole,attr"`
	UploadRole          bool      `xml:"uploadRole,attr"`
	PlaylistRole        bool      `xml:"playlistRole,attr"`
	CoverArtRole        bool      `xml:"coverArtRole,attr"`
	CommentRole         bool      `xml:"commentRole,attr"`
	PodcastRole         bool      `xml:"podcastRole,attr"`
	StreamRole          bool      `xml:"streamRole,attr"`
	JukeboxRole         bool      `xml:"jukeboxRole,attr"`
	ShareRole           bool      `xml:"shareRole,attr"`
	VideoConversionRole bool      `xml:"videoConversionRole,attr"`
	AvatarLastChanged   time.Time `xml:"avatarLastChanged,attr,omitempty"`
}

func (t *User) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T User
	var layout struct {
		*T
		AvatarLastChanged *xsdDateTime `xml:"avatarLastChanged,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.AvatarLastChanged = (*xsdDateTime)(&layout.T.AvatarLastChanged)
	return e.EncodeElement(layout, start)
}
func (t *User) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T User
	var overlay struct {
		*T
		AvatarLastChanged *xsdDateTime `xml:"avatarLastChanged,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.AvatarLastChanged = (*xsdDateTime)(&overlay.T.AvatarLastChanged)
	return d.DecodeElement(&overlay, &start)
}

type users struct {
	User []*User `xml:"http://subsonic.org/restapi user,omitempty"`
}

type xsdDateTime time.Time

func (t *xsdDateTime) UnmarshalText(text []byte) error {
	return _unmarshalTime(text, (*time.Time)(t), "2006-01-02T15:04:05.999999999")
}
func (t xsdDateTime) MarshalText() ([]byte, error) {
	return []byte((time.Time)(t).Format("2006-01-02T15:04:05.999999999")), nil
}
func (t xsdDateTime) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if (time.Time)(t).IsZero() {
		return nil
	}
	m, err := t.MarshalText()
	if err != nil {
		return err
	}
	return e.EncodeElement(m, start)
}
func (t xsdDateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	if (time.Time)(t).IsZero() {
		return xml.Attr{}, nil
	}
	m, err := t.MarshalText()
	return xml.Attr{Name: name, Value: string(m)}, err
}
func _unmarshalTime(text []byte, t *time.Time, format string) (err error) {
	s := string(bytes.TrimSpace(text))
	*t, err = time.Parse(format, s)
	if _, ok := err.(*time.ParseError); ok {
		*t, err = time.Parse(format+"Z07:00", s)
	}
	return err
}
