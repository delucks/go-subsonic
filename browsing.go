package main

import "log"

type MusicFolder struct {
	Id   int    `json:"id"` // subsonic returns an int, navidrome a string
	Name string `json:"name"`
}

type MusicFolderContainer struct {
	Folders []*MusicFolder `json:"musicFolder"`
}

func (s *SubsonicClient) GetMusicFolders() []*MusicFolder {
	resp, err := s.Get("getMusicFolders", nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.MusicFolders.Folders
}

type Artist struct {
	ID             string `json:"id"`
	Name           string `json:"name"`
	ArtistImageURL string `json:"artistImageUrl"`
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
func (s *SubsonicClient) GetIndexes(parameters map[string]string) *IndexContainer {
	resp, err := s.Get("getIndexes", parameters)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.Indexes
}

func (s *SubsonicClient) GetMusicDirectory(id string) *SubsonicResponse {
	resp, err := s.Get("getMusicDirectory", map[string]string{"id": id})
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp
}
