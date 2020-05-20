package subsonic

import "errors"

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

// CreatePlaylist creates (or updates) a playlist.
// Optional Parameters:
// * songId     ID of a song in the playlist. Use one songId parameter for each song in the playlist.
// Mutually Exclusive Parameters:
// * playlistId The playlist ID.
// * name       The human-readable name of the playlist.
// This returns a Playlist object in Subsonic > 1.14.0, so it cannot consistently return a *Playlist
func (s *Client) CreatePlaylist(parameters map[string]string) error {
	_, idPresent := parameters["playlistId"]
	_, namePresent := parameters["name"]
	if !(idPresent || namePresent) {
		return errors.New("One of name or playlistId is mandatory, to create or update a playlist respectively")
	}
	_, err := s.Get("createPlaylist", parameters)
	if err != nil {
		return err
	}
	return nil
}
