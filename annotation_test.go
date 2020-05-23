package subsonic

import (
	"fmt"
	"testing"
	"time"
)

func runAnnotationTests(client Client, t *testing.T) {
	sampleArtist := getSampleArtist(client)
	sampleSong := getSampleSong(client)
	sampleAlbum := getSampleAlbum(client)
	itemsToStar := StarParameters{
		SongIDs:   []string{sampleSong.ID},
		AlbumIDs:  []string{sampleAlbum.ID},
		ArtistIDs: []string{sampleArtist.ID},
	}

	t.Run("Star", func(t *testing.T) {
		err := client.Star(itemsToStar)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Unstar", func(t *testing.T) {
		err := client.Unstar(itemsToStar)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("SetRating", func(t *testing.T) {
		err := client.SetRating(sampleSong.ID, -1)
		if err == nil {
			t.Error("Bounds checking on ratings failed")
		}
		err = client.SetRating(sampleSong.ID, 10)
		if err == nil {
			t.Error("Bounds checking on ratings failed")
		}
		err = client.SetRating(sampleSong.ID, 4)
		if err != nil {
			t.Error(err)
		}
		err = client.SetRating(sampleSong.ID, 0)
		if err != nil {
			t.Error(err)
		}
		err = client.SetRating(sampleArtist.ID, 5)
		if err != nil {
			t.Error(err)
		}
		err = client.SetRating(sampleArtist.ID, 0)
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Scrobble", func(t *testing.T) {
		err := client.Scrobble(sampleSong.ID, map[string]string{
			"time": "not-an-int",
		})
		if err == nil {
			t.Error("Typecheck on time argument failed")
		}
		err = client.Scrobble(sampleSong.ID, map[string]string{
			"submission": "not-a-bool",
		})
		if err == nil {
			t.Error("Typecheck on submission argument failed")
		}
		err = client.Scrobble(sampleSong.ID, map[string]string{
			"time":       fmt.Sprintf("%d", time.Now().Unix()),
			"submission": "true",
		})
		if err != nil {
			t.Error(err)
		}
		err = client.Scrobble(sampleSong.ID, nil)
		if err != nil {
			t.Error(err)
		}
	})
}
