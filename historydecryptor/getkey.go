// +build !darwin

package historydecryptor

import (
	"fmt"
)

func GetByteKey(keyStr string) ([]byte, error) {
	if len(keyStr) == 0 {
		return nil, fmt.Errorf("Use -k <key> parameter to supply the key.")
	}
	key, err := DecodeB64Key([]byte(strKey))
	if err != nil {
		return nil, err
	}
	return key, nil
}
