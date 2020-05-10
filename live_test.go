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
		license := client.GetLicense()
		if !license.Valid {
			t.Errorf("Invalid license returned- %#v\n", license)
		}
	})
	t.Run("GetMusicFolders", func(t *testing.T) {
		folders := client.GetMusicFolders()
		if len(folders) < 1 {
			t.Error("No music folders were returned from the API")
		}
		for _, f := range folders {
			t.Log(f.Name)
		}
	})
	t.Run("GetIndexes", func(t *testing.T) {
		// Compare no-args usage versus usage with the folder ID
		idx := client.GetIndexes(nil)
		specified := client.GetIndexes(map[string]string{"musicFolderId": "0"})
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
}

func runAirsonicTests(client SubsonicClient, t *testing.T) {
	// Subsonic/Airsonic uses numeric IDs
	t.Run("GetMusicDirectory", func(t *testing.T) {
		// TODO replace this magic number with a song ID when search2 is ready
		dir := client.GetMusicDirectory("5")
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
		dir := client.GetMusicDirectory("6b59470bff90cf113faa72dc01f84995")
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
