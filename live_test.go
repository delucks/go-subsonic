package main

import (
	"net/http"
	"testing"
)

func runCommonTests(client SubsonicClient, t *testing.T) {
	// These test the library's ability to unmarshal server responses
	t.Run("Ping", func(t *testing.T) {
		if !client.Ping() {
			t.Error("Ping failed (somehow)")
		}
	})
	t.Run("License", func(t *testing.T) {
		license, err := client.GetLicense()
		if err != nil {
			t.Error(err)
		}
		if !license.Valid {
			t.Errorf("Invalid license returned- %#v\n", license)
		}
	})
	t.Run("GetMusicFolders", func(t *testing.T) {
		folders, err := client.GetMusicFolders()
		if err != nil {
			t.Error(err)
		}
		if len(folders) < 1 {
			t.Error("No music folders were returned from the API")
		}
		for _, f := range folders {
			t.Log(f.Name)
		}
	})
	t.Run("GetIndexes", func(t *testing.T) {
		// Compare no-args usage versus usage with the folder ID
		idx, err := client.GetIndexes(nil)
		if err != nil {
			t.Error(err)
		}
		specified, err := client.GetIndexes(map[string]string{"musicFolderId": "0"})
		if err != nil {
			t.Error(err)
		}
		if idx.LastModified != specified.LastModified {
			t.Errorf("LastModified differs: %v -> %v (specified)", idx.LastModified, specified.LastModified)
		}
		if idx.IgnoredArticles != specified.IgnoredArticles {
			t.Errorf("IgnoredArticles differs: %s -> %s (specified)", idx.IgnoredArticles, specified.IgnoredArticles)
		}
		for position, index := range idx.Indexes {
			if index.Name != specified.Indexes[position].Name {
				t.Errorf("Names differ: %s -> %s (specified)", index.Name, specified.Indexes[position].Name)
			}
		}
	})
	t.Run("GetGenres", func(t *testing.T) {
		genres, err := client.GetGenres()
		if err != nil {
			t.Error(err)
		}
		for _, g := range genres {
			if g.Value == "" {
				t.Error("Empty genre present")
			}
			if g.SongCount == 0 {
				t.Errorf("Genre %s has no songs", g.Value)
			}
		}
	})
	t.Run("GetArtists", func(t *testing.T) {
		idx, err := client.GetArtists(nil)
		if err != nil {
			t.Error(err)
		}
		specified, err := client.GetArtists(map[string]string{"musicFolderId": "0"})
		if err != nil {
			t.Error(err)
		}
		if idx.IgnoredArticles != specified.IgnoredArticles {
			t.Errorf("IgnoredArticles differs: %s -> %s (specified)", idx.IgnoredArticles, specified.IgnoredArticles)
		}
		for position, index := range idx.Indexes {
			if index.Name != specified.Indexes[position].Name {
				t.Errorf("Names differ: %s -> %s (specified)", index.Name, specified.Indexes[position].Name)
			}
		}
	})
}

func runAirsonicTests(client SubsonicClient, t *testing.T) {
	// Subsonic/Airsonic uses numeric IDs
	t.Run("GetMusicDirectory", func(t *testing.T) {
		// TODO replace this magic number with a song ID when search2 is ready
		dir, err := client.GetMusicDirectory("5")
		if err != nil {
			t.Error(err)
		}
		if dir.ID == "" {
			t.Error("Directory has an empty ID")
		}
		if dir.Name == "" {
			t.Error("Directory has an empty Name")
		}
		for _, child := range dir.Children {
			if child.ID == "" {
				t.Log(child.Title)
				t.Errorf("Child %s has an empty ID", child.Title)
			}
		}
	})
	t.Run("GetArtist", func(t *testing.T) {
		artist, err := client.GetArtist("1") // the subsonic demo server does not have an artist 0
		if err != nil {
			t.Error(err)
		}
		if len(artist.Albums) != artist.AlbumCount {
			t.Errorf("Artist %s has %d albums in the 'album' key, but an AlbumCount of %d", artist.Name, len(artist.Albums), artist.AlbumCount)
		}
	})
	t.Run("GetAlbum", func(t *testing.T) {
		album, err := client.GetAlbum("1")
		if err != nil {
			t.Error(err)
		}
		if len(album.Songs) != album.SongCount {
			t.Errorf("Album %s has %d songs in the 'song' key, but an songCount of %d", album.Name, len(album.Songs), album.SongCount)
		}
	})
	t.Run("GetSong", func(t *testing.T) {
		song, err := client.GetSong("27")
		if err != nil {
			t.Error(err)
		}
		if song.ID == "" {
			t.Errorf("Song was not returned properly, %#v\n", song)
		}
	})
	t.Run("GetArtistInfo", func(t *testing.T) {
		ai, err := client.GetArtistInfo("3", nil)
		if err != nil {
			t.Error(err)
		}
		if ai.Biography == "" {
			t.Error("Empty biography, invalid response")
		}
		ai, err = client.GetArtistInfo2("1", nil)
		if err != nil {
			t.Error(err)
		}
		if ai.Biography == "" {
			t.Error("Empty biography, invalid response")
		}
	})
	t.Run("GetAlbumInfo", func(t *testing.T) {
		ai, err := client.GetAlbumInfo("48")
		if err != nil {
			t.Error(err)
		}
		if ai.MusicBrainzID == "" {
			t.Logf("%#v\n", ai)
			t.Error("Empty MB id from GetAlbumInfo, invalid response")
		}
		ai, err = client.GetAlbumInfo2("1")
		if err != nil {
			t.Error(err)
		}
		if ai.MusicBrainzID == "" {
			t.Logf("%#v\n", ai)
			t.Error("Empty MB id from GetAlbumInfo2, invalid response")
		}
	})
	t.Run("GetSimilarSongs", func(t *testing.T) {
		_, err := client.GetSimilarSongs("48", nil)
		if err != nil {
			t.Error(err)
		}
		// Cannot check for song contents here because GetSimilarSongs on ID 48 may or may not return data
		songs, err := client.GetSimilarSongs2("1", nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Error("GetSimilarSongs2 returned nil recommendations for ID 1!")
		}
		// Make sure the count argument is getting properly passed
		songs, err = client.GetSimilarSongs2("1", map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Count argument did not work properly: got %d songs in response to a request for one", len(songs))
		}
	})
}

/*
func TestNavidrome(t *testing.T) {
	client := SubsonicClient{
		client:     &http.Client{},
		BaseUrl:    "http://192.168.1.7:4040/",
		User:       "test",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("blah")
	if err != nil {
		t.Error(err)
	}
	runCommonTests(client, t)
	// Navidrome uses UUIDs (strings)
	t.Run("GetMusicDirectory", func(t *testing.T) {
		// TODO replace this magic uuid with a real one when search2 is ready
		dir, err := client.GetMusicDirectory("6b59470bff90cf113faa72dc01f84995")
		if err != nil {
			t.Error(err)
		}
		if dir.ID == "" {
			t.Error("Directory has an empty ID")
		}
		if dir.Name == "" {
			t.Error("Directory has an empty Name")
		}
		for _, child := range dir.Children {
			t.Log(child.Title)
			if child.ID == "" {
				t.Errorf("Child %s has an empty ID", child.Title)
			}
		}
	})
}
*/

func TestAirsonic(t *testing.T) {
	client := SubsonicClient{
		client:     &http.Client{},
		BaseUrl:    "http://127.0.0.1:8080/",
		User:       "admin",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("admin")
	if err != nil {
		t.Error(err)
	}
	runCommonTests(client, t)
	runAirsonicTests(client, t)
}

func TestSubsonic(t *testing.T) {
	client := SubsonicClient{
		client:     &http.Client{},
		BaseUrl:    "http://demo.subsonic.org/",
		User:       "guest5",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("guest")
	if err != nil {
		t.Error(err)
	}
	runCommonTests(client, t)
	runAirsonicTests(client, t)
}
