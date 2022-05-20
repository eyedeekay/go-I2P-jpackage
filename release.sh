#! /usr/bin/env sh

if [ -f "$HOME/github-release-config.sh" ]; then
    . "$HOME/github-release-config.sh"
fi
cd "$HOME/go/src/github.com/eyedeekay/go-I2P-jpackage"
git pull upstream main

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     export UPLOAD_OS=linux;;
    Darwin*)    exit 1;;
    CYGWIN*)    export UPLOAD_OS=windows;;
    MINGW*)     export UPLOAD_OS=windows;;
    *)          exit 1;;
esac
echo ${machine}


make jpackage release upload