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

type Artist struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	AlbumCount     int    `json:"albumCount"`
	ArtistImageURL string `json:"artistImageUrl"` // subsonic only
	CoverArt       string `json:"coverArt"`       // subsonic only
}

type Index struct {
	Name    string   `json:"name"`
	Artists []Artist `json:"artist"`
}

type IndexContainer struct {
	LastModified    int64   `json:"lastModified"` // subsonic returns an int64, navidrome a string
	IgnoredArticles string  `json:"ignoredArticles"`
	Indexes         []Index `json:"index"`
}

/*
 * Parameters:
 *   musicFolderId   If specified, only return artists in the music folder with the given ID. See getMusicFolders.
 *   ifModifiedSince If specified, only return a result if the artist collection has changed since the given time (in milliseconds since 1 Jan 1970).
 */
func (s *SubsonicClient) GetIndexes(parameters map[string]string) (*IndexContainer, error) {
	resp, err := s.Get("getIndexes", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Indexes, nil
}

type Child struct {
	ID          string    `json:"id"`
	Parent      string    `json:"parent"`
	IsDir       bool      `json:"isDir"`
	Title       string    `json:"title"`
	Album       string    `json:"album"`
	CoverArt    string    `json:"coverArt"`
	Size        int64     `json:"size"` // string in navidrome
	ContentType string    `json:"contentType"`
	Suffix      string    `json:"suffix"`
	Duration    int       `json:"duration"`
	BitRate     int       `json:"bitRate"`
	Path        string    `json:"path"`
	IsVideo     bool      `json:"isVideo"`
	Created     time.Time `json:"created"`
	Type        string    `json:"type"`
	PlayCount   int       `json:"playCount"` // airsonic only
	Artist      string    `json:"artist"`    // this and all following fields are navidrome only
	Track       int       `json:"track"`
	Year        int       `json:"year"`
	Genre       string    `json:"genre"`
	DiscNumber  int       `json:"discNumber"`
	AlbumID     string    `json:"albumId"`
	ArtistID    string    `json:"artistId"`
}

type Directory struct {
	Children   []*Child `json:"child"`
	ID         string   `json:"id"`
	Name       string   `json:"name"`
	PlayCount  int      `json:"playCount"`  // airsonic only
	AlbumCount int      `json:"albumCount"` // navidrome only
	Parent     string   `json:"parent"`     // navidrome only
}

// The parameter is an object ID from the database
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

/*
 * Parameters:
 *   musicFolderId   If specified, only return artists in the music folder with the given ID. See getMusicFolders.
 */
func (s *SubsonicClient) GetArtists(parameters map[string]string) (*ArtistsContainer, error) {
	resp, err := s.Get("getArtists", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Artists, nil
}
