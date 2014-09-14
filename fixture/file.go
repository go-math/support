package fixture

import (
	"io/ioutil"
)

func MakeTempFile() string {
	file, _ := ioutil.TempFile("", "fixture")
	file.Close()

	return file.Name()
}
