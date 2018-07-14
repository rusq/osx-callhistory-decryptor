// +build darwin

package main

import (
	"fmt"

	keychain "github.com/keybase/go-keychain"
)

func getByteKey(keyStr string) ([]byte, error) {
	var key []byte
	var err error
	if len(keyStr) == 0 {
		key, err = getCallHistoryPwd(true)
	} else {
		key, err = decodeB64Key([]byte(keyStr))
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
	key, err := decodeB64Key(passwd)
	if err != nil {
		return nil, err
	}
	return key, nil
}
