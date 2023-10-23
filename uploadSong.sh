#!/bin/bash

sanitizeInput() {
  local input="$1"
  sanitized="${input//[^a-zA-Z0-9 ]/}"
  sanitized="${sanitized// /}"
  sanitized="${sanitized//mp3/}"
  echo "$sanitized"
}

if [ $# -ne 1 ]; then
  echo "Usage: $0 <URL_Youtube>"
  exit 1
fi

URL_DEL_VIDEO="$1"
output_file="song.mp3"
yt-dlp --extract-audio --audio-format mp3 --add-metadata --embed-thumbnail --output "$output_file" "$URL_DEL_VIDEO"

if [ $? -eq 0 ]; then
  artist=$(exiftool -Artist "$output_file" | cut -d ":" -f 2 | xargs)
  title=$(exiftool -Title "$output_file" | cut -d ":" -f 2 | xargs)

  sanitized_artist=$(sanitizeInput "$artist")
  sanitized_title=$(sanitizeInput "$title")

  new_filename="/etc/music/${sanitized_artist}${sanitized_title}"
  mv "$output_file" "$new_filename"

  echo "Download completato con successo. File rinominato in: $new_filename"
else
  echo "Errore durante il download del video da YouTube."
  exit 1
fi

curl -X POST -F "song=@$new_filename" "http://localhost:4000/songs"