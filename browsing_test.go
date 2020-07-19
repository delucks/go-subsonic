package subsonic

import (
	"testing"
	"time"
)

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
		// This test is potentially affected by race condition between the first and second
		// GetIndexes call. It's clear that LastModified can differ at any time, but if this
		// is running against an actively scanning collection with more music to be added,
		// the index content can differ as well.
		// tl;dr if this test fails when you're running it against an active music collection and not
		// the test collection downloaded by test.sh, don't worry about it. If you have suggestions
		// for how to change this test please make an issue/PR!
		//
		// Compare no-args usage versus usage with the folder ID
		idx, err := client.GetIndexes(nil)
		if err != nil {
			t.Error(err)
		}
		specified, err := client.GetIndexes(map[string]string{"musicFolderId": "0"})
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
		var empty time.Time
		if album.Created == empty {
			t.Errorf("Album %#v has an empty created date", album)
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
		var empty time.Time
		if song.Created == empty {
			t.Errorf("Song %#v has an empty created date", song)
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
