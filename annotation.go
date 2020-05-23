package subsonic

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
)

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

// SetRating sets the rating of a music file.
func (s *Client) SetRating(id string, rating int) error {
	if rating > 5 || rating < 0 {
		return errors.New("Rating can only be in the range 0-5")
	}
	params := map[string]string{
		"id":     id,
		"rating": fmt.Sprintf("%d", rating),
	}
	_, err := s.Get("setRating", params)
	if err != nil {
		return err
	}
	return nil
}

// Scrobble submits a song to last.fm if the user has configured their credentials to do so.
//
// Optional Parameters:
//   time:            (Since 1.8.0) The time (in milliseconds since 1 Jan 1970) at which the song was listened to.
//   submission:      Whether this is a "submission" (true) or a "now playing" (false) notification. Defaults to a submission.
func (s *Client) Scrobble(id string, parameters map[string]string) error {
	params := map[string]string{
		"id": id,
	}
	if scrobbleTime, ok := parameters["time"]; ok {
		_, err := strconv.Atoi(scrobbleTime)
		if err != nil {
			return fmt.Errorf("%s is not a unix-style timestamp", scrobbleTime)
		}
		params["time"] = scrobbleTime
	}
	if submission, ok := parameters["submission"]; ok {
		_, err := strconv.ParseBool(submission)
		if err != nil {
			return fmt.Errorf("%s is not boolean", submission)
		}
		params["submission"] = submission
	}
	_, err := s.Get("scrobble", params)
	if err != nil {
		return err
	}
	return nil
}
