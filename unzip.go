package I2P

import (
	"archive/tar"
	"compress/gzip"
	"embed"
	"flag"
	"fmt"
	"io"
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
	flag.Parse()
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

func Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return fmt.Errorf("Untar: Open() failed %s", err.Error())
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("Untar: Next() failed %s", err.Error())
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return fmt.Errorf("Untar: MkdirAll failed %s", err.Error())
			}
			continue
		}

		file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
		if err != nil {
			log.Println(err)
			continue
		}
		log.Println("Untar:", path)
		defer file.Close()
		_, err = io.Copy(file, tarReader)
		if err != nil {
			return fmt.Errorf("Untar: Copy() failed%s", err.Error())
		}
	}
	return nil
}

func UnGzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return fmt.Errorf("UnGzip: Open() failed %s", err.Error())
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return fmt.Errorf("UnGzip: NewReader() failed %s", err.Error())
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return fmt.Errorf("UnGzip: Create() failed%s", err.Error())
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return fmt.Errorf("UnGzip: Copy() failed %s", err.Error())
}
