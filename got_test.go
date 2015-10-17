package main

import (
	"io/ioutil"
	"testing"
)

func TestMain(t *testing.T) {
	files, err := ioutil.ReadDir("cmd")
	if err != nil {
		t.Fatal(err)
	}

	if len(app.Commands) != len(files) {
		t.Fatal("app.Commands size is not equal command file size")
	}
}
