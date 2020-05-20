package subsonic

// GetPlaylists returns all playlists a user is allowed to play.
// Optional Parameters:
// * user: get playlists visible to this username rather than the current user. Must have admin permission.
func (s *Client) GetPlaylists(parameters map[string]string) ([]*Playlist, error) {
	resp, err := s.Get("getPlaylists", parameters)
	if err != nil {
		return nil, err
	}
	return resp.Playlists.Entries, nil
}

// GetPlaylist returns a listing of files in a saved playlist.
func (s *Client) GetPlaylist(id string) (*Playlist, error) {
	resp, err := s.Get("getPlaylist", map[string]string{"id": id})
	if err != nil {
		return nil, err
	}
	return resp.Playlist, nil
}
