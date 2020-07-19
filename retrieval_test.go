package subsonic

import (
	"image"
	"testing"
)

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
		sample := make([]byte, 8)
		_, err = contents.Read(sample)
		if err != nil {
			t.Errorf("Error reading sample %v: %v", sample, err)
		}
	})

	t.Run("Download", func(t *testing.T) {
		contents, err := client.Download(sampleSong.ID)
		if err != nil {
			t.Fatal(err)
		}
		if contents == nil {
			t.Fatal("No content returned")
		}
		sample := make([]byte, 8)
		_, err = contents.Read(sample)
		if err != nil {
			t.Errorf("Error reading sample %v: %v", sample, err)
		}
	})

	t.Run("GetCoverArt", func(t *testing.T) {
		img, err := client.GetCoverArt(sampleSong.ID, nil)
		if err != nil {
			t.Fatal(err)
		}
		var empty image.Rectangle
		if img.Bounds() == empty {
			t.Fatalf("Image %#v has empty bounds", img)
		}
	})

	t.Run("GetAvatar", func(t *testing.T) {
		// Avatars aren't always returned by servers, so this test is a no-op
		_, err := client.GetAvatar(client.User)
		if err != nil {
			t.Log(err)
		}
	})
}
