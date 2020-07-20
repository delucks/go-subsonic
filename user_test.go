package subsonic

import "testing"

func runUserTests(client Client, t *testing.T) {
	t.Run("GetUser", func(t *testing.T) {
		user, err := client.GetUser("admin")
		if err != nil {
			t.Error(err)
		}
		if user == nil {
			t.Error("No user returned for 'admin'")
		}
	})
}
