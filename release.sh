#! /usr/bin/env sh

if [ -f "$HOME/github-release-config.sh" ]; then
    . "$HOME/github-release-config.sh"
fi
cd "$HOME/go/src/github.com/eyedeekay/go-I2P-jpackage"
git pull upstream main
make clean jpackage release upload