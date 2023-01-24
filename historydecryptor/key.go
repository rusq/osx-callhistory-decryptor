package historydecryptor

import (
	"encoding/base64"
	"errors"
)

var ErrNoKey = errors.New("key is not supplied")

// DecodeB64Key decodes the provided key from base64 encoding
func DecodeB64Key(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, ErrNoKey
	}
	ret := make([]byte, base64.StdEncoding.DecodedLen(len(key)))
	l, err := base64.StdEncoding.Decode(ret, key)
	if err != nil {
		return nil, err
	}
	return ret[:l], nil
}
