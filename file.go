package fixture

import (
	"io/ioutil"
)

// MakeFile creates a temporary file and returns its full path.
func MakeFile() string {
	file, _ := ioutil.TempFile("", "fixture")
	file.Close()

	return file.Name()
}
