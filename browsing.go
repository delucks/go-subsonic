package main

import (
	"time"
)

type MusicFolder struct {
	Id   int    `json:"id"` // subsonic returns an int, navidrome a string
	Name string `json:"name"`
}

type MusicFolderContainer struct {
	Folders []*MusicFolder `json:"musicFolder"`
}

func (s *SubsonicClient) GetMusicFolders() ([]*MusicFolder, error) {
	resp, err := s.Get("getMusicFolders", nil)
	if err != nil {
		return nil, err
	}
	return resp.MusicFolders.Folders, nil
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
	Artist    string    `json:"artist"`
	ArtistID  string    `json:"artistId"`
	SongCount int       `json:"songCount"`
	Duration  int       `json:"duration"`
	Created   time.Time `json:"created"`
	Year      int       `json:"year"`
	Genre     string    `json:"genre"`
	PlayCount int       `json:"playCount"`
	CoverArt  string    `json:"coverArt"`
	Songs     []*Song   `json:"song"`    // populated by getAlbum
	IsDir     bool      `json:"isDir"`   // navidrome only
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

// Index contains a by-letter representation of every item in the database.
type Index struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artist"`
}

type IndexContainer struct {
	LastModified    int64   `json:"lastModified"` // subsonic returns an int64, navidrome a string
	IgnoredArticles string  `json:"ignoredArticles"`
	Indexes         []Index `json:"index"`
}

// GetIndexes returns the index of entries by letter/number.
// Optional Parameters:
// * musicFolderId   If specified, only return artists in the music folder with the given ID. See getMusicFolders.
// * ifModifiedSince If specified, only return a result if the artist collection has changed since the given time (in milliseconds since 1 Jan 1970).
func (s *SubsonicClient) GetIndexes(parameters map[string]string) (*IndexContainer, error) {
	resp, err := s.Get("getIndexes", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Indexes, nil
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

// GetMusicDirectory returns the context around an object ID from the database.
func (s *SubsonicClient) GetMusicDirectory(id string) (*Directory, error) {
	resp, err := s.Get("getMusicDirectory", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Directory, nil
}

type Genre struct {
	SongCount  int    `json:"songCount"`
	AlbumCount int    `json:"albumCount"`
	Value      string `json:"value"`
}

type GenreContainer struct {
	Genre []*Genre `json:"genre"`
}

// GetGenres returns all genres in the server.
func (s *SubsonicClient) GetGenres() ([]*Genre, error) {
	resp, err := s.Get("getGenres", nil)
	if err != nil {
		return nil, err
	}
	return resp.Genres.Genre, nil
}

type ArtistsContainer struct {
	IgnoredArticles string  `json:"ignoredArticles"`
	Indexes         []Index `json:"index"`
}

// GetArtists returns all artists in the server. If specified, musicFolderId will return only the artists from a specific folder.
func (s *SubsonicClient) GetArtists(parameters map[string]string) (*ArtistsContainer, error) {
	resp, err := s.Get("getArtists", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Artists, nil
}

// GetAlbum returns an Artist by ID.
func (s *SubsonicClient) GetArtist(id string) (*Artist, error) {
	resp, err := s.Get("getArtist", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Artist, nil
}

// GetAlbum returns an Album by ID.
func (s *SubsonicClient) GetAlbum(id string) (*Album, error) {
	resp, err := s.Get("getAlbum", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Album, nil
}

// GetSong returns a Song by ID.
func (s *SubsonicClient) GetSong(id string) (*Song, error) {
	resp, err := s.Get("getSong", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Song, nil
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

// GetArtistInfo returns biography, image links, and similar artists from last.fm.
// Optional Parameters:
// * count:             Max number of similar artists to return.
// * includeNotPresent: Whether to return artists that are not present in the media library.
func (s *SubsonicClient) GetArtistInfo(id string, parameters map[string]string) (*ArtistInfo, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getArtistInfo", params)
	if err != nil {
		return nil, err
	}
	return resp.ArtistInfo, nil
}

// GetArtistInfo2 returns biography, image links, and similar artists like GetArtistInfo, but using id3 tags.
// Optional Parameters:
// * count:             Max number of similar artists to return.
// * includeNotPresent: Whether to return artists that are not present in the media library.
func (s *SubsonicClient) GetArtistInfo2(id string, parameters map[string]string) (*ArtistInfo, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getArtistInfo2", params)
	if err != nil {
		return nil, err
	}
	return resp.ArtistInfo2, nil
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

// GetAlbumInfo returns album notes, image data, etc using data from last.fm.
// This accepts both album and song IDs.
func (s *SubsonicClient) GetAlbumInfo(id string) (*AlbumInfo, error) {
	resp, err := s.Get("getAlbumInfo", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.AlbumInfo, nil
}

// GetAlbumInfo2 returns the same data as GetAlbumInfo, but organized by id3 tag.
// It only accepts album IDs.
func (s *SubsonicClient) GetAlbumInfo2(id string) (*AlbumInfo, error) {
	resp, err := s.Get("getAlbumInfo2", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.AlbumInfo, nil
}

type SimilarSongs struct {
	Songs []*Song `json:"song"`
}

// GetSimilarSongs finds similar songs to an album, track, or artist.
// This is mostly used for radio features. This accepts artist, album, or song IDs.
// Optional Parameters:
// * count: Number of songs to return
func (s *SubsonicClient) GetSimilarSongs(id string, parameters map[string]string) ([]*Song, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSimilarSongs", params)
	if err != nil {
		return nil, err
	}
	return resp.SimilarSongs.Songs, nil
}

// GetSimilarSongs2 finds similar songs like GetSimilarSongs, but using id3 tags.
// Optional Parameters:
// * count: Number of songs to return
func (s *SubsonicClient) GetSimilarSongs2(id string, parameters map[string]string) ([]*Song, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSimilarSongs2", params)
	if err != nil {
		return nil, err
	}
	return resp.SimilarSongs2.Songs, nil
}
