package subsonic

/* This file was automatically generated from the xsd schema provided by Subsonic, then manually modified.
 *   xsdgen -o xml.go -pkg subsonic -ns "http://subsonic.org/restapi" subsonic-rest-api-1.16.1.xsd
 * Changes from the original include:
 * - Adding missing name (value of xml element) for each genre
 * - Capitalize "ID" in struct names
 */

import (
	"bytes"
	"encoding/xml"
	"time"
)

// AlbumID3 is an album that's organized by ID3 tags.
type AlbumID3 struct {
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

func (t *AlbumID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AlbumID3
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *AlbumID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AlbumID3
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
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

type AlbumList struct {
	Album []*Child `xml:"http://subsonic.org/restapi album,omitempty"`
}

type AlbumList2 struct {
	Album []*AlbumID3 `xml:"http://subsonic.org/restapi album,omitempty"`
}

// AlbumWithSongsID3 is an Album organized by ID3 tags with songs, obtained by a call to getAlbum.
type AlbumWithSongsID3 struct {
	Song      []*Child  `xml:"http://subsonic.org/restapi song,omitempty"`
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

func (t *AlbumWithSongsID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T AlbumWithSongsID3
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *AlbumWithSongsID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T AlbumWithSongsID3
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

// Artist is a single artist from the database.
type Artist struct {
	Name           string    `xml:"name,attr"`
	ArtistImageUrl string    `xml:"artistImageUrl,attr,omitempty"`
	Starred        time.Time `xml:"starred,attr,omitempty"`
	UserRating     int       `xml:"userRating,attr,omitempty"`
	AverageRating  float64   `xml:"averageRating,attr,omitempty"`
}

func (t *Artist) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Artist
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Artist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Artist
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

// ArtistID3 is a single artist from the database, organized by ID3 tag.
type ArtistID3 struct {
	Name           string    `xml:"name,attr"`
	CoverArt       string    `xml:"coverArt,attr,omitempty"`
	ArtistImageUrl string    `xml:"artistImageUrl,attr,omitempty"`
	AlbumCount     int       `xml:"albumCount,attr"`
	Starred        time.Time `xml:"starred,attr,omitempty"`
}

func (t *ArtistID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ArtistID3
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *ArtistID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtistID3
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo/GetArtistInfo2
type ArtistInfo struct {
	SimilarArtist  []*Artist `xml:"http://subsonic.org/restapi similarArtist,omitempty"`
	Biography      string    `xml:"http://subsonic.org/restapi biography,omitempty"`
	MusicBrainzID  string    `xml:"http://subsonic.org/restapi musicBrainzId,omitempty"`
	LastFmUrl      string    `xml:"http://subsonic.org/restapi lastFmUrl,omitempty"`
	SmallImageUrl  string    `xml:"http://subsonic.org/restapi smallImageUrl,omitempty"`
	MediumImageUrl string    `xml:"http://subsonic.org/restapi mediumImageUrl,omitempty"`
	LargeImageUrl  string    `xml:"http://subsonic.org/restapi largeImageUrl,omitempty"`
}

// ArtistInfo is all auxillary information about an artist from GetArtistInfo/GetArtistInfo2
type ArtistInfo2 struct {
	SimilarArtist  []*ArtistID3 `xml:"http://subsonic.org/restapi similarArtist,omitempty"`
	Biography      string       `xml:"http://subsonic.org/restapi biography,omitempty"`
	MusicBrainzID  string       `xml:"http://subsonic.org/restapi musicBrainzId,omitempty"`
	LastFmUrl      string       `xml:"http://subsonic.org/restapi lastFmUrl,omitempty"`
	SmallImageUrl  string       `xml:"http://subsonic.org/restapi smallImageUrl,omitempty"`
	MediumImageUrl string       `xml:"http://subsonic.org/restapi mediumImageUrl,omitempty"`
	LargeImageUrl  string       `xml:"http://subsonic.org/restapi largeImageUrl,omitempty"`
}

// TODO is this type necessary?
type ArtistInfoBase struct {
	Biography      string `xml:"http://subsonic.org/restapi biography,omitempty"`
	MusicBrainzID  string `xml:"http://subsonic.org/restapi musicBrainzId,omitempty"`
	LastFmUrl      string `xml:"http://subsonic.org/restapi lastFmUrl,omitempty"`
	SmallImageUrl  string `xml:"http://subsonic.org/restapi smallImageUrl,omitempty"`
	MediumImageUrl string `xml:"http://subsonic.org/restapi mediumImageUrl,omitempty"`
	LargeImageUrl  string `xml:"http://subsonic.org/restapi largeImageUrl,omitempty"`
}

type ArtistWithAlbumsID3 struct {
	Album          []*AlbumID3 `xml:"http://subsonic.org/restapi album,omitempty"`
	Name           string      `xml:"name,attr"`
	CoverArt       string      `xml:"coverArt,attr,omitempty"`
	ArtistImageUrl string      `xml:"artistImageUrl,attr,omitempty"`
	AlbumCount     int         `xml:"albumCount,attr"`
	Starred        time.Time   `xml:"starred,attr,omitempty"`
}

func (t *ArtistWithAlbumsID3) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T ArtistWithAlbumsID3
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *ArtistWithAlbumsID3) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T ArtistWithAlbumsID3
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

type ArtistsID3 struct {
	Index           []*IndexID3 `xml:"http://subsonic.org/restapi index,omitempty"`
	IgnoredArticles string      `xml:"ignoredArticles,attr"`
}

type AudioTrack struct {
	Name         string `xml:"name,attr,omitempty"`
	LanguageCode string `xml:"languageCode,attr,omitempty"`
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

type Bookmarks struct {
	Bookmark []*Bookmark `xml:"http://subsonic.org/restapi bookmark,omitempty"`
}

type Captions struct {
	Name string `xml:"name,attr,omitempty"`
}

type ChatMessage struct {
	Username string `xml:"username,attr"`
	Time     int64  `xml:"time,attr"`
	Message  string `xml:"message,attr"`
}

type ChatMessages struct {
	ChatMessage []*ChatMessage `xml:"http://subsonic.org/restapi chatMessage,omitempty"`
}

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
	Type                  MediaType `xml:"type,attr,omitempty"`
	BookmarkPosition      int64     `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int       `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int       `xml:"originalHeight,attr,omitempty"`
}

func (t *Child) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Child
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Child) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Child
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

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

func (t *Directory) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Directory
	var layout struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *Directory) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Directory
	var overlay struct {
		*T
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
}

type Error struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:"message,attr,omitempty"`
}

type Genre struct {
	Name       string `xml:",chardata"` // Added manually
	SongCount  int    `xml:"songCount,attr"`
	AlbumCount int    `xml:"albumCount,attr"`
}

type Genres struct {
	Genre []*Genre `xml:"http://subsonic.org/restapi genre,omitempty"`
}

type Index struct {
	Artist []*Artist `xml:"http://subsonic.org/restapi artist,omitempty"`
	Name   string    `xml:"name,attr"`
}

type IndexID3 struct {
	Artist []*ArtistID3 `xml:"http://subsonic.org/restapi artist,omitempty"`
	Name   string       `xml:"name,attr"`
}

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

type InternetRadioStations struct {
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

func (t *License) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T License
	var layout struct {
		*T
		LicenseExpires *xsdDateTime `xml:"licenseExpires,attr,omitempty"`
		TrialExpires   *xsdDateTime `xml:"trialExpires,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.LicenseExpires = (*xsdDateTime)(&layout.T.LicenseExpires)
	layout.TrialExpires = (*xsdDateTime)(&layout.T.TrialExpires)
	return e.EncodeElement(layout, start)
}
func (t *License) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T License
	var overlay struct {
		*T
		LicenseExpires *xsdDateTime `xml:"licenseExpires,attr,omitempty"`
		TrialExpires   *xsdDateTime `xml:"trialExpires,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.LicenseExpires = (*xsdDateTime)(&overlay.T.LicenseExpires)
	overlay.TrialExpires = (*xsdDateTime)(&overlay.T.TrialExpires)
	return d.DecodeElement(&overlay, &start)
}

type Lyrics struct {
	Artist string `xml:"artist,attr,omitempty"`
	Title  string `xml:"title,attr,omitempty"`
}

// May be one of music, podcast, audiobook, video
type MediaType string

// MusicFolder is a representation of a source of music files added to the server. It is identified primarily by the numeric ID.
type MusicFolder struct {
	ID   string `xml:"id,attr"`
	Name string `xml:"name,attr,omitempty"`
}

type MusicFolders struct {
	MusicFolder []*MusicFolder `xml:"http://subsonic.org/restapi musicFolder,omitempty"`
}

type NewestPodcasts struct {
	Episode []*PodcastEpisode `xml:"http://subsonic.org/restapi episode,omitempty"`
}

type NowPlaying struct {
	Entry []*NowPlayingEntry `xml:"http://subsonic.org/restapi entry,omitempty"`
}

type NowPlayingEntry struct {
	Username              string     `xml:"username,attr"`
	MinutesAgo            int        `xml:"minutesAgo,attr"`
	PlayerID              int        `xml:"playerId,attr"`
	PlayerName            string     `xml:"playerName,attr,omitempty"`
	Parent                string     `xml:"parent,attr,omitempty"`
	IsDir                 bool       `xml:"isDir,attr"`
	Title                 string     `xml:"title,attr"`
	Album                 string     `xml:"album,attr,omitempty"`
	Artist                string     `xml:"artist,attr,omitempty"`
	Track                 int        `xml:"track,attr,omitempty"`
	Year                  int        `xml:"year,attr,omitempty"`
	Genre                 string     `xml:"genre,attr,omitempty"`
	CoverArt              string     `xml:"coverArt,attr,omitempty"`
	Size                  int64      `xml:"size,attr,omitempty"`
	ContentType           string     `xml:"contentType,attr,omitempty"`
	Suffix                string     `xml:"suffix,attr,omitempty"`
	TranscodedContentType string     `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string     `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int        `xml:"duration,attr,omitempty"`
	BitRate               int        `xml:"bitRate,attr,omitempty"`
	Path                  string     `xml:"path,attr,omitempty"`
	IsVideo               bool       `xml:"isVideo,attr,omitempty"`
	UserRating            int        `xml:"userRating,attr,omitempty"`
	AverageRating         float64    `xml:"averageRating,attr,omitempty"`
	PlayCount             int64      `xml:"playCount,attr,omitempty"`
	DiscNumber            int        `xml:"discNumber,attr,omitempty"`
	Created               time.Time  `xml:"created,attr,omitempty"`
	Starred               time.Time  `xml:"starred,attr,omitempty"`
	AlbumID               string     `xml:"albumId,attr,omitempty"`
	ArtistID              string     `xml:"artistId,attr,omitempty"`
	Type                  *MediaType `xml:"type,attr,omitempty"`
	BookmarkPosition      int64      `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int        `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int        `xml:"originalHeight,attr,omitempty"`
}

func (t *NowPlayingEntry) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T NowPlayingEntry
	var layout struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	layout.T = (*T)(t)
	layout.Created = (*xsdDateTime)(&layout.T.Created)
	layout.Starred = (*xsdDateTime)(&layout.T.Starred)
	return e.EncodeElement(layout, start)
}
func (t *NowPlayingEntry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T NowPlayingEntry
	var overlay struct {
		*T
		Created *xsdDateTime `xml:"created,attr,omitempty"`
		Starred *xsdDateTime `xml:"starred,attr,omitempty"`
	}
	overlay.T = (*T)(t)
	overlay.Created = (*xsdDateTime)(&overlay.T.Created)
	overlay.Starred = (*xsdDateTime)(&overlay.T.Starred)
	return d.DecodeElement(&overlay, &start)
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

type Playlist struct {
	ID          string    `xml:"id,attr"` // Added manually
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

func (t *Playlist) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T Playlist
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
func (t *Playlist) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T Playlist
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

type PlaylistWithSongs struct {
	ID          string    `xml:"id,attr"` // Manually added
	Entry       []*Child  `xml:"http://subsonic.org/restapi entry,omitempty"`
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

func (t *PlaylistWithSongs) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	type T PlaylistWithSongs
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
func (t *PlaylistWithSongs) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type T PlaylistWithSongs
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

type Playlists struct {
	Playlist []*Playlist `xml:"http://subsonic.org/restapi playlist,omitempty"`
}

type PodcastChannel struct {
	Episode          []*PodcastEpisode `xml:"http://subsonic.org/restapi episode,omitempty"`
	Url              string            `xml:"url,attr"`
	Title            string            `xml:"title,attr,omitempty"`
	Description      string            `xml:"description,attr,omitempty"`
	CoverArt         string            `xml:"coverArt,attr,omitempty"`
	OriginalImageUrl string            `xml:"originalImageUrl,attr,omitempty"`
	Status           *PodcastStatus    `xml:"status,attr"`
	ErrorMessage     string            `xml:"errorMessage,attr,omitempty"`
}

type PodcastEpisode struct {
	StreamID              string         `xml:"streamId,attr,omitempty"`
	ChannelID             string         `xml:"channelId,attr"`
	Description           string         `xml:"description,attr,omitempty"`
	Status                *PodcastStatus `xml:"status,attr"`
	PublishDate           time.Time      `xml:"publishDate,attr,omitempty"`
	Parent                string         `xml:"parent,attr,omitempty"`
	IsDir                 bool           `xml:"isDir,attr"`
	Title                 string         `xml:"title,attr"`
	Album                 string         `xml:"album,attr,omitempty"`
	Artist                string         `xml:"artist,attr,omitempty"`
	Track                 int            `xml:"track,attr,omitempty"`
	Year                  int            `xml:"year,attr,omitempty"`
	Genre                 string         `xml:"genre,attr,omitempty"`
	CoverArt              string         `xml:"coverArt,attr,omitempty"`
	Size                  int64          `xml:"size,attr,omitempty"`
	ContentType           string         `xml:"contentType,attr,omitempty"`
	Suffix                string         `xml:"suffix,attr,omitempty"`
	TranscodedContentType string         `xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string         `xml:"transcodedSuffix,attr,omitempty"`
	Duration              int            `xml:"duration,attr,omitempty"`
	BitRate               int            `xml:"bitRate,attr,omitempty"`
	Path                  string         `xml:"path,attr,omitempty"`
	IsVideo               bool           `xml:"isVideo,attr,omitempty"`
	UserRating            int            `xml:"userRating,attr,omitempty"`
	AverageRating         float64        `xml:"averageRating,attr,omitempty"`
	PlayCount             int64          `xml:"playCount,attr,omitempty"`
	DiscNumber            int            `xml:"discNumber,attr,omitempty"`
	Created               time.Time      `xml:"created,attr,omitempty"`
	Starred               time.Time      `xml:"starred,attr,omitempty"`
	AlbumID               string         `xml:"albumId,attr,omitempty"`
	ArtistID              string         `xml:"artistId,attr,omitempty"`
	Type                  *MediaType     `xml:"type,attr,omitempty"`
	BookmarkPosition      int64          `xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int            `xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int            `xml:"originalHeight,attr,omitempty"`
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

// May be one of new, downloading, completed, error, deleted, skipped
type PodcastStatus string

type Podcasts struct {
	Channel []*PodcastChannel `xml:"http://subsonic.org/restapi channel,omitempty"`
}

// Response is the main target for unmarshalling JSON data from the API - everything within the "subsonic-response" key
type Response struct {
	MusicFolders          *MusicFolders          `xml:"http://subsonic.org/restapi musicFolders"`
	Indexes               *Indexes               `xml:"http://subsonic.org/restapi indexes"`
	Directory             *Directory             `xml:"http://subsonic.org/restapi directory"`
	Genres                *Genres                `xml:"http://subsonic.org/restapi genres"`
	Artists               *ArtistsID3            `xml:"http://subsonic.org/restapi artists"`
	Artist                *ArtistWithAlbumsID3   `xml:"http://subsonic.org/restapi artist"`
	Album                 *AlbumWithSongsID3     `xml:"http://subsonic.org/restapi album"`
	Song                  *Child                 `xml:"http://subsonic.org/restapi song"`
	Videos                *Videos                `xml:"http://subsonic.org/restapi videos"`
	VideoInfo             *VideoInfo             `xml:"http://subsonic.org/restapi videoInfo"`
	NowPlaying            *NowPlaying            `xml:"http://subsonic.org/restapi nowPlaying"`
	SearchResult          *SearchResult          `xml:"http://subsonic.org/restapi searchResult"`
	SearchResult2         *SearchResult2         `xml:"http://subsonic.org/restapi searchResult2"`
	SearchResult3         *SearchResult3         `xml:"http://subsonic.org/restapi searchResult3"`
	Playlists             *Playlists             `xml:"http://subsonic.org/restapi playlists"`
	Playlist              *PlaylistWithSongs     `xml:"http://subsonic.org/restapi playlist"`
	JukeboxStatus         *JukeboxStatus         `xml:"http://subsonic.org/restapi jukeboxStatus"`
	JukeboxPlaylist       *JukeboxPlaylist       `xml:"http://subsonic.org/restapi jukeboxPlaylist"`
	License               *License               `xml:"http://subsonic.org/restapi license"`
	Users                 *Users                 `xml:"http://subsonic.org/restapi users"`
	User                  *User                  `xml:"http://subsonic.org/restapi user"`
	ChatMessages          *ChatMessages          `xml:"http://subsonic.org/restapi chatMessages"`
	AlbumList             *AlbumList             `xml:"http://subsonic.org/restapi albumList"`
	AlbumList2            *AlbumList2            `xml:"http://subsonic.org/restapi albumList2"`
	RandomSongs           *Songs                 `xml:"http://subsonic.org/restapi randomSongs"`
	SongsByGenre          *Songs                 `xml:"http://subsonic.org/restapi songsByGenre"`
	Lyrics                *Lyrics                `xml:"http://subsonic.org/restapi lyrics"`
	Podcasts              *Podcasts              `xml:"http://subsonic.org/restapi podcasts"`
	NewestPodcasts        *NewestPodcasts        `xml:"http://subsonic.org/restapi newestPodcasts"`
	InternetRadioStations *InternetRadioStations `xml:"http://subsonic.org/restapi internetRadioStations"`
	Bookmarks             *Bookmarks             `xml:"http://subsonic.org/restapi bookmarks"`
	PlayQueue             *PlayQueue             `xml:"http://subsonic.org/restapi playQueue"`
	Shares                *Shares                `xml:"http://subsonic.org/restapi shares"`
	Starred               *Starred               `xml:"http://subsonic.org/restapi starred"`
	Starred2              *Starred2              `xml:"http://subsonic.org/restapi starred2"`
	AlbumInfo             *AlbumInfo             `xml:"http://subsonic.org/restapi albumInfo"`
	ArtistInfo            *ArtistInfo            `xml:"http://subsonic.org/restapi artistInfo"`
	ArtistInfo2           *ArtistInfo2           `xml:"http://subsonic.org/restapi artistInfo2"`
	SimilarSongs          *SimilarSongs          `xml:"http://subsonic.org/restapi similarSongs"`
	SimilarSongs2         *SimilarSongs2         `xml:"http://subsonic.org/restapi similarSongs2"`
	TopSongs              *TopSongs              `xml:"http://subsonic.org/restapi topSongs"`
	ScanStatus            *ScanStatus            `xml:"http://subsonic.org/restapi scanStatus"`
	Error                 *Error                 `xml:"http://subsonic.org/restapi error"`
	Status                ResponseStatus         `xml:"status,attr"`
	Version               Version                `xml:"version,attr"`
}

// May be one of ok, failed
type ResponseStatus string

type ScanStatus struct {
	Scanning bool  `xml:"scanning,attr"`
	Count    int64 `xml:"count,attr,omitempty"`
}

type SearchResult struct {
	Match     []*Child `xml:"http://subsonic.org/restapi match,omitempty"`
	TotalHits int      `xml:"totalHits,attr"`
}

type SearchResult2 struct {
	Artist []*Artist `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*Child  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child  `xml:"http://subsonic.org/restapi song,omitempty"`
}

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

type Shares struct {
	Share []*Share `xml:"http://subsonic.org/restapi share,omitempty"`
}

type SimilarSongs struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

type SimilarSongs2 struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

type Songs struct {
	Song []*Child `xml:"http://subsonic.org/restapi song,omitempty"`
}

type Starred struct {
	Artist []*Artist `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*Child  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child  `xml:"http://subsonic.org/restapi song,omitempty"`
}

type Starred2 struct {
	Artist []*ArtistID3 `xml:"http://subsonic.org/restapi artist,omitempty"`
	Album  []*AlbumID3  `xml:"http://subsonic.org/restapi album,omitempty"`
	Song   []*Child     `xml:"http://subsonic.org/restapi song,omitempty"`
}

type TopSongs struct {
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

type Users struct {
	User []*User `xml:"http://subsonic.org/restapi user,omitempty"`
}

// Must match the pattern \d+\.\d+\.\d+
type Version string

type VideoConversion struct {
	BitRate      int `xml:"bitRate,attr,omitempty"`
	AudioTrackID int `xml:"audioTrackId,attr,omitempty"`
}

type VideoInfo struct {
	Captions   []*Captions        `xml:"http://subsonic.org/restapi captions,omitempty"`
	AudioTrack []*AudioTrack      `xml:"http://subsonic.org/restapi audioTrack,omitempty"`
	Conversion []*VideoConversion `xml:"http://subsonic.org/restapi conversion,omitempty"`
}

type Videos struct {
	Video []*Child `xml:"http://subsonic.org/restapi video,omitempty"`
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
