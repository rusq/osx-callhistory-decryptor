// +build !darwin

package main

import (
	"fmt"
)

func getByteKey(keyStr string) ([]byte, error) {
	var key []byte
	var err error
	if len(strKey) == 0 {
		return nil, fmt.Errorf("Use -k <key> parameter to supply the key.")
	} else {
		key, err = decodeB64Key([]byte(strKey))
	}
	if err != nil {
		return nil, err
	}
	return key, nil
}
