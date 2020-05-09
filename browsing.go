package main

import "log"

type MusicFolder struct {
	Id   int    `json:"id"`
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
	LastModified    string  `json:"lastModified"` // subsonic returns an int64, navidrome a string
	IgnoredArticles string  `json:"ignoredArticles"`
	Indexes         []Index `json:"index"`
}

func (s *SubsonicClient) GetIndexes() *IndexContainer {
	resp, err := s.Get("getIndexes", nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.Indexes
}
