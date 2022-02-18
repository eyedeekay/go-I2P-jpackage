package I2P

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/walle/targz"
)

//go:embed build.I2P.tar.gz
var Content embed.FS

func Unpack(dir string) error {
	//untar build.I2P.tar.gz to a directory specified by -dir
	if err := SetEnv(dir); err != nil {
		return err
	}
	iname := "build.I2P.tar.gz"
	fname := "jpackage.I2P.tar.gz"
	dir, err := filepath.Abs(dir)
	if err != nil {
		return fmt.Errorf("Unpack: Abs() failed: %s", err.Error())
	}
	os.MkdirAll(dir, 0755)
	fpath := filepath.Join(dir, fname)
	//tfpath := filepath.Join(dir, tfname)
	ufpath := filepath.Join(dir, "I2P")
	log.Println("Unpacking", fpath, "to", ufpath)
	file, err := Content.Open(iname)
	if err != nil {
		return fmt.Errorf("Unpack: Open() failed: %s", err.Error())
	}
	defer file.Close()
	// read the file into a byte slice
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return fmt.Errorf("Unpack: ReadAll() failed: %s", err.Error())
	}
	err = ioutil.WriteFile(fpath, data, 0644)
	if err != nil {
		return fmt.Errorf("Unpack: Write() failed: %s", err.Error())
	}
	err = UnTarGzip(fpath, ufpath)
	if err != nil {
		return fmt.Errorf("Unpack: UnTarGzip() failed: %s", err.Error())
	}
	log.Println("Removing", fpath)
	err = os.Remove(fpath)
	if err != nil {
		return fmt.Errorf("Unpack: Remove() failed: %s", err.Error())
	}
	return nil
}

func UnTarGzip(source, target string) error {
	return targz.Extract(source, target)
}
