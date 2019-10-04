/*
OS X Call history decryptor
Copyright (C) 2016  n0fate (GPL2 license)
Copyright (C) 2018  rusq (golang implementation)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program.  If not, see <https://www.gnu.org/licenses/>.
*/
package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rusq/callhistory/historydecoder"
)

const (
	version = "1.1"
)

var (
	envKey = os.Getenv("KEY")

	strKey         string //key from Keychain
	filename       string // call history database filename
	outputFilename string // output filename
	versionOnly    bool   // output version and quit
)

func init() {
	keyMsg := "Base64 key value from OS X keychain."
	flag.StringVar(&strKey, "k", envKey, keyMsg)
	homeDir := os.Getenv("HOME")
	flag.StringVar(&filename, "f", "CallHistory.storedata", "filename with call data. Get it from:\n"+homeDir+"/Library/Application Support/CallHistoryDB/\n")
	flag.StringVar(&outputFilename, "o", "", "output csv filename.  If not specified, result is output to stdout")
	flag.BoolVar(&versionOnly, "v", false, "print version and quit")
}

func printHeader() {
	fmt.Fprintf(os.Stderr, "MacOS X Call History Decryptor v.%s © 2018 rusq\n"+
		"Based on Call History Decryptor © 2016 n0fate\n",
		version)
}

func main() {
	flag.Parse()

	printHeader()
	if versionOnly {
		return
	}

	key, err := getByteKey(strKey)
	if err != nil {
		log.Fatal(err)
	}

	outFile := os.Stdout
	if outputFilename != "" || outputFilename == "-" {
		outFile, err := os.Create(outputFilename)
		if err != nil {
			log.Fatal(err)
		}
		defer outFile.Close()
	}

	fmt.Printf("Starting. Filename: %s\n", filename)

	numRecords, err := historydecoder.DecipherHistory(filename, key, outFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFinished. %d records processed\n", numRecords)
}

// decodeB64Key decodes the provided key from base64 encoding
func decodeB64Key(key []byte) ([]byte, error) {
	if len(key) == 0 {
		return nil, fmt.Errorf("Empty key")
	}
	ret := make([]byte, base64.StdEncoding.DecodedLen(len(key)))
	if l, err := base64.StdEncoding.Decode(ret, key); err != nil {
		return nil, err
	} else {
		ret = ret[:l]
	}
	return ret, nil
}
