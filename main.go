/*
OS X Call history decryptor
Copyright (C) 2016  n0fate (GPL2 license)
Copyright (C) 2019  rusq   (golang implementation)

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
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/rusq/osx-callhistory-decryptor/historydecryptor"
)

var defFile = filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "CallHistoryDB", "CallHistory.storedata")

var (
	strKey         = flag.String("k", os.Getenv("KEY"), "Base64 key value from OS X keychain, on macOS may be omitted.")
	omitDecryption = flag.Bool("no-key", false, "omit decryption")
	outputFilename = flag.String("o", "", "output csv filename.  If not specified, result is output to stdout")
	versionOnly    = flag.Bool("v", false, "print version and quit")
	timeFormat     = flag.String("time-format", historydecryptor.DefTimeFmt, "CSV output time `format`")

	build = "v.0.0-development"
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [flags] [filename]\n"+
			"If filename is not specified, will try to locate it at the default location\n(%s)\n"+
			"The program will not modify the original file, and will operate on a copy in a\n"+
			"temporary directory\n\n"+
			"Flags:\n",
			os.Args[0], defFile)
		flag.PrintDefaults()
	}
}

func printHeader() {
	fmt.Fprintf(os.Stderr, "MacOS X Call History Decryptor %s © 2018-2021 rusq\n"+
		"Based on Call History Decryptor © 2016 n0fate\n",
		build)
}

func main() {
	flag.Parse()

	printHeader()
	if *versionOnly {
		return
	}

	var src string
	if flag.NArg() > 0 {
		src = flag.Arg(0)
	} else {
		src = defFile
	}

	if err := run(src, *outputFilename, *strKey, *omitDecryption); err != nil {
		log.Fatal(err)
	}
}

func run(src, dst string, strKey string, omitDecryption bool) error {
	log.Printf("*** database filename: %q", src)
	dbfile, err := copytemp(src)
	if err != nil {
		return err
	}
	defer os.Remove(dbfile)
	log.Printf("*** temporary file (will be removed): %q", dbfile)

	var key []byte
	if omitDecryption {
		key = []byte{}
	} else {
		key, err = historydecryptor.GetByteKey(strKey)
		if err != nil {
			return fmt.Errorf("%w: make sure you have supplied the key via -k <key> or KEY env variable", err)
		}
	}

	var output = os.Stdout
	if dst != "" && dst != "-" {
		f, err := os.Create(dst)
		if err != nil {
			return fmt.Errorf("unable to create the output file: %w", err)
		}
		defer f.Close()
		output = f
	}

	numRecords, err := historydecryptor.DecipherHistory(dbfile, key, output, historydecryptor.OptTimeFormat(*timeFormat))
	if err != nil {
		return err
	}

	log.Printf("*** finished. %d records processed\n", numRecords)
	return nil
}

// copytemp copies the file to a temporary location and returns its name.
func copytemp(filename string) (string, error) {
	src, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer src.Close()

	dst, err := ioutil.TempFile("", "callhistory*")
	if err != nil {
		return "", err
	}
	defer dst.Close()
	if n, err := io.Copy(dst, src); err != nil {
		return "", err
	} else if n == 0 {
		return "", fmt.Errorf("copy: %s: file has zero size", filename)
	}
	return dst.Name(), nil
}
