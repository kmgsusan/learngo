package mydict

import "errors"

// Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errValueExists = errors.New("Exists value")

// Search dict word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

// Add value
func (d Dictionary) Add(word, value string) error {
	_, err := d.Search(word)
	// if err == errNotFound {
	// 	d[word] = word
	// 	return nil
	// }
	// return errValueExists
	switch err {
	case errNotFound:
		d[word] = value
	case nil:
		return errValueExists
	}
	return nil
}
