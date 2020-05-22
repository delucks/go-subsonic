package subsonic

import (
	"errors"
	"fmt"
)

func validateListType(input string) bool {
	validTypes := map[string]bool{
		"random":               true,
		"newest":               true,
		"highest":              true,
		"frequent":             true,
		"recent":               true,
		"alphabeticalByName":   true,
		"alphabeticalByArtist": true,
		"starred":              true,
		"byYear":               true,
		"byGenre":              true,
	}
	_, ok := validTypes[input]
	return ok
}

// GetAlbumList returns a list of random, newest, highest rated etc. albums. Similar to the album lists on the home page of the Subsonic web interface.
//
// Optional Parameters:
//   size:           The number of albums to return. Max 500, default 10.
//   offset:         The list offset. Useful if you for example want to page through the list of newest albums.
//   fromYear:       The first year in the range. If fromYear > toYear a reverse chronological list is returned.
//   toYear:         The last year in the range.
//   genre:          The name of the genre, e.g., "Rock".
//   musicFolderId:  (Since 1.11.0) Only return albums in the music folder with the given ID. See getMusicFolders.
//
// toYear and fromYear are required parameters when type == "byYear". genre is a required parameter when type == "byGenre".
func (s *Client) GetAlbumList(listType string, parameters map[string]string) ([]*Child, error) {
	if !validateListType(listType) {
		return nil, fmt.Errorf("List type %s is invalid, see http://www.subsonic.org/pages/api.jsp#getAlbumList", listType)
	}
	if listType == "byYear" {
		_, ok := parameters["fromYear"]
		if !ok {
			return nil, errors.New("Required argument fromYear was not found when using GetAlbumList byYear")
		}
		_, ok = parameters["toYear"]
		if !ok {
			return nil, errors.New("Required argument toYear was not found when using GetAlbumList byYear")
		}
	} else if listType == "byGenre" {
		_, ok := parameters["genre"]
		if !ok {
			return nil, errors.New("Required argument genre was not found when using GetAlbumList byGenre")
		}
	}
	params := make(map[string]string)
	params["type"] = listType
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getAlbumList", params)
	if err != nil {
		return nil, err
	}
	return resp.AlbumList.Album, nil
}

// GetAlbumList2 returns a list of albums like GetAlbumList, but organized according to id3 tags.
//
// Optional Parameters:
//   size:           The number of albums to return. Max 500, default 10.
//   offset:         The list offset. Useful if you for example want to page through the list of newest albums.
//   fromYear:       The first year in the range. If fromYear > toYear a reverse chronological list is returned.
//   toYear:         The last year in the range.
//   genre:          The name of the genre, e.g., "Rock".
//   musicFolderId:  (Since 1.11.0) Only return albums in the music folder with the given ID. See getMusicFolders.
//
// toYear and fromYear are required parameters when type == "byYear". genre is a required parameter when type == "byGenre".
func (s *Client) GetAlbumList2(listType string, parameters map[string]string) ([]*AlbumID3, error) {
	if !validateListType(listType) {
		return nil, fmt.Errorf("List type %s is invalid, see http://www.subsonic.org/pages/api.jsp#getAlbumList", listType)
	}
	if listType == "byYear" {
		_, ok := parameters["fromYear"]
		if !ok {
			return nil, errors.New("Required argument fromYear was not found when using GetAlbumList2 byYear")
		}
		_, ok = parameters["toYear"]
		if !ok {
			return nil, errors.New("Required argument toYear was not found when using GetAlbumList2 byYear")
		}
	} else if listType == "byGenre" {
		_, ok := parameters["genre"]
		if !ok {
			return nil, errors.New("Required argument genre was not found when using GetAlbumList2 byGenre")
		}
	}
	params := make(map[string]string)
	params["type"] = listType
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getAlbumList2", params)
	if err != nil {
		return nil, err
	}
	return resp.AlbumList2.Album, nil
}

// GetRandomSongs returns a randomly selected set of songs limited by the optional parameters.
//
// Optional Parameters:
//   size:           The maximum number of songs to return. Max 500, default 10.
//   genre:          Only returns songs belonging to this genre.
//   fromYear:       Only return songs published after or in this year.
//   toYear:         Only return songs published before or in this year.
//   musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *Client) GetRandomSongs(parameters map[string]string) ([]*Child, error) {
	resp, err := s.Get("getRandomSongs", parameters)
	if err != nil {
		return nil, err
	}
	return resp.RandomSongs.Song, nil
}

// GetSongsByGenre returns songs in a given genre name.
//
// Optional Parameters:
//   count:          The maximum number of songs to return. Max 500, default 10.
//   offset:         The offset. Useful if you want to page through the songs in a genre.
//   musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *Client) GetSongsByGenre(name string, parameters map[string]string) ([]*Child, error) {
	params := make(map[string]string)
	params["genre"] = name
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSongsByGenre", params)
	if err != nil {
		return nil, err
	}
	return resp.SongsByGenre.Song, nil
}

// GetNowPlaying returns what is currently being played by all users.
func (s *Client) GetNowPlaying() ([]*NowPlayingEntry, error) {
	resp, err := s.Get("getNowPlaying", nil)
	if err != nil {
		return nil, err
	}
	return resp.NowPlaying.Entry, nil
}

// GetStarred returns starred albums, artists, and songs.
//
// Optional Parameters:
//   musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *Client) GetStarred(parameters map[string]string) (*Starred, error) {
	resp, err := s.Get("getStarred", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Starred, nil
}

// GetStarred2 returns starred albums, artists, and songs arranged by id3 tag.
//
// Optional Parameters:
//   musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *Client) GetStarred2(parameters map[string]string) (*Starred2, error) {
	resp, err := s.Get("getStarred2", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Starred2, nil
}
