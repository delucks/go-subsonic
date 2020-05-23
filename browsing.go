package subsonic

// GetMusicFolders returns all configured top-level music folders.
func (s *Client) GetMusicFolders() ([]*MusicFolder, error) {
	resp, err := s.Get("getMusicFolders", nil)
	if err != nil {
		return nil, err
	}
	return resp.MusicFolders.MusicFolder, nil
}

// GetIndexes returns the index of entries by letter/number.
//
// Optional Parameters:
//   musicFolderId:    Only return songs in the music folder with the given ID. See getMusicFolders.
//   ifModifiedSince:  If specified, only return a result if the artist collection has changed since the given time (in milliseconds since 1 Jan 1970).
func (s *Client) GetIndexes(parameters map[string]string) (*Indexes, error) {
	resp, err := s.Get("getIndexes", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Indexes, nil
}

// GetMusicDirectory returns a listing of all files in a music directory. Typically used to get list of albums for an artist, or list of songs for an album.
// The ID can be an album, song, or artist - anything considered within the directory hierarchy of Subsonic.
func (s *Client) GetMusicDirectory(id string) (*Directory, error) {
	resp, err := s.Get("getMusicDirectory", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Directory, nil
}

// GetGenres returns all genres in the server.
func (s *Client) GetGenres() ([]*Genre, error) {
	resp, err := s.Get("getGenres", nil)
	if err != nil {
		return nil, err
	}
	return resp.Genres.Genre, nil
}

// GetArtists returns all artists in the server.
//
// Optional Parameters:
//   musicFolderId:  Only return songs in the music folder with the given ID. See getMusicFolders.
func (s *Client) GetArtists(parameters map[string]string) (*ArtistsID3, error) {
	resp, err := s.Get("getArtists", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Artists, nil
}

// GetAlbum returns an Artist by ID.
func (s *Client) GetArtist(id string) (*ArtistID3, error) {
	resp, err := s.Get("getArtist", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Artist, nil
}

// GetAlbum returns an Album by ID.
func (s *Client) GetAlbum(id string) (*AlbumID3, error) {
	resp, err := s.Get("getAlbum", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Album, nil
}

// GetSong returns a Song by ID.
func (s *Client) GetSong(id string) (*Child, error) {
	resp, err := s.Get("getSong", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Song, nil
}

// GetArtistInfo returns biography, image links, and similar artists from last.fm.
//
// Optional Parameters:
//   count:             Max number of similar artists to return.
//   includeNotPresent: Whether to return artists that are not present in the media library.
func (s *Client) GetArtistInfo(id string, parameters map[string]string) (*ArtistInfo, error) {
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
//
// Optional Parameters:
//   count:             Max number of similar artists to return.
//   includeNotPresent: Whether to return artists that are not present in the media library.
func (s *Client) GetArtistInfo2(id string, parameters map[string]string) (*ArtistInfo2, error) {
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

// GetAlbumInfo returns album notes, image data, etc using data from last.fm.
// This accepts both album and song IDs.
func (s *Client) GetAlbumInfo(id string) (*AlbumInfo, error) {
	resp, err := s.Get("getAlbumInfo", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.AlbumInfo, nil
}

// GetAlbumInfo2 returns the same data as GetAlbumInfo, but organized by id3 tag.
// It only accepts album IDs.
func (s *Client) GetAlbumInfo2(id string) (*AlbumInfo, error) {
	resp, err := s.Get("getAlbumInfo2", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.AlbumInfo, nil
}

// GetSimilarSongs finds similar songs to an album, track, or artist.
// This is mostly used for radio features. This accepts artist, album, or song IDs.
//
// Optional Parameters:
//   count: Number of songs to return
func (s *Client) GetSimilarSongs(id string, parameters map[string]string) ([]*Child, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSimilarSongs", params)
	if err != nil {
		return nil, err
	}
	return resp.SimilarSongs.Song, nil
}

// GetSimilarSongs2 finds similar songs like GetSimilarSongs, but using id3 tags.
//
// Optional Parameters:
//   count: Number of songs to return
func (s *Client) GetSimilarSongs2(id string, parameters map[string]string) ([]*Child, error) {
	params := make(map[string]string)
	params["id"] = id
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getSimilarSongs2", params)
	if err != nil {
		return nil, err
	}
	return resp.SimilarSongs2.Song, nil
}

// GetTopSongs returns the top songs for a given artist by name.
//
// Optional Parameters:
//   count: Number of songs to return
func (s *Client) GetTopSongs(name string, parameters map[string]string) ([]*Child, error) {
	params := make(map[string]string)
	params["artist"] = name
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getTopSongs", params)
	if err != nil {
		return nil, err
	}
	return resp.TopSongs.Song, nil
}
