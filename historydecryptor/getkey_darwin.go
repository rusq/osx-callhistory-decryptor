// +build darwin

package historydecryptor

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

// GetByteKey decodes the key.  If keyStr is not set, it attempts to get the key
// from the keychain.  If it is set - it's better be a base64 encoded key!
func GetByteKey(keyStr string) (key []byte, err error) {
	if len(keyStr) == 0 {
		key, err = getCallHistoryPwd(true)
	} else {
		key, err = DecodeB64Key([]byte(keyStr))
	}
	if err != nil {
		return nil, err
	}
	return key, nil
}

func getCallHistoryPwd(interactive bool) ([]byte, error) {
	if interactive {
		fmt.Printf("Enter your account password to access keychain when prompted...\n")
	}
	passwd, err := keychain.GetGenericPassword("Call History User Data Key", "", "Call History User Data Key", "")
	if err != nil {
		return nil, err
	}
	key, err := DecodeB64Key(passwd)
	if err != nil {
		return nil, err
	}
	return key, nil
}
