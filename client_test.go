package subsonic

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func runClientTests(client Client, t *testing.T) {
	t.Run("Ping", func(t *testing.T) {
		if !client.Ping() {
			t.Error("Ping failed (somehow)")
		}
	})
	t.Run("License", func(t *testing.T) {
		license, err := client.GetLicense()
		if err != nil {
			t.Error(err)
		}
		if !license.Valid {
			t.Errorf("Invalid license returned- %#v\n", license)
		}
	})
}

func validateBaseParams(s *Client, r *http.Request) error {
	wantU := s.User
	gotU := r.URL.Query().Get("u")
	if gotU != wantU {
		return fmt.Errorf("Missing expected parameter 'u'. got='%s' want='%s'", gotU, wantU)
	}

	gotV := r.URL.Query().Get("v")
	if !strings.HasPrefix(gotV, "1.") {
		return fmt.Errorf("Missing expected parameter 'v'. got='%s' want='1.*'", gotV)
	}

	wantC := s.ClientName
	gotC := r.URL.Query().Get("c")
	if gotC != wantC {
		return fmt.Errorf("Missing expected parameter 'c'. got='%s' want='%s'", gotC, wantC)
	}

	return nil
}

func TestPing(t *testing.T) {
	s := &Client{
		Client:     &http.Client{},
		User:       "testUser",
		ClientName: "testClient",
	}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := validateBaseParams(s, r); err != nil {
			t.Errorf("validateBaseParams failed: %s", err)
		}
		switch {
		case strings.HasSuffix(r.URL.Path, "/rest/ping"):
			io.WriteString(w, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.1.1"> </subsonic-response>`)
		default:
			t.Errorf("Unexpected request: %s", r.URL)
		}
	}))
	defer svr.Close()
	s.BaseUrl = svr.URL

	if s.Ping() != true {
		t.Errorf("Incorrect Ping response.")
	}
}

func TestPingDotView(t *testing.T) {
	s := &Client{
		Client:     &http.Client{},
		User:       "testUser",
		ClientName: "testClient",
	}
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if err := validateBaseParams(s, r); err != nil {
			t.Errorf("validateBaseParams failed: %s", err)
		}
		switch {
		case strings.HasSuffix(r.URL.Path, "/rest/ping"):
			w.WriteHeader(404)
		case strings.HasSuffix(r.URL.Path, "/rest/ping.view"):
			io.WriteString(w, `<subsonic-response xmlns="http://subsonic.org/restapi" status="ok" version="1.1.1"> </subsonic-response>`)
		default:
			t.Errorf("Unexpected request: %s", r.URL)
		}
	}))
	defer svr.Close()
	s.BaseUrl = svr.URL

	if s.Ping() != false {
		t.Errorf("Expected Ping false for invalid endpoint.")
	}

	s.RequireDotView = true
	if s.Ping() != true {
		t.Errorf("Incorrect Ping response.")
	}

}
