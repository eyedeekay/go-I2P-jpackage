package I2P

import (
	"embed"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

//go:embed build.I2P.tar.gz
var Content embed.FS

func Unpack(dir string) {
	//untar build.I2P.tar.gz to a directory specified by -dir
	flag.Parse()
	iname := "build.I2P.tar.gz"
	fname := "jpackage.I2P.tar.gz"
	tfname := "jpackage.I2P.tar"
	os.MkdirAll(dir, 0755)
	fpath := filepath.Join(dir, fname)
	tfpath := filepath.Join(dir, tfname)
	ufpath := filepath.Join(dir, "I2P")
	file, err := Content.Open(iname)
	if err != nil {
		log.Fatal("main: Open() failed: ", err.Error())
	}
	defer file.Close()
	// read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("main: ReadAll() failed: ", err.Error())
	}
	err = ioutil.WriteFile(fpath, data, 0644)
	if err != nil {
		log.Fatal("main: Write() failed: ", err.Error())
	}
	err = UnGzip(fpath, tfpath)
	if err != nil {
		log.Fatal("main: Untar() failed: ", err.Error())
	}
	err = Untar(tfpath, ufpath)
	if err != nil {
		log.Fatal("main: Untar() failed: ", err.Error())
	}
	err = os.Remove(fpath)
	if err != nil {
		log.Fatal("main: Remove() failed: ", err.Error())
	}
	err = os.Remove(tfpath)
	if err != nil {
		log.Fatal("main: Remove() failed: ", err.Error())
	}
}
