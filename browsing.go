package main

import "log"

func (s *SubsonicClient) GetMusicFolders() *SubsonicResponse {
	resp, err := s.Get("getMusicFolders", nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp
}
