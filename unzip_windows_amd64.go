package I2P

import (
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
)

//go:embed build.windows.I2P.tar.xz
var Content embed.FS

func (d *Daemon) Unpack() error {
	//untar build.I2P.tar.xz to a directory specified by -dir
	if err := SetEnv(d.Dir); err != nil {
		return err
	}
	iname := "build." + runtime.GOOS + ".I2P.tar.xz"
	fname := "jpackage.I2P.tar.xz"
	dir, err := filepath.Abs(d.Dir)
	if err != nil {
		return fmt.Errorf("Unpack: Abs() failed: %s", err.Error())
	}
	os.MkdirAll(dir, 0755)
	fpath := filepath.Join(dir, fname)
	//tfpath := filepath.Join(dir, tfname)
	ufpath := filepath.Join(d.Dir)
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
	err = UnTarXzip(fpath, ufpath)
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
