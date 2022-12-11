// Package subsonic implements an API client library for Subsonic-compatible music streaming servers.
//
// This project handles communication with a remote *sonic server, but does not handle playback of media. The library user should be prepared to do something with the stream of audio in bytes, like decoding and playing that audio on a sound card.
// The list of API endpoints implemented is available on the project's github page.
//
// The API is divided between functions with no suffix, and functions that have a "2" suffix (or "3" in the case of Search3).
// Generally, things with "2" on the end are organized by file tags rather than folder structure. This is how you'd expect most music players to work and is recommended.
// The variants without a suffix organize the library by directory structure; artists are a directory, albums are children of that directory, songs (subsonic.Child) are children of albums.
// This has some disadvantages: possibly duplicating items with identical directory names, treating songs and albums in much the same fashion, and being more difficult to query consistently.
package subsonic

import (
	"crypto/md5"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"path"
)

const (
	supportedApiVersion = "1.8.0"
	libraryVersion      = "0.0.5"
)

type Client struct {
	Client       *http.Client
	BaseUrl      string
	User         string
	ClientName   string
	PasswordAuth bool
	password     string
	salt         string
	token        string
}

func generateSalt() string {
	var corpus = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	// length is minimum 6, but let's use ten to start
	b := make([]rune, 10)
	for i := range b {
		b[i] = corpus[rand.Intn(len(corpus))]
	}
	return string(b)
}

// Authenticate authenticates the current user with a provided password. The password is salted before transmission and requires Subsonic > 1.13.0.
func (s *Client) Authenticate(password string) error {
	if s.PasswordAuth {
		s.password = password
	} else {
		salt := generateSalt()
		h := md5.New()
		_, err := io.WriteString(h, password)
		if err != nil {
			return err
		}
		_, err = io.WriteString(h, salt)
		if err != nil {
			return err
		}
		s.salt = salt
		s.token = fmt.Sprintf("%x", h.Sum(nil))
	}

	// Test authentication
	// Don't use the s.Ping method because that always returns true as long as the servers is up.
	resp, err := s.Get("ping", nil)
	if err != nil {
		return fmt.Errorf("Authentication failed: %s", err)
	}

	if resp.Error != nil {
		return fmt.Errorf("Authentication failed: %s", resp.Error.Message)
	}

	return nil
}

// Request performs a HTTP request against the Subsonic server as the current user.
func (s *Client) Request(method string, endpoint string, params url.Values) (*http.Response, error) {
	req, err := s.setupRequest(method, endpoint, params)
	if err != nil {
		return nil, err
	}

	resp, err := s.Client.Do(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Get is a convenience interface to issue a GET request and parse the response body (99% of Subsonic API calls)
func (s *Client) Get(endpoint string, params map[string]string) (*Response, error) {
	parameters := url.Values{}
	for k, v := range params {
		parameters.Add(k, v)
	}
	return s.getValues(endpoint, parameters)
}

func (s *Client) setupRequest(method string, endpoint string, params url.Values) (*http.Request, error) {
	baseUrl, err := url.Parse(s.BaseUrl)
	if err != nil {
		return nil, err
	}
	baseUrl.Path = path.Join(baseUrl.Path, "/rest/", endpoint)
	req, err := http.NewRequest(method, baseUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("f", "xml")
	q.Add("v", supportedApiVersion)
	q.Add("c", s.ClientName)
	q.Add("u", s.User)
	if s.PasswordAuth {
		q.Add("p", s.password)
	} else {
		q.Add("t", s.token)
		q.Add("s", s.salt)
	}

	for key, values := range params {
		for _, val := range values {
			q.Add(key, val)
		}
	}
	req.URL.RawQuery = q.Encode()
	//log.Printf("%s %s", method, req.URL.String())
	return req, nil
}

// getValues is a convenience interface to issue a GET request and parse the response body. It supports multiple values by way of the url.Values argument.
func (s *Client) getValues(endpoint string, params url.Values) (*Response, error) {
	response, err := s.Request("GET", endpoint, params)
	if err != nil {
		return nil, err
	}
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	parsed := Response{}
	err = xml.Unmarshal(responseBody, &parsed)
	if err != nil {
		return nil, err
	}
	if parsed.Error != nil {
		return nil, fmt.Errorf("Error #%d: %s\n", parsed.Error.Code, parsed.Error.Message)
	}
	//log.Printf("%s: %s\n", endpoint, string(responseBody))
	return &parsed, nil
}

// Ping is used to test connectivity with the server. It returns true if the server is up.
func (s *Client) Ping() bool {
	_, err := s.Request("GET", "ping", nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

// GetLicense retrieves details about the software license. Subsonic requires a license after a 30-day trial, compatible applications have a perpetually valid license.
func (s *Client) GetLicense() (*License, error) {
	resp, err := s.Get("getLicense", nil)
	if err != nil {
		return nil, err
	}
	return resp.License, nil
}
