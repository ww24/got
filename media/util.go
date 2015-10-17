package media

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"
)

// Download method
func Download(url string) (err error) {
	filename := path.Base(url)
	filename = filename[:strings.LastIndex(filename, ":")]

	resp, err := http.Get(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return
	}
	defer file.Close()

	io.Copy(file, resp.Body)

	return
}
