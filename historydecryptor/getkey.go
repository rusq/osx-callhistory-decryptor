//go:build !darwin
// +build !darwin

package historydecryptor

func GetByteKey(keyStr string) ([]byte, error) {
	if len(keyStr) == 0 {
		return nil, ErrNoKey
	}
	key, err := DecodeB64Key([]byte(keyStr))
	if err != nil {
		return nil, err
	}
	return key, nil
}
