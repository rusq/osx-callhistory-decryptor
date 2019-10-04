// Package converts call history time
package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/rusq/osx-callhistory-decryptor/historydecryptor"
)

func usage() {
	fmt.Printf("Converts callhistory date.\n"+
		"usage: %s <time from zcallhistory>\n"+
		"i.e.: %s 568354924.058314\n",
		os.Args[0], os.Args[0])
}

func main() {
	if len(os.Args) != 2 {
		usage()
		os.Exit(1)
	}

	date, err := convert(os.Args[1])
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	fmt.Printf("%s\n", date)
}

func convert(floatOffset string) (time.Time, error) {
	t, err := strconv.ParseFloat(floatOffset, 64)
	if err != nil {
		return time.Time{}, err
	}
	return historydecoder.CalcCallTime(t), nil
}
