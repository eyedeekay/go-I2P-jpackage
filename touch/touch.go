package main

import (
	"go/build"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := filepath.Join(build.Default.GOPATH, "src", "github.com", "eyedeekay", "go-I2P-jpackage")
	if f, err := os.Stat(filepath.Join(dir, "build.I2P.tar.gz")); err == nil {
		if f.Size() < 10 {
			os.Remove(filepath.Join(dir, "build.I2P.tar.gz"))
			ioutil.WriteFile(filepath.Join(dir, "build.I2P.tar.gz"), []byte(""), 0644)
			return
		}
	} else {
		os.Remove(filepath.Join(dir, "build.I2P.tar.gz"))
		ioutil.WriteFile(filepath.Join(dir, "build.I2P.tar.gz"), []byte(""), 0644)
		return
	}
}
