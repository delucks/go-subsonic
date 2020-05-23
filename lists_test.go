package subsonic

import (
	"testing"
	"time"
)

func runListsTests(client Client, t *testing.T) {
	sampleGenre := getSampleGenre(client)

	t.Run("GetAlbumList", func(t *testing.T) {
		_, err := client.GetAlbumList("foobar", nil)
		if err == nil {
			t.Error("No error was returned with an invalid listType argument")
		}
		_, err = client.GetAlbumList("byYear", nil)
		if err == nil {
			t.Error("Failed to validate byYear parameters")
		}
		_, err = client.GetAlbumList("byYear", map[string]string{"fromYear": "1990"})
		if err == nil {
			t.Error("Failed to validate partial byYear parameters")
		}
		_, err = client.GetAlbumList("byGenre", nil)
		if err == nil {
			t.Error("Failed to validate byGenre parameters")
		}
		albums, err := client.GetAlbumList("random", nil)
		if err != nil {
			t.Error(err)
		}
		if albums == nil {
			t.Error("No albums were returned in a call to random getAlbumList")
		}
		for _, album := range albums {
			if album.Title == "" {
				t.Errorf("Album %#v has an empty name :(", album)
			}
		}
		// Work out genre matching
		albums, err = client.GetAlbumList("byGenre", map[string]string{"genre": sampleGenre.Name})
		if err != nil {
			t.Error(err)
		}
		if albums == nil || len(albums) < 1 {
			t.Error("No albums were returned in a call to a byGenre getAlbumList")
		}
		var empty time.Time
		for _, album := range albums {
			if album.Created == empty {
				t.Errorf("Album %#v has empty created time", album)
			}
		}
	})

	t.Run("GetAlbumList2", func(t *testing.T) {
		// Test incorrect parameters
		_, err := client.GetAlbumList2("foobar", nil)
		if err == nil {
			t.Error("No error was returned with an invalid listType argument")
		}
		_, err = client.GetAlbumList2("byYear", nil)
		if err == nil {
			t.Error("Failed to validate byYear parameters")
		}
		_, err = client.GetAlbumList2("byYear", map[string]string{"fromYear": "1990"})
		if err == nil {
			t.Error("Failed to validate partial byYear parameters")
		}
		_, err = client.GetAlbumList2("byGenre", nil)
		if err == nil {
			t.Error("Failed to validate byGenre parameters")
		}
		// Test with proper parameters
		albums, err := client.GetAlbumList2("newest", nil)
		if err != nil {
			t.Error(err)
		}
		if albums == nil {
			t.Error("No albums were returned in a call to newest getAlbumList2")
		}
		var empty time.Time
		for _, album := range albums {
			if album.Name == "" {
				t.Errorf("Album %#v has an empty name", album)
			}
			if album.Created == empty {
				t.Errorf("Album %#v has empty created time", album)
			}
		}
	})

	t.Run("GetRandomSongs", func(t *testing.T) {
		songs, err := client.GetRandomSongs(nil)
		if err != nil || songs == nil {
			t.Error("Basic call to getRandomSongs failed")
		}
		var empty time.Time
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
		songs, err = client.GetRandomSongs(map[string]string{"size": "1"})
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by getRandomSongs failed: expected 1, length actual %d", len(songs))
		}
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
	})

	t.Run("GetSongsByGenre", func(t *testing.T) {
		songs, err := client.GetSongsByGenre(sampleGenre.Name, nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Errorf("No songs returned for genre %v", sampleGenre)
		}
		songs, err = client.GetSongsByGenre(sampleGenre.Name, map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by GetSongsByGenre failed: expected 1, length actual %d", len(songs))
		}
		var empty time.Time
		for _, song := range songs {
			if song.Created == empty {
				t.Errorf("Song %#v had an empty created", song)
			}
		}
	})

	t.Run("GetNowPlaying", func(t *testing.T) {
		// This test is essentially a no-op because we can't depend on the state of playing something in a test environment
		entries, err := client.GetNowPlaying()
		if err != nil {
			t.Error(err)
		}
		var empty time.Time
		for _, nowPlaying := range entries {
			//t.Logf("NowPlaying %d minutes ago, created %v", nowPlaying.MinutesAgo, nowPlaying.Created.Format("2006-01-02T15:04:05.999999-07:00"))
			if nowPlaying.Created == empty {
				t.Errorf("NowPlayingEntry %#v had an empty created", nowPlaying)
			}
		}
	})

	t.Run("GetStarred", func(t *testing.T) {
		// State dependent test
		_, err := client.GetStarred(nil)
		if err != nil {
			t.Error(err)
		}
		_, err = client.GetStarred2(nil)
		if err != nil {
			t.Error(err)
		}
	})
}
