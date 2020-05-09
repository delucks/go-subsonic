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
