package vdotype

import (
	"strings"
	"testing"

	"github.com/shamsher31/govdoext"
)

func TestGet(t *testing.T) {
	vdoType, err := Get("_testdata/test-video.mp4")

	if err != nil {
		t.Fatal(err)
	}

	if isVdoTypeFromExt(vdoType) == false {
		t.Fatal("Not a valid video type")
	}
}

func isVdoTypeFromExt(p string) bool {
	ext := vdoext.Get()
	for i := 0; i < len(ext); i++ {
		if strings.Contains(ext[i], p[6:len(p)]) {
			return true
		}
	}
	return false
}
