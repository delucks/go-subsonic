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
}
