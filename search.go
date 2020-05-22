package subsonic

// Search2 returns albums, artists and songs matching the given search criteria. Supports paging through the result.
//
// Optional Parameters:
//   artistCount:     Maximum number of artists to return. (Default 20)
//   artistOffset:    Search result offset for artists. Used for paging.
//   albumCount:      Maximum number of albums to return. (Default 20)
//   albumOffset:     Search result offset for albums. Used for paging.
//   songCount:       Maximum number of songs to return. (Default 20)
//   songOffset:      Search result offset for songs. Used for paging.
//   musicFolderId:   (Since 1.12.0) Only return results from the music folder with the given ID. See getMusicFolders.
func (s *Client) Search2(query string, parameters map[string]string) (*SearchResult2, error) {
	params := make(map[string]string)
	params["query"] = query
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("search2", params)
	if err != nil {
		return nil, err
	}
	return resp.SearchResult2, nil
}

// Search3 returns albums, artists and songs matching the given search criteria like Search2, but organized according to id3 tags.
// Optional Parameters:
//   artistCount:     Maximum number of artists to return. (Default 20)
//   artistOffset:    Search result offset for artists. Used for paging.
//   albumCount:      Maximum number of albums to return. (Default 20)
//   albumOffset:     Search result offset for albums. Used for paging.
//   songCount:       Maximum number of songs to return. (Default 20)
//   songOffset:      Search result offset for songs. Used for paging.
//   musicFolderId:   (Since 1.12.0) Only return results from the music folder with the given ID. See getMusicFolders.
func (s *Client) Search3(query string, parameters map[string]string) (*SearchResult3, error) {
	params := make(map[string]string)
	params["query"] = query
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("search3", params)
	if err != nil {
		return nil, err
	}
	return resp.SearchResult3, nil
}
