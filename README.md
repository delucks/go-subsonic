# go-subsonic

[![GoDoc](https://godoc.org/github.com/delucks/go-subsonic?status.svg)](https://godoc.org/github.com/delucks/go-subsonic)

This is an API client library for Subsonic and Subsonic-compatible music servers. It is tested on Subsonic, Airsonic, and Navidrome.

# Testing

Tests for this library run live against instances of Subsonic-compatible servers. A docker-compose setup comes with this repo to run instances of Airsonic and Navidrome for testing purposes. To set up the test environment and run tests against the docker containers, run `test.sh`. This test runner does the following:

- Download CC-licensed sample music into `build/music`
- Set up a location for configuration for both servers in `build/data`
- Run `docker-compose up` to initialize both servers with the music & configuration directories as volumes
- Wait until the servers are up
- Run tests for both Subsonic and Navidrome

The tests for Subsonic run against the demo server, so please be mindful to exclude them from all but final test runs in respect for their server load. To manually tests for only one server type, use the `-run` argument to `go test` like so:

```
# Just Airsonic
go test . -test.v -run Airsonic
# Subsonic and Airsonic
go test . -test.v -run '(Sub|Air)sonic'
# Navidrome
go test . -test.v -run Navidrome
```

If you intend on testing with your own server, modify the top-level TestSubsonic, TestAirsonic, or TestNavidrome functions with your server's address and credentials. Please test this library with other Subsonic-compatible streaming servers, and file an issue if something is amiss!

# API Support

## System

- [x] ping (1.0.0)
- [x] getLicense (1.0.0)

## Browsing

- [x] getMusicFolders (1.0.0)
- [x] getIndexes (1.0.0)
- [x] getMusicDirectory
- [x] getGenres (1.9.0)
- [x] getArtists (1.8.0)
- [x] getArtist (1.8.0)
- [x] getAlbum (1.0.0)
- [x] getSong (1.8.0)
- [x] getArtistInfo (1.11.0)
- [x] getArtistInfo2 (1.11.0)
- [x] getAlbumInfo (1.14.0)
- [x] getAlbumInfo2 (1.14.0)
- [x] getSimilarSongs (1.11.0)
- [x] getSimilarSongs2 (1.11.0)
- [x] getTopSongs (1.13.0)

## Album/song lists

- [x] getAlbumList (1.2.0)
- [x] getAlbumList2 (1.8.0)
- [x] getRandomSongs (1.2.0)
- [x] getSongsByGenre (1.9.0)
- [x] getNowPlaying (1.0.0)
- [x] getStarred (1.8.0)
- [x] getStarred2 (1.8.0)

## Searching

- [x] search2 (1.4.0)
- [x] search3 (1.8.0)

## Playlists

- [x] getPlaylists (1.0.0)
- [x] getPlaylist (1.0.0)
- [x] createPlaylist (1.2.0)
- [x] updatePlaylist (1.8.0)
- [x] deletePlaylist (1.2.0)

## Media retrieval

- [x] stream (1.0.0)
- [x] download (1.0.0)
- [x] getCoverArt (1.0.0)
- [ ] getLyrics (1.2.0)
- [x] getAvatar (1.8.0)

## Media annotation

- [x] star (1.8.0)
- [x] unstar (1.8.0)
- [x] setRating (1.6.0)
- [x] scrobble (1.5.0)

## User management

- [x] getUser (1.3.0)
- [x] getUsers (1.8.0)
- [x] createUser (1.1.0)
- [x] updateUser (1.10.1)
- [x] deleteUser (1.3.0)
- [x] changePassword (1.1.0)

## Media library scanning

- [x] getScanStatus (1.15.0)
- [x] startScan (1.15.0)

## Bookmarks

- [ ] getBookmarks (1.9.0)
- [ ] createBookmark (1.9.0)
- [ ] deleteBookmark (1.9.0)
- [ ] getPlayQueue (1.12.0)
- [ ] savePlayQueue (1.12.0)

## Sharing

- [ ] getShares (1.6.0)
- [ ] createShare (1.6.0)
- [ ] updateShare (1.6.0)
- [ ] deleteShare (1.6.0)

## Podcast

- [ ] getPodcasts (1.6.0)
- [ ] getNewestPodcasts (1.13.0)
- [ ] refreshPodcasts (1.9.0)
- [ ] createPodcastChannel (1.9.0)
- [ ] deletePodcastChannel (1.9.0)
- [ ] deletePodcastEpisode (1.9.0)
- [ ] downloadPodcastEpisode (1.9.0)

## Jukebox

- [ ] jukeboxControl (1.2.0)

## Internet radio

- [ ] getInternetRadioStations (1.9.0)
- [ ] createInternetRadioStation (1.16.0)
- [ ] updateInternetRadioStation (1.16.0)
- [ ] deleteInternetRadioStation (1.16.0)

## Chat

- [ ] getChatMessages (1.2.0)
- [ ] addChatMessage (1.2.0)

## Out of Scope

Video endpoints are currently out of scope- please file an issue if you would like support for them. The deprecated "search" endpoint is unimplemented.

- [ ] hls.m3u8 (1.8.0)
- [ ] getVideos (1.8.0)
- [ ] getVideoInfo (1.14.0)
- [ ] getCaptions (1.14.0)
- [ ] search (1.0.0)
