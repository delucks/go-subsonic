package subsonic

import "fmt"

type AlbumList struct {
	Albums []*Album `json:"album"`
}

func validateListType(input string) bool {
	validTypes := map[string]bool{
		"random":               true,
		"newest":               true,
		"highest":              true,
		"frequent":             true,
		"recent":               true,
		"alphabeticalByName":   true,
		"alphabeticalByArtist": true,
		"starred":              true,
		"byYear":               true,
		"byGenre":              true,
	}
	_, ok := validTypes[input]
	return ok
}

// GetAlbumList returns a list of random, newest, highest rated etc. albums. Similar to the album lists on the home page of the Subsonic web interface.
// Optional Parameters:
// size          No              10      The number of albums to return. Max 500.
// offset        No              0       The list offset. Useful if you for example want to page through the list of newest albums.
// fromYear      Yes (if type is         The first year in the range. If fromYear > toYear a reverse chronological list is returned.
//               byYear)
// toYear        Yes (if type is         The last year in the range.
//               byYear)
// genre         Yes (if type is         The name of the genre, e.g., "Rock".
//               byGenre)
// musicFolderId No                      (Since 1.11.0) Only return albums in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetAlbumList(listType string, parameters map[string]string) ([]*Album, error) {
	if !validateListType(listType) {
		return nil, fmt.Errorf("List type %s is invalid, see http://www.subsonic.org/pages/api.jsp#getAlbumList", listType)
	}
	params := make(map[string]string)
	params["type"] = listType
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getAlbumList", params)
	if err != nil {
		return nil, err
	}
	return resp.AlbumList.Albums, nil
}

// GetAlbumList2 returns a list of albums like GetAlbumList, but organized according to id3 tags.
// Optional Parameters:
// size          No              10      The number of albums to return. Max 500.
// offset        No              0       The list offset. Useful if you for example want to page through the list of newest albums.
// fromYear      Yes (if type is         The first year in the range. If fromYear > toYear a reverse chronological list is returned.
//               byYear)
// toYear        Yes (if type is         The last year in the range.
//               byYear)
// genre         Yes (if type is         The name of the genre, e.g., "Rock".
//               byGenre)
// musicFolderId No                      (Since 1.11.0) Only return albums in the music folder with the given ID. See getMusicFolders.
func (s *SubsonicClient) GetAlbumList2(listType string, parameters map[string]string) ([]*Album, error) {
	if !validateListType(listType) {
		return nil, fmt.Errorf("List type %s is invalid, see http://www.subsonic.org/pages/api.jsp#getAlbumList", listType)
	}
	params := make(map[string]string)
	params["type"] = listType
	for k, v := range parameters {
		params[k] = v
	}
	resp, err := s.Get("getAlbumList2", params)
	if err != nil {
		return nil, err
	}
	return resp.AlbumList2.Albums, nil
}