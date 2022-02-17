package main

import (
	"flag"

	"github.com/eyedeekay/go-I2P-jpackage"
)

var (
	dir = flag.String("dir", "", "directory to unpack to")
)

func main() {
	I2P.Unpack(*dir)
}
