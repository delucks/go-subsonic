// Package subsonic implements an API client library for Subsonic-compatible music streaming servers
package subsonic

import (
	"crypto/md5"
	"encoding/json"
	"errors"
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
	libraryVersion      = "0.0.3"
)

type SubsonicClient struct {
	client     *http.Client
	BaseUrl    string
	User       string
	ClientName string
	salt       string
	token      string
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

func (s *SubsonicClient) Authenticate(password string) error {
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
	// Test authentication
	if !s.Ping() {
		return errors.New("Authentication failed")
	}
	return nil
}

func (s *SubsonicClient) Request(method string, endpoint string, params map[string]string) ([]byte, error) {
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
	q.Add("f", "json")
	q.Add("v", supportedApiVersion)
	q.Add("c", s.ClientName)
	q.Add("u", s.User)
	q.Add("t", s.token)
	q.Add("s", s.salt)
	for key, val := range params {
		q.Add(key, val)
	}
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	//log.Printf("%s\n", req.URL)
	//log.Printf("%s\n", contents)
	return contents, nil
}

// Get is a convenience interface to issue a GET request and parse the response body (99% of Subsonic API calls)
func (s *SubsonicClient) Get(endpoint string, params map[string]string) (*SubsonicResponse, error) {
	responseBody, err := s.Request("GET", endpoint, params)
	if err != nil {
		return nil, err
	}
	parsed := APIResponse{}
	err = json.Unmarshal(responseBody, &parsed)
	if err != nil {
		return nil, err
	}
	resp := parsed.Response
	if resp.Error != nil {
		return nil, fmt.Errorf("Error #%d: %s\n", resp.Error.Code, resp.Error.Message)
	}
	return resp, nil
}

func (s *SubsonicClient) Ping() bool {
	_, err := s.Request("GET", "ping", nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *SubsonicClient) GetLicense() (*LicenseValidity, error) {
	resp, err := s.Get("getLicense", nil)
	if err != nil {
		return nil, err
	}
	return resp.License, nil
}
