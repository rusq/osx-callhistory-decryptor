package historydecoder

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"reflect"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

const (
	keyStr = "0123456789012345"
)

var testmatrix = [][]string{
	{"Now is the time for all good men to come to the aid of their country.",
		"5B7GQZvdkzifcv6bmRm2jdZEH5xUQGL2Ool68eFnGhs2c+QfOAc84d4oCqCBoWJa16W47j/8beN2QeZ4V1gcYpFE+FfNwtgUigIy0WrdmvvjiyggHcsY58ZBbKRKLQ0j+Bq6mDo="},
	{"The quick brown fox jumps over the lazy dog.",
		"R6ecAkF0S5stCsXvY534a75vg8KIyvAQMQqPEtPKwg5xOWJ1RFkAYZFXVouwreMlxsZRms0vlEsZNSfF3nNDh6Bvz9AVJplBEMau1w=="},
	{"+64220123456", "hNVo3z/cSBbCjpyAeZYKA7g2CF0saEhKlGKKclNFwe8UrxxgBFA/DCvHdyg="},
	{"+79251111111", "w4hQXKHc5aOkuL83kG9lE/ZsGRXsgd0CG/szEo7EGFVPMk7C5NZeAEJovIw="},
	{"0220220202", "ZUOxJzUX2lBY9Ug3EWPQ9Zfx7cggwyMgkmUBRbNl0ZYMRlpL9MYi718K"},
	{"", ""},
}

func TestEnDecipher(t *testing.T) {

	key := []byte(keyStr)
	for _, data := range testmatrix {
		ct, err := Cipher([]byte(data[0]), key)
		if err != nil {
			t.Fatal(err)
		}
		pt, err := Decipher(ct, key)
		if err != nil {
			t.Fatal(err)
		}
		if string(pt) != data[0] {
			t.Fatal("data != pt")
		}
	}
}

func TestDecipher(t *testing.T) {

	for _, data := range testmatrix {
		ct, err := base64.StdEncoding.DecodeString(data[1])
		if err != nil {
			t.Error(err)
		}
		pt, err := Decipher(ct, []byte(keyStr))
		if err != nil {
			t.Error(err)
		}
		if string(pt) != data[0] {
			t.Errorf("failed to decode %s: %s", data[0], string(pt))
		}
	}
}

const (
	sampledbData = `Date,Answered?,Outgoing?,Type,Country,Number/Address
2014-09-30 07:42:04Z,true,false,CellPhone,ru,
2014-10-19 14:32:18Z,true,false,CellPhone,ru,
2014-11-11 13:39:09Z,false,true,CellPhone,ru,
2014-11-13 17:01:21Z,true,false,CellPhone,ru,
2014-11-17 07:43:44Z,false,false,CellPhone,ru,
2014-11-21 08:32:59Z,false,true,CellPhone,ru,
2014-11-21 09:31:40Z,true,false,CellPhone,ru,
2014-11-21 17:23:41Z,false,true,CellPhone,ru,
2015-05-19 05:19:12Z,true,false,CellPhone,nz,
2015-06-10 02:56:01Z,true,false,CellPhone,nz,
`
)

func TestDecipherHistory(t *testing.T) {
	type args struct {
		database string
		key      []byte
	}
	tests := []struct {
		name       string
		args       args
		want       int
		wantOutput string
		wantErr    bool
	}{
		{"Parsing sample db", args{"sample_db/sample.storedata", []byte{0, 0}}, 10, sampledbData, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output := &bytes.Buffer{}
			got, err := DecipherHistory(tt.args.database, tt.args.key, output)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecipherHistory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DecipherHistory() = %v, want %v", got, tt.want)
			}
			if gotOutput := output.String(); gotOutput != tt.wantOutput {
				t.Errorf("DecipherHistory() = %v, want %v", gotOutput, tt.wantOutput)
				fmt.Printf("!!%v!!", gotOutput)
			}
		})
	}
}

func Test_calcCallTime(t *testing.T) {
	type args struct {
		callOffset float64
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{"zero offset", args{0.0}, time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{"negative offset", args{-10.0}, time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{"positive offset", args{1000.0}, time.Date(2001, time.January, 1, 0, 16, 40, 0, time.UTC)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcCallTime(tt.args.callOffset); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("calcCallTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
