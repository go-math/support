// Package fixture provides auxiliary routines for testing.
package fixture

import (
	"encoding/json"
	"os"
)

// Load reads data from a JSON file.
func Load(path string, data interface{}) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	if err = dec.Decode(data); err != nil {
		return err
	}

	return nil
}
