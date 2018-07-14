package main

import (
	"reflect"
	"testing"
)

func Test_decodeB64Key(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"", args{[]byte("U2FtcGxlIGtleQo=")}, []byte("Sample key\n"), false},
		{"", args{[]byte("")}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := decodeB64Key(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("decodeB64Key() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decodeB64Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
