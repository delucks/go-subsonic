package subsonic

import (
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"testing"
)

func getSampleGenre(client Client) *Genre {
	genres, err := client.GetGenres()
	if err != nil {
		return nil
	}
	for _, selection := range genres {
		if strings.Contains(selection.Name, "Empty") {
			// the Empty token should be skipped
			continue
		}
		if selection.SongCount > 0 && selection.AlbumCount > 0 {
			return selection
		}
	}
	return nil
}

func getSampleArtist(client Client) *ArtistID3 {
	artists, err := client.GetArtists(nil)
	if err != nil {
		return nil
	}
	return artists.Index[len(artists.Index)-1].Artist[0]
}

func getSampleSong(client Client) *Child {
	songs, err := client.GetRandomSongs(nil)
	if err != nil {
		return nil
	}
	return songs[rand.Intn(len(songs))]
}

func getSampleAlbum(client Client) *AlbumID3 {
	albums, err := client.GetAlbumList2("newest", nil)
	if err != nil {
		return nil
	}
	return albums[rand.Intn(len(albums))]
}

func findPlaylistByName(client Client, name string) (*Playlist, error) {
	playlists, err := client.GetPlaylists(nil)
	if err != nil {
		return nil, err
	}
	for _, p := range playlists {
		if p.Name == name {
			return p, nil
		}
	}
	return nil, fmt.Errorf("Could not find playlist %s", name)
}

func getSampleArtistFolder(client Client) string {
	// this uses the folder-based navigation structure to find one valid child in the hierarchy
	indexes, err := client.GetIndexes(nil)
	if err != nil {
		return ""
	}
	for _, index := range indexes.Index {
		if len(index.Artist) > 0 {
			return index.Artist[0].ID
		}
	}
	return ""
}

func runAirsonicTests(client Client, t *testing.T) {
	// These are tests for functionality not implemented in Navidrome yet
	// Compatibility list: https://www.navidrome.org/docs/developers/subsonic-api/
	sampleArtist := getSampleArtist(client)
	sampleAlbum := getSampleAlbum(client)

	// Browsing
	t.Run("GetSimilarSongs", func(t *testing.T) {
		sampleFolder := getSampleArtistFolder(client)
		_, err := client.GetSimilarSongs(sampleFolder, nil)
		if err != nil {
			t.Error(err)
		}
		songs, err := client.GetSimilarSongs2(sampleFolder, nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Errorf("GetSimilarSongs2 returned nil recommendations for artist %#v!", sampleFolder)
		}
		// Make sure the count argument is getting properly passed
		songs, err = client.GetSimilarSongs2(sampleFolder, map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Count argument did not work properly: got %d songs in response to a request for one", len(songs))
		}
	})

	t.Run("GetAlbumInfo", func(t *testing.T) {
		ai, err := client.GetAlbumInfo(sampleAlbum.ID)
		if err != nil {
			t.Error(err)
		}
		empty := AlbumInfo{}
		if *ai == empty {
			// This can't fail the test even though it's anomalous because many albums
			// do not have information available for them in the databases
			t.Logf("Empty AlbumInfo returned for album %s by %s", sampleAlbum.Name, sampleAlbum.Artist)
		}
		ai, err = client.GetAlbumInfo2(sampleAlbum.ID)
		if err != nil {
			t.Error(err)
		}
		if *ai == empty {
			t.Logf("Empty AlbumInfo2 returned for album %s by %s", sampleAlbum.Name, sampleAlbum.Artist)
		}
	})

	t.Run("GetTopSongs", func(t *testing.T) {
		songs, err := client.GetTopSongs(sampleArtist.Name, nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Errorf("No top songs returned for known artist %s", sampleArtist.Name)
		}
		songs, err = client.GetTopSongs(sampleArtist.Name, map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Incorrect song count returned from call to getTopSongs, %d actual 1 expected", len(songs))
		}
	})
}

func TestNavidrome(t *testing.T) {
	client := Client{
		Client:     &http.Client{},
		BaseUrl:    "http://127.0.0.1:4533/",
		User:       "admin",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("admin")
	if err != nil {
		t.Error(err)
	}
	runClientTests(client, t)
	runBrowsingTests(client, t)
	runListsTests(client, t)
	runPlaylistTests(client, t)
	runRetrievalTests(client, t)
	runSearchTests(client, t)
	runAnnotationTests(client, t)
}

func TestAirsonic(t *testing.T) {
	client := Client{
		Client:     &http.Client{},
		BaseUrl:    "http://127.0.0.1:4040/",
		User:       "admin",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("admin")
	if err != nil {
		t.Error(err)
	}
	runClientTests(client, t)
	runBrowsingTests(client, t)
	runListsTests(client, t)
	runPlaylistTests(client, t)
	runRetrievalTests(client, t)
	runSearchTests(client, t)
	runAnnotationTests(client, t)

	runAirsonicTests(client, t)
	runUserTests(client, t)
	// Initiating a scan interferes with other tests in the suite by removing indexes temporarily, so we run it last
	runScanningTests(client, t)
}

func TestSubsonic(t *testing.T) {
	client := Client{
		Client:     &http.Client{},
		BaseUrl:    "http://demo.subsonic.org/",
		User:       "guest5",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("guest")
	if err != nil {
		t.Error(err)
	}
	runClientTests(client, t)
	runBrowsingTests(client, t)
	runListsTests(client, t)
	runRetrievalTests(client, t)
	runSearchTests(client, t)
	runAnnotationTests(client, t)

	runAirsonicTests(client, t)
	runUserTests(client, t)
	runScanningTests(client, t)
}
