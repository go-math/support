package fixture

import (
	"io/ioutil"
)

// MakeTemporaryFile creates a temporary file and returns its full path.
func MakeTemporaryFile() string {
	file, _ := ioutil.TempFile("", "fixture")
	file.Close()

	return file.Name()
}
