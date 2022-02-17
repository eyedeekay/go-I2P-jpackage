package I2P

import (
	"archive/tar"
	"compress/gzip"
	"embed"
	"flag"
	"io"
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
	dir, err := filepath.Abs(dir)
	if err != nil {
		log.Fatal(err)
	}
	os.MkdirAll(dir, 0755)
	fpath := filepath.Join(dir, fname)
	tfpath := filepath.Join(dir, tfname)
	ufpath := filepath.Join(dir, "I2P")
	log.Println("Unpacking", fpath, "to", ufpath)
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
	log.Println("Unpacking", fpath, "to", tfpath)
	err = UnGzip(fpath, tfpath)
	if err != nil {
		log.Fatal("main: UnGzip() failed: ", err.Error())
	}
	log.Println("Unpacking", tfpath, "to", ufpath)
	err = Untar(tfpath, ufpath)
	if err != nil {
		log.Fatal("main: Untar() failed: ", err.Error())
	}
	log.Println("Removing", fpath)
	err = os.Remove(fpath)
	if err != nil {
		log.Fatal("main: Remove() failed: ", err.Error())
	}
	log.Println("Removing", tfpath)
	err = os.Remove(tfpath)
	if err != nil {
		log.Fatal("main: Remove() failed: ", err.Error())
	}
}

func Untar(tarball, target string) error {
	reader, err := os.Open(tarball)
	if err != nil {
		return err
	}
	defer reader.Close()
	tarReader := tar.NewReader(reader)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		} else if err != nil {
			return err
		}

		path := filepath.Join(target, header.Name)
		info := header.FileInfo()
		if info.IsDir() {
			if err = os.MkdirAll(path, info.Mode()); err != nil {
				return err
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
			return err
		}
	}
	return nil
}

func UnGzip(source, target string) error {
	reader, err := os.Open(source)
	if err != nil {
		return err
	}
	defer reader.Close()

	archive, err := gzip.NewReader(reader)
	if err != nil {
		return err
	}
	defer archive.Close()

	target = filepath.Join(target, archive.Name)
	writer, err := os.Create(target)
	if err != nil {
		return err
	}
	defer writer.Close()

	_, err = io.Copy(writer, archive)
	return err
}
