package fixture

import (
	"io/ioutil"
)

// MakeTempFile creates a temporary file and returns its full path.
func MakeTempFile() string {
	file, _ := ioutil.TempFile("", "fixture")
	file.Close()

	return file.Name()
}
