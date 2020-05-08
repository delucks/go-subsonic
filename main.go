package main

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
	libraryVersion      = "0.0.1"
)

type LicenseValidity struct {
	Valid bool `json:"valid"`
}

type SubsonicResponse struct {
	Status        string `json:"status"`
	Version       string `json:"version"`
	Type          string `json:"type"`
	ServerVersion string `json:"serverVersion"`
	License       *LicenseValidity
}

type APIResponse struct {
	Response *SubsonicResponse `json:"subsonic-response"`
}

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
	io.WriteString(h, password)
	io.WriteString(h, salt)
	s.salt = salt
	s.token = fmt.Sprintf("%x", h.Sum(nil))
	// Test authentication
	if !s.Ping() {
		return errors.New("Invalid authentication parameters!")
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
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return nil, err
	}
	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	log.Printf("%s\n", contents)
	return contents, nil
}

func (s *SubsonicClient) Ping() bool {
	_, err := s.Request("GET", "ping", nil)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func (s *SubsonicClient) GetLicense() *SubsonicResponse {
	contents, err := s.Request("GET", "getLicense", nil)
	if err != nil {
		log.Println(err)
		return nil
	}
	resp := APIResponse{}
	err = json.Unmarshal(contents, &resp)
	if err != nil {
		log.Println(err)
		return nil
	}
	return resp.Response
}

func main() {
	client := SubsonicClient{
		client:     &http.Client{},
		BaseUrl:    "http://192.168.1.7:4040/",
		User:       "test",
		ClientName: "go-subsonic_" + libraryVersion,
	}
	err := client.Authenticate("blah")
	if err != nil {
		log.Fatal(err)
	}
	lic := client.GetLicense()
	fmt.Printf("%#v\n", lic)
}
