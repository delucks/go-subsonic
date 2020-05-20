package subsonic

import (
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func getRandomGenre(client Client) *Genre {
	rand.Seed(time.Now().Unix())
	genres, err := client.GetGenres()
	if err != nil {
		return nil
	}
	// make sure the genre has songs present
	selection := genres[rand.Intn(len(genres))]
	for {
		if selection.SongCount > 0 {
			break
		}
		selection = genres[rand.Intn(len(genres))]
	}
	return selection
}

func getSampleArtist(client Client) *Artist {
	artists, err := client.GetArtists(nil)
	if err != nil {
		return nil
	}
	return artists.Indexes[len(artists.Indexes)-1].Artists[0]
}

func getSamplePlaylist(client Client) *Playlist {
	playlists, err := client.GetPlaylists(nil)
	if err != nil {
		return nil
	}
	return playlists[rand.Intn(len(playlists))]
}

func runCommonTests(client Client, t *testing.T) {
	sampleArtist := getSampleArtist(client)
	sampleGenre := getRandomGenre(client)
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
		albums, err = client.GetAlbumList("byGenre", map[string]string{"genre": sampleGenre.Value})
		if err != nil {
			t.Error(err)
		}
		if albums == nil || len(albums) < 1 {
			t.Error("No albums were returned in a call to a byGenre getAlbumList")
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
		for _, album := range albums {
			if album.Name == "" {
				t.Errorf("Album %#v has an empty name :(", album)
			}
		}
	})
	t.Run("GetRandomSongs", func(t *testing.T) {
		songs, err := client.GetRandomSongs(nil)
		if err != nil || songs == nil {
			t.Error("Basic call to getRandomSongs failed")
		}
		songs, err = client.GetRandomSongs(map[string]string{"size": "1"})
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by getRandomSongs failed: expected 1, length actual %d", len(songs))
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
	t.Run("GetSongsByGenre", func(t *testing.T) {
		songs, err := client.GetSongsByGenre(sampleGenre.Value, nil)
		if err != nil {
			t.Error(err)
		}
		if songs == nil {
			t.Errorf("No songs returned for genre %v", sampleGenre)
		}
		songs, err = client.GetSongsByGenre(sampleGenre.Value, map[string]string{"count": "1"})
		if err != nil {
			t.Error(err)
		}
		if len(songs) != 1 {
			t.Errorf("Limiting songs returned by GetSongsByGenre failed: expected 1, length actual %d", len(songs))
		}
	})
	t.Run("GetNowPlaying", func(t *testing.T) {
		// This test is essentially a no-op because we can't depend on the state of playing something in a test environment
		entries, err := client.GetNowPlaying()
		if err != nil {
			t.Error(err)
		}
		for _, nowPlaying := range entries {
			t.Logf("%#v", nowPlaying)
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
	t.Run("Search2", func(t *testing.T) {
		results, err := client.Search2(sampleArtist.Name, nil)
		if err != nil {
			t.Error(err)
		}
		// The non-id3 matching does not consistently return an artist, but it does erturn that artist's albums
		if len(results.Albums) == 0 {
			t.Errorf("Could not find any albums for a known artist %s", sampleArtist.Name)
		}
	})
	t.Run("Search3", func(t *testing.T) {
		results, err := client.Search3(sampleArtist.Name, nil)
		if err != nil {
			t.Error(err)
		}
		if len(results.Artists) == 0 {
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
		if len(results.Artists) != 1 || len(results.Songs) != 0 || len(results.Albums) != 0 {
			t.Errorf("Improperly limited results of search for %s: %#v", sampleArtist.Name, results)
		}
	})
	t.Run("Stream", func(t *testing.T) {
		// Purposefully choose an ID that returns an error
		_, err := client.Stream("1", nil)
		if err == nil {
			t.Error("An error was not returned on ID 1")
		}
		contents, err := client.Stream("33", nil)
		if err != nil {
			t.Error(err)
		}
		if contents == nil {
			t.Error("No content returned")
		}
	})
	t.Run("GetPlaylists", func(t *testing.T) {
		playlists, err := client.GetPlaylists(nil)
		if err != nil {
			t.Error(err)
		}
		for _, p := range playlists {
			if p.ID == "" {
				t.Errorf("Invalid playlist returned %#v", p)
			}
		}
	})
	t.Run("GetPlaylist", func(t *testing.T) {
		sample := getSamplePlaylist(client)
		if sample == nil {
			t.Error("Failed to get sample playlist")
		}
		playlist, err := client.GetPlaylist(sample.ID)
		if err != nil {
			t.Error(err)
		}
		if playlist.ID == "" {
			t.Errorf("Invalid playlist returned %#v", playlist)
		}
	})
}

func runAirsonicTests(client Client, t *testing.T) {
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
	client := Client{
		Client:     &http.Client{},
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
	client := Client{
		Client:     &http.Client{},
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
	runCommonTests(client, t)
	runAirsonicTests(client, t)
}
