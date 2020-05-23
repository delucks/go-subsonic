package subsonic

import "net/url"

// StarParameters are used to identify songs, albums, and artists (or some subset of those) at the same time.
// subsonic.Star and subsonic.Unstar both use StarParameters to identify things to star.
type StarParameters struct {
	SongIDs   []string
	AlbumIDs  []string
	ArtistIDs []string
}

// Star adds the star annotation to songs, albums, and artists.
func (s *Client) Star(parameters StarParameters) error {
	params := url.Values{}
	for _, song := range parameters.SongIDs {
		params.Add("id", song)
	}
	for _, album := range parameters.AlbumIDs {
		params.Add("albumId", album)
	}
	for _, artist := range parameters.ArtistIDs {
		params.Add("artistId", artist)
	}
	_, err := s.getValues("star", params)
	if err != nil {
		return err
	}
	return nil
}

// Unstar removes the star annotation from songs, albums, and artists.
func (s *Client) Unstar(parameters StarParameters) error {
	params := url.Values{}
	for _, song := range parameters.SongIDs {
		params.Add("id", song)
	}
	for _, album := range parameters.AlbumIDs {
		params.Add("albumId", album)
	}
	for _, artist := range parameters.ArtistIDs {
		params.Add("artistId", artist)
	}
	_, err := s.getValues("unstar", params)
	if err != nil {
		return err
	}
	return nil
}
