#!/bin/bash

URL="http://localhost:8095/new"

for i in $(seq 1 50); do
  TITLE="TestTitle_$i"
  echo "Sending request $i with title: $TITLE"

  curl -s -o /dev/null -w "HTTP %{http_code}\n" \
    -X POST "$URL" \
    -H "Content-Type: application/json" \
    -d "{
      \"title\": \"$TITLE\",
      \"author\": \"AutoTester\",
      \"content\": \"Some content for note $i\"
    }"

  sleep 0.5  # небольшая пауза, можно убрать
done
