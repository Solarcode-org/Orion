#!/bin/sh
set -e
rm -rf manpages
mkdir manpages

go run . manpages 

for FILE in "manpages"/*; do
    if [ -f "$FILE" ]; then
        gzip -c -9 "$FILE" >"$FILE".gz
    fi
done
