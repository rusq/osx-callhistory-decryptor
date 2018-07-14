// +build !darwin

package main

import (
	"reflect"
	"testing"
)

func Test_getByteKey(t *testing.T) {
	type args struct {
		keyStr string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"case1", args{keyStr: "VGVzdCBrZXkK"}, []byte("Test key\n"), false},
		{"Invalid base64", args{"xxx"}, nil, true},
		{"Empty key", args{""}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getByteKey(tt.args.keyStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("getByteKey() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getByteKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
