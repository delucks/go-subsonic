package subsonic

import "testing"

func runSearchTests(client Client, t *testing.T) {
	sampleArtist := getSampleArtist(client)

	t.Run("Search2", func(t *testing.T) {
		results, err := client.Search2(sampleArtist.Name, nil)
		if err != nil {
			t.Error(err)
		}
		// The non-id3 matching does not consistently return an artist, but it does erturn that artist's albums
		if len(results.Album) == 0 {
			t.Errorf("Could not find any albums for a known artist %s", sampleArtist.Name)
		}
	})

	t.Run("Search3", func(t *testing.T) {
		results, err := client.Search3(sampleArtist.Name, nil)
		if err != nil {
			t.Error(err)
		}
		if len(results.Artist) == 0 {
			t.Errorf("Could not find a known artist %s", sampleArtist.Name)
		}
		returnOneArtist := map[string]string{
			"artistCount": "1",
			"songCount":   "0",
			"albumCount":  "0",
		}
		results, err = client.Search3(sampleArtist.Name, returnOneArtist)
		if err != nil {
			t.Error(err)
		}
		if len(results.Artist) != 1 || len(results.Song) != 0 || len(results.Album) != 0 {
			t.Errorf("Improperly limited results of search for %s: %#v", sampleArtist.Name, results)
		}
	})
}
