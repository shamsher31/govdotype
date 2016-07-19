// Package vdotype return type of video based on mime type
package vdotype // import "github.com/shamsher31/govdotype"

import (
	"errors"
	"github.com/shamsher31/govdoext"
	"log"
	"net/http"
	"os"
	"strings"
)

// Get returns the type of Video
func Get(p string) (string, error) {
	file, err := os.Open(p)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	buff := make([]byte, 512)

	_, err = file.Read(buff)

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	filetype := http.DetectContentType(buff)

	ext := vdoext.Get()

	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], filetype[6:len(filetype)]) {
			return filetype, nil
		}
	}

	return "", errors.New("Invalid video type")

}
