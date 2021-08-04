while read p; do
  curl -X POST -H 'Content-Type: application/json' -d "$p" http://localhost:8080/albums
done <albums.json
