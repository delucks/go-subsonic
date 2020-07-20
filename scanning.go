package subsonic

// GetScanStatus returns the current status for media library scanning.
func (c *Client) GetScanStatus() (*ScanStatus, error) {
	resp, err := c.Get("getScanStatus", nil)
	if err != nil {
		return nil, err
	}
	return resp.ScanStatus, nil
}

// StartScan initiates a rescan of the media libraries.
func (c *Client) StartScan() (*ScanStatus, error) {
	resp, err := c.Get("startScan", nil)
	if err != nil {
		return nil, err
	}
	return resp.ScanStatus, nil
}
