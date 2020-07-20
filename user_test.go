package subsonic

import (
	"fmt"
	"testing"
	"time"
)

func runUserTests(client Client, t *testing.T) {
	t.Run("GetUsers", func(t *testing.T) {
		users, err := client.GetUsers()
		if err != nil {
			t.Error(err)
		}
		if len(users) == 0 {
			t.Error("No users returned by GetUsers")
		}
	})

	// The rest of this test is state-heavy: it creates a test user then deletes it after accessing it
	testUserName := fmt.Sprintf("test_user_%v", time.Now().Unix())

	t.Run("CreateUser", func(t *testing.T) {
		beforeUsers, err := client.GetUsers()
		if err != nil {
			t.Error(err)
		}
		// the test user cannot be admin role because they cannot be modified or deleted except by themselves
		err = client.CreateUser(testUserName, "testPassword", "foo@bar.com", map[string]string{
			"jukeboxRole": "true",
		})
		if err != nil {
			t.Error(err)
		}
		afterUsers, err := client.GetUsers()
		if err != nil {
			t.Error(err)
		}
		if len(beforeUsers) == len(afterUsers) {
			t.Errorf("It does not seem that a new user was added between old %#v (len %d) and new %#v (len %d)", beforeUsers, len(beforeUsers), afterUsers, len(afterUsers))
		}
	})

	t.Run("GetUser", func(t *testing.T) {
		user, err := client.GetUser(testUserName)
		if err != nil {
			t.Error(err)
		}
		if user == nil {
			t.Errorf("No user returned for '%s'", testUserName)
		}
	})

	t.Run("UpdateUser", func(t *testing.T) {
		oldUser, err := client.GetUser(testUserName)
		if err != nil {
			t.Error(err)
		}
		err = client.UpdateUser(testUserName, map[string]string{
			"jukeboxRole": "false",
			"password":    "some junk",
		})
		if err != nil {
			t.Error(err)
		}
		newUser, err := client.GetUser(testUserName)
		if err != nil {
			t.Error(err)
		}
		if newUser.JukeboxRole == oldUser.JukeboxRole {
			t.Errorf("Updating user %s's JukeboxRole failed: before %v, after %v", testUserName, oldUser.JukeboxRole, newUser.JukeboxRole)
		}
	})

	t.Run("ChangePassword", func(t *testing.T) {
		err := client.ChangePassword(testUserName, "new password")
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("DeleteUser", func(t *testing.T) {
		beforeUsers, err := client.GetUsers()
		if err != nil {
			t.Error(err)
		}
		err = client.DeleteUser(testUserName)
		if err != nil {
			t.Error(err)
		}
		afterUsers, err := client.GetUsers()
		if err != nil {
			t.Error(err)
		}
		if len(beforeUsers) == len(afterUsers) {
			t.Errorf("It does not seem that user was deleted between old %#v (len %d) and new %#v (len %d)", beforeUsers, len(beforeUsers), afterUsers, len(afterUsers))
		}
	})
}
