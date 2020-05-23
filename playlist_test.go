package subsonic

import (
	"fmt"
	"testing"
	"time"
)

func runPlaylistTests(client Client, t *testing.T) {
	t.Run("GetPlaylists", func(t *testing.T) {
		playlists, err := client.GetPlaylists(nil)
		if err != nil {
			t.Error(err)
		}
		var empty time.Time
		for _, p := range playlists {
			if p.ID == "" {
				t.Errorf("Invalid playlist returned %#v", p)
			}
			if p.Created == empty {
				t.Errorf("Playlist %#v had an empty created", p)
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
		var empty time.Time
		if playlist.Created == empty {
			t.Errorf("Playlist %#v had an empty created", playlist)
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
