package subsonic

import "testing"

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
}
