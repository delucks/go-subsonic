package subsonic

// GetUser gets details about a given user, including which authorization roles and folder access it has. Can be used to enable/disable certain features in the client, such as jukebox control.
func (c *Client) GetUser(username string) (*User, error) {
	resp, err := c.Get("getUser", map[string]string{"username": username})
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}

// GetUsers gets details about all users, including which authorization roles and folder access they have. Only users with admin privileges are allowed to call this method.
func (c *Client) GetUsers() ([]*User, error) {
	resp, err := c.Get("getUsers", nil)
	if err != nil {
		return nil, err
	}
	return resp.Users.User, nil
}

// CreateUser creates a new Subsonic user.
//
// Optional Parameters		Default	Description
//   ldapAuthenticated		false		Whether the user is authenicated in LDAP.
//   adminRole						false		Whether the user is administrator.
//   settingsRole					true		Whether the user is allowed to change personal settings and password.
//   streamRole						true		Whether the user is allowed to play files.
//   jukeboxRole					false		Whether the user is allowed to play files in jukebox mode.
//   downloadRole					false		Whether the user is allowed to download files.
//   uploadRole						false		Whether the user is allowed to upload files.
//   playlistRole					false		Whether the user is allowed to create and delete playlists. Since 1.8.0, changing this role has no effect.
//   coverArtRole					false		Whether the user is allowed to change cover art and tags.
//   commentRole					false		Whether the user is allowed to create and edit comments and ratings.
//   podcastRole					false		Whether the user is allowed to administrate Podcasts.
//   shareRole						false		(Since 1.8.0) Whether the user is allowed to share files with anyone.
//   videoConversionRole	false		(Since 1.15.0) Whether the user is allowed to start video conversions.
//   musicFolderId				All 		(Since 1.12.0) IDs of the music folders the user is allowed access to. Include the parameter once for each folder.
func (c *Client) CreateUser(username, password, email string, parameters map[string]string) error {
	params := make(map[string]string)
	params["username"] = username
	params["password"] = password
	params["email"] = email
	for k, v := range parameters {
		params[k] = v
	}
	_, err := c.Get("createUser", params)
	if err != nil {
		return err
	}
	return nil
}

// UpdateUser modifies an existing Subsonic user.
//
// Optional Parameters:
//   password							The password of the user, either in clear text of hex-encoded.
//   email								The email address of the user.
//   ldapAuthenticated		Whether the user is authenicated in LDAP.
//   adminRole						Whether the user is administrator.
//   settingsRole					Whether the user is allowed to change personal settings and password.
//   streamRole						Whether the user is allowed to play files.
//   jukeboxRole					Whether the user is allowed to play files in jukebox mode.
//   downloadRole					Whether the user is allowed to download files.
//   uploadRole						Whether the user is allowed to upload files.
//   coverArtRole					Whether the user is allowed to change cover art and tags.
//   commentRole					Whether the user is allowed to create and edit comments and ratings.
//   podcastRole					Whether the user is allowed to administrate Podcasts.
//   shareRole						(Since 1.8.0) Whether the user is allowed to share files with anyone.
//   videoConversionRole	(Since 1.15.0) Whether the user is allowed to start video conversions.
//   musicFolderId				(Since 1.12.0) IDs of the music folders the user is allowed access to. Include the parameter once for each folder.
//   maxBitRate						(Since 1.13.0) The maximum bit rate (in Kbps) for the user. Audio streams of higher bit rates are automatically downsampled to this bit rate. Legal values: 0 (no limit), 32, 40, 48, 56, 64, 80, 96, 112, 128, 160, 192, 224, 256, 320.
func (c *Client) UpdateUser(username string, parameters map[string]string) error {
	params := make(map[string]string)
	params["username"] = username
	for k, v := range parameters {
		params[k] = v
	}
	_, err := c.Get("updateUser", params)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUser deletes an existing Subsonic user.
func (s *Client) DeleteUser(username string) error {
	_, err := s.Get("deleteUser", map[string]string{"username": username})
	if err != nil {
		return err
	}
	return nil
}

// ChangePassword changes the password of an existing Subsonic user, using the following parameters. You can only change your own password unless you have admin privileges.
func (c *Client) ChangePassword(username, password string) error {
	_, err := c.Get("changePassword", map[string]string{
		"username": username,
		"password": password,
	})
	if err != nil {
		return err
	}
	return nil
}
