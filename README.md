# go-subsonic

This is an API client library for Subsonic and Subsonic-compatible music servers. It has been tested on Subsonic, Airsonic, and Navidrome.

Unless support is requested, video endpoints will be out of scope.

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
- [x] getAlbum
- [x] getSong (1.8.0)
- [x] getArtistInfo (1.11.0)
- [x] getArtistInfo2 (1.11.0)
- [x] getAlbumInfo (1.14.0)
- [x] getAlbumInfo2 (1.14.0)
- [x] getSimilarSongs (1.11.0)
- [x] getSimilarSongs2 (1.11.0)
- [x] getTopSongs (1.13.0)
- [ ] getVideos (1.8.0)
- [ ] getVideoInfo (1.14.0)

## Album/song lists

- [x] getAlbumList (1.2.0)
- [x] getAlbumList2 (1.8.0)
- [ ] getRandomSongs (1.2.0)
- [ ] getSongsByGenre (1.9.0)
- [ ] getNowPlaying (1.0.0)
- [ ] getStarred (1.8.0)
- [ ] getStarred2 (1.8.0)

## Searching

- [ ] search (1.0.0)
- [ ] search2 (1.4.0)
- [ ] search3 (1.8.0)

## Playlists

- [ ] getPlaylists (1.0.0)
- [ ] getPlaylist (1.0.0)
- [ ] createPlaylist (1.2.0)
- [ ] updatePlaylist (1.8.0)
- [ ] deletePlaylist (1.2.0)

## Media retrieval

- [ ] stream (1.0.0)
- [ ] download (1.0.0)
- [ ] hls.m3u8 (1.8.0)
- [ ] getCaptions (1.14.0)
- [ ] getCoverArt (1.0.0)
- [ ] getLyrics (1.2.0)
- [ ] getAvatar (1.8.0)

## Media annotation

- [ ] star (1.8.0)
- [ ] unstar (1.8.0)
- [ ] setRating (1.6.0)
- [ ] scrobble (1.5.0)

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

## User management

- [ ] getUser (1.3.0)
- [ ] getUsers (1.8.0)
- [ ] createUser (1.1.0)
- [ ] updateUser (1.10.1)
- [ ] deleteUser (1.3.0)
- [ ] changePassword (1.1.0)

## Bookmarks

- [ ] getBookmarks (1.9.0)
- [ ] createBookmark (1.9.0)
- [ ] deleteBookmark (1.9.0)
- [ ] getPlayQueue (1.12.0)
- [ ] savePlayQueue (1.12.0)

## Media library scanning

- [ ] getScanStatus (1.15.0)
- [ ] startScan (1.15.0)
