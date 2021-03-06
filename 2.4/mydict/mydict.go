package mydict

import "errors"

// Dictionary map
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")

// Search value
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}
