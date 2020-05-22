package subsonic

import (
	"fmt"
	"math/rand"
	"net/http"
	"testing"
	"time"
)

func getSampleGenre(client Client) *Genre {
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

func runClientTests(client Client, t *testing.T) {
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
}

func runBrowsingTests(client Client, t *testing.T) {
	sampleArtist := getSampleArtist(client)
	sampleSong := getSampleSong(client)
	sampleAlbum := getSampleAlbum(client)

	// These test the library's ability to unmarshal server responses
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
		for position, index := range idx.Index {
			if index.Name != specified.Index[position].Name {
				t.Errorf("Names differ: %s -> %s (specified)", index.Name, specified.Index[position].Name)
			}
		}
	})

	t.Run("GetMusicDirectory", func(t *testing.T) {
		dir, err := client.GetMusicDirectory(sampleArtist.ID)
		if err != nil {
			t.Error(err)
		}
		if dir.ID == "" {
			t.Error("Directory has an empty ID")
		}
		if dir.Name == "" {
			t.Error("Directory has an empty Name")
		}
		for _, child := range dir.Child {
			if child.ID == "" {
				t.Log(child.Title)
				t.Errorf("Child %s has an empty ID", child.Title)
			}
		}
	})

	t.Run("GetGenres", func(t *testing.T) {
		genres, err := client.GetGenres()
		if err != nil {
			t.Error(err)
		}
		for _, g := range genres {
			if g.Name == "" {
				t.Error("Empty genre present")
			}
			if g.SongCount == 0 {
				t.Errorf("Genre %s has no songs", g.Name)
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
		for position, index := range idx.Index {
			if index.Name != specified.Index[position].Name {
				t.Errorf("Names differ: %s -> %s (specified)", index.Name, specified.Index[position].Name)
			}
		}
	})

	t.Run("GetArtist", func(t *testing.T) {
		artist, err := client.GetArtist(sampleArtist.ID)
		if err != nil {
			t.Error(err)
		}
		if len(artist.Album) != artist.AlbumCount {
			t.Errorf("Artist %s has %d albums in the 'album' key, but an AlbumCount of %d", artist.Name, len(artist.Album), artist.AlbumCount)
		}
	})

	t.Run("GetAlbum", func(t *testing.T) {
		album, err := client.GetAlbum(sampleAlbum.ID)
		if err != nil {
			t.Error(err)
		}
		if len(album.Song) != album.SongCount {
			t.Errorf("Album %s has %d songs in the 'song' key, but an songCount of %d", album.Name, len(album.Song), album.SongCount)
		}
	})

	t.Run("GetSong", func(t *testing.T) {
		song, err := client.GetSong(sampleSong.ID)
		if err != nil {
			t.Error(err)
		}
		if song.ID == "" {
			t.Errorf("Song was not returned properly, %#v\n", song)
		}
	})

	t.Run("GetArtistInfo", func(t *testing.T) {
		ai, err := client.GetArtistInfo(sampleArtist.ID, nil)
		if err != nil {
			t.Error(err)
		}
		if ai.Biography == "" {
			t.Logf("Empty ArtistInfo returned for artist %s", sampleArtist.Name)
		}
		ai2, err := client.GetArtistInfo2(sampleArtist.ID, nil)
		if err != nil {
			t.Error(err)
		}
		if ai2.Biography == "" {
			t.Logf("Empty ArtistInfo2 returned for artist %s", sampleArtist.Name)
		}
	})
}

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
}

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

func runRetrievalTests(client Client, t *testing.T) {
	sampleSong := getSampleSong(client)

	t.Run("Stream", func(t *testing.T) {
		// Purposefully choose an ID that returns an error
		_, err := client.Stream("1", nil)
		if err == nil {
			t.Error("An error was not returned on ID 1")
		}
		contents, err := client.Stream(sampleSong.ID, nil)
		if err != nil {
			t.Error(err)
		}
		if contents == nil {
			t.Error("No content returned")
		}
	})
}

func runPlaylistTests(client Client, t *testing.T) {
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
	// State-heavy test for playlist CRUD
	testPlaylistName := fmt.Sprintf("Test playlist %v", time.Now().Unix())
	t.Run("CreatePlaylist", func(t *testing.T) {
		err := client.CreatePlaylist(map[string]string{
			"name": testPlaylistName,
		})
		if err != nil {
			t.Error(err)
		}
	})
	testPlaylist, err := findPlaylistByName(client, testPlaylistName)
	if err != nil {
		t.Log(err)
	}
	t.Run("GetPlaylist", func(t *testing.T) {
		playlist, err := client.GetPlaylist(testPlaylist.ID)
		if err != nil {
			t.Error(err)
		}
		if playlist.ID == "" || playlist.ID != testPlaylist.ID {
			t.Errorf("Invalid playlist returned %#v", playlist)
		}
	})
	t.Run("UpdatePlaylist", func(t *testing.T) {
		err = client.UpdatePlaylist(testPlaylist.ID, map[string]string{
			"comment": "Whee there buddy",
		})
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("DeletePlaylist", func(t *testing.T) {
		err = client.DeletePlaylist(testPlaylist.ID)
		if err != nil {
			t.Error(err)
		}
	})
}

func runAirsonicTests(client Client, t *testing.T) {
	// These are not implemented in Navidrome yet
	sampleArtist := getSampleArtist(client)
	sampleAlbum := getSampleAlbum(client)

	// Browsing
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
		BaseUrl:    "http://192.168.1.7:4040/",
		User:       "test",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("blah")
	if err != nil {
		t.Error(err)
	}
	runClientTests(client, t)
	runBrowsingTests(client, t)
	runListsTests(client, t)
	runPlaylistTests(client, t)
	runRetrievalTests(client, t)
	runSearchTests(client, t)
}

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
	runClientTests(client, t)
	runBrowsingTests(client, t)
	runListsTests(client, t)
	runPlaylistTests(client, t)
	runRetrievalTests(client, t)
	runSearchTests(client, t)
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
	runClientTests(client, t)
	runBrowsingTests(client, t)
	runListsTests(client, t)
	runRetrievalTests(client, t)
	runSearchTests(client, t)
	runAirsonicTests(client, t)
}
