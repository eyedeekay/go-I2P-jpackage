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
	if err != nil {
		log.Fatal(err)
	}
	I2Pdaemon, err := I2P.NewDaemon(dir, *generate)
	if err != nil {
		log.Println(err)
		return
	}
	err = I2Pdaemon.RunCommand()
	if err != nil {
		log.Fatal(err)
	}
}
