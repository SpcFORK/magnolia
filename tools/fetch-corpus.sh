#!/usr/bin/env bash
# fetch-corpus.sh -- download and verify the JS ablation corpus.
#
# Corpora are NOT committed to the repository; they are fetched on demand
# into build/corpus/ (which is .gitignored) and verified against pinned
# SHA-256 hashes. Small hand-written fixtures live in test/fixtures/js/.
#
# Usage: ./tools/fetch-corpus.sh [corpus-dir]
set -euo pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
CORPUS_DIR="${1:-$ROOT/build/corpus}"
mkdir -p "$CORPUS_DIR"
cd "$CORPUS_DIR"

fetch() { # name url sha256
	local name="$1" url="$2" sha="$3"
	local archive="$CORPUS_DIR/$name.tar.gz"
	if [ ! -f "$archive" ]; then
		echo ">> fetching $name from $url"
		curl -sL --fail -o "$archive" "$url"
	else
		echo ">> reusing $archive"
	fi
	echo "$sha  $archive" | sha256sum --check --quiet - \
		&& echo ">> $name sha256 OK"
}

# --- jQuery 1.0 (GitHub tag "1.0"; npm starts at 1.5.1) -------------------
fetch jquery-1.0 \
	"https://codeload.github.com/jquery/jquery/tar.gz/refs/tags/1.0" \
	"2eb4c7744a8f381c545418345809edc6fac14f6c47bdc2584d1326cd29578a2a"
if [ ! -f "$CORPUS_DIR/jquery-1.0/src/jquery/jquery.js" ]; then
	tar xzf "$CORPUS_DIR/jquery-1.0.tar.gz" -C "$CORPUS_DIR" jquery-1.0/src
fi
# Source concatenation order per the jQuery 1.0 Makefile/build.xml
cat "$CORPUS_DIR"/jquery-1.0/src/jquery/*.js \
	"$CORPUS_DIR"/jquery-1.0/src/event/*.js \
	"$CORPUS_DIR"/jquery-1.0/src/fx/*.js \
	"$CORPUS_DIR"/jquery-1.0/src/ajax/*.js \
	> "$CORPUS_DIR/jquery-1.0/jquery-1.0.js"
# Dist builds historically substituted the $Rev$ placeholder with the release
# number; do the same for the smoke corpus (jquery.fn.jquery === '1.0').
sed -i 's/\$Rev\$/1.0/g' "$CORPUS_DIR/jquery-1.0/jquery-1.0.js"

# --- three.js r58 (stretch corpus; cf. threejs-oak-complete.patch.txt) ----
fetch three-r58 \
	"https://codeload.github.com/mrdoob/three.js/tar.gz/refs/tags/r58" \
	"076988d415e2c3e88c23a81eb9979c03f1b73b7866657c9cbedc324fe88b81c0"
if [ ! -f "$CORPUS_DIR/three-r58/build/three.js" ]; then
	tar xzf "$CORPUS_DIR/three-r58.tar.gz" -C "$CORPUS_DIR" three.js-r58/build
	mv "$CORPUS_DIR/three.js-r58" "$CORPUS_DIR/three-r58"
fi

echo ">> corpus ready at $CORPUS_DIR"
