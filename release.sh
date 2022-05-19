#! /usr/bin/env sh

if [ -f "$HOME/github-release-config.sh" ]; then
    . "$HOME/github-release-config.sh"
fi
git pull upstream main
make clean jpackage release upload