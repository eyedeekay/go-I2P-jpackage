#! /usr/bin/env sh

if [ -f "$HOME/github-release-config.sh" ]; then
    . "$HOME/github-release-config.sh"
fi

make clean jpackage release upload