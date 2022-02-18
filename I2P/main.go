package main

import (
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/eyedeekay/go-I2P-jpackage"
)

var (
	dir      = flag.String("dir", ".", "directory to unpack to")
	generate = flag.Bool("generate", false, "download and compile i2p.firefox to generate the build-dependencies")
)

func main() {
	flag.Parse()
	dir, err := filepath.Abs(*dir)
	if *generate {
		if f, err := os.Stat(filepath.Join(dir, "build.I2P.tar.gz")); err == nil {
			if f.Size() < 10 {
				os.Remove(filepath.Join(dir, "build.I2P.tar.gz"))
				ioutil.WriteFile(filepath.Join(dir, "build.I2P.tar.gz"), []byte(""), 0644)
				return
			}
		}
		err := I2P.Generate(dir)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("I2P", filepath.Join(dir, "I2P"))
	if err != nil {
		log.Fatal(err)
	}
	err = os.Setenv("I2P_CONFIG", filepath.Join(dir, "I2P", "config"))
	if err != nil {
		log.Fatal(err)
	}
	err = I2P.Unpack(dir)
	if err != nil {
		log.Fatal(err)
	}
	err = I2P.RunCommand(dir)
	if err != nil {
		log.Fatal(err)
	}
}
