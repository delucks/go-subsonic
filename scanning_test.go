package subsonic

import "testing"

func runScanningTests(client Client, t *testing.T) {
	t.Run("GetScanStatus", func(t *testing.T) {
		_, err := client.GetScanStatus()
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("StartScan", func(t *testing.T) {
		status, err := client.StartScan()
		if err != nil {
			t.Error(err)
		}
		if !status.Scanning {
			t.Error("Scan doesn't seem to have started after call to startScan")
		}
	})
}
