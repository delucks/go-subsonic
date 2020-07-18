#!/usr/bin/env bash

# This script performs automated testing of go-subsonic on Airsonic and Navidrome music servers.
# It wraps docker-compose to download a sample music library before creating instances of these servers,
# allowing the go-subsonic tests to run fully against the sample instances.

err() {
  echo "$1" >&2
  exit 1
}

for dependency in curl docker-compose; do
  hash "$dependency" 2>/dev/null || err "$dependency must be installed"
done

SOURCE_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

download_audionautix() {
  # Audionautix Acoustic is a CC licensed album from Jason Shaw available from Free Music Archive.
  # https://archive.org/details/Audionautix_Acoustic-9870
  BASEURL="https://archive.org/download/Audionautix_Acoustic-9870/Jason_Shaw_-_"
  TRACKS=("12_MORNINGS.ogg" "ACOUSTIC_BLUES.ogg" "FUNKY_JUNKY.ogg" "JENNYS_THEME.ogg" "LAZY_DAY.ogg" "MOUNTAIN_SUN.ogg" "ONE_FINE_DAY.ogg" "RIVER_MEDITATION.ogg" "ROCKY_TOP.ogg" "RUNNING_WATERS.ogg" "SERENITY.ogg" "SIDEWALK.ogg" "SNAPPY.ogg" "SOLO_ACOUSTIC_GUITAR.ogg" "SOUTH_OF_THE_BORDER.ogg" "TENNESEE_HAYRIDE.ogg" "THINGAMAJIG.ogg" "TRAVEL_LIGHT.ogg" "WHEELS.ogg" "WORDS.ogg")
  DESTINATION="${SOURCE_DIR}/build/music/Jason Shaw/Audionautix"
  mkdir -p "$DESTINATION"
  for track in "${TRACKS[@]}"
  do
    if test -f "$DESTINATION/$track"
    then
      echo "Skipping download of Audionautix $track" >&2
    else
      curl -L -o "$DESTINATION/$track" "${BASEURL}${track}"
    fi
  done
}

download_grabbag() {
  # Grab Bag is a CC-licensed jazz album from Jahzzar available from Free Music Archive.
  # https://archive.org/details/Grab_Bag-12446
  BASEURL="https://archive.org/download/Grab_Bag-12446/"
  DESTINATION="${SOURCE_DIR}/build/music/Jahzzar/Grab Bag"
  TRACKS=("01_-_Dummy.ogg" "02_-_Candlelight.ogg" "03_-_Trust.ogg" "04_-_Guilty.ogg" "05_-_Storm.ogg")
  mkdir -p "$DESTINATION"
  for track in "${TRACKS[@]}"
  do
    if test -f "$DESTINATION/$track"
    then
      echo "Skipping download of Grab Bag $track" >&2
    else
      curl -L -o "$DESTINATION/$track" "${BASEURL}${track}"
    fi
  done
}


download_fourseasons() {
  # This is a CC-licensed recording of Vivaldi's The Four Seasons performed by John Harrison with the Wichita State University Chamber Players
  # https://archive.org/details/The_Four_Seasons_Vivaldi-10361
  BASEURL="https://archive.org/download/The_Four_Seasons_Vivaldi-10361/John_Harrison_with_the_Wichita_State_University_Chamber_Players_-_"
  DESTINATION="${SOURCE_DIR}/build/music/Vivaldi/The Four Seasons"
  TRACKS=("01_-_Spring_Mvt_1_Allegro.ogg" "02_-_Spring_Mvt_2_Largo.ogg" "03_-_Spring_Mvt_3_Allegro_pastorale.ogg" "04_-_Summer_Mvt_1_Allegro_non_molto.ogg" "05_-_Summer_Mvt_2_Adagio.ogg" "06_-_Summer_Mvt_3_Presto.ogg" "07_-_Autumn_Mvt_1_Allegro.ogg" "08_-_Autumn_Mvt_2_Adagio_molto.ogg" "09_-_Autumn_Mvt_3_Allegro.ogg" "10_-_Winter_Mvt_1_Allegro_non_molto.ogg" "11_-_Winter_Mvt_2_Largo.ogg" "12_-_Winter_Mvt_3_Allegro.ogg")
  mkdir -p "$DESTINATION"
  for track in "${TRACKS[@]}"
  do
    if test -f "$DESTINATION/$track"
    then
      echo "Skipping download of The Four Seasons $track" >&2
    else
      curl -L -o "$DESTINATION/$track" "${BASEURL}${track}"
    fi
  done
}

download_sample_audio() {
  download_audionautix
  download_grabbag
  download_fourseasons
}

configure_airsonic() {
  cat << DOG > build/data/airsonic.properties
JWTKey=q7q8u331n25gkvgjiehutl3e4u
SettingsChanged=$(date +%s)000
LastScanned=0
IndexCreationInterval=1
DOG
}

main() {
  # Download free CC-licensed music into the build/music directory
  download_sample_audio
  # Create or restart the docker containers of Airsonic and Navidrome
  if [[ $(docker-compose top) ]]
  then
    # If the current composition is running, restart it to pick up possible changes
    docker-compose down
    configure_airsonic  # This must occur in the middle so settings aren't overwritten
    docker-compose up -d
  else
    # Otherwise, bring up the docker containers
    configure_airsonic
    docker-compose up -d
  fi
  # TODO call test function for Airsonic and Navidrome only
}

main
