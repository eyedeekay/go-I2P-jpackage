package I2P

import (
	"embed"
	"io/ioutil"
	"log"
)

//go:embed license/*
var License embed.FS

func GetLicenses() (string, error) {
	// loop over directory contents of the embed.FS License and
	// return a string of all the files contents combined.
	// This is used to display the license agreement to the user.
	var s string
	licenses, err := License.ReadDir("license")
	if err != nil {
		return "", err
	}
	for _, license := range licenses {
		if license.IsDir() {
			continue
		}
		file, err := License.Open("license/" + license.Name())
		if err != nil {
			continue
		}
		data, err := ioutil.ReadAll(file)
		if err != nil {
			continue
		}
		s += string(data)
	}
	return s, nil
}

func PrintLicenses() {
	s, err := GetLicenses()
	if err != nil {
		log.Println(err)
	}
	log.Println(s)
}
