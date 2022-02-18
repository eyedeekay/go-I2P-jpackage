package main

import (
	"flag"
	"log"
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
		err := I2P.Generate(dir)
		if err != nil {
			log.Fatal(err)
		}
		return
	}
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
