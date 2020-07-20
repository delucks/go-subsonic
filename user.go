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
