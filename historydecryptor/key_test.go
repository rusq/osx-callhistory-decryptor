package historydecryptor

import (
	"reflect"
	"testing"
)

func Test_DecodeB64Key(t *testing.T) {
	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		{"case1", args{key: []byte("VGVzdCBrZXkK")}, []byte("Test key\n"), false},
		{"Invalid base64", args{[]byte("xxx")}, nil, true},
		{"Empty key", args{[]byte("")}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeB64Key(tt.args.key)
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
