package auth

import (
	"crypto/aes"
	"reflect"
	"testing"
)

func TestNewAesCbcPkcs7Cipher(t *testing.T) {
	successBlock, _ := aes.NewCipher([]byte("xtRMca7jhJCGVwfq"))

	type args struct {
		key []byte
	}
	tests := []struct {
		name    string
		args    args
		want    *AesCbcPkcs7Cipher
		wantErr bool
	}{
		{
			name: "Success create 128bit key of Aes",
			args: args{
				key: []byte("xtRMca7jhJCGVwfq"),
			},
			want: &AesCbcPkcs7Cipher{
				block: successBlock,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewAesCbcPkcs7Cipher(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewAesCbcPkcs7Cipher() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAesCbcPkcs7Cipher() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcPkcs7Cipher_pad(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		c    *AesCbcPkcs7Cipher
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.pad(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcPkcs7Cipher.pad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcPkcs7Cipher_unpad(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name string
		c    *AesCbcPkcs7Cipher
		args args
		want []byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.unpad(tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcPkcs7Cipher.unpad() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcPkcs7Cipher_Encrypt(t *testing.T) {
	type args struct {
		plain []byte
	}
	tests := []struct {
		name    string
		c       *AesCbcPkcs7Cipher
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Encrypt(tt.args.plain)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcPkcs7Cipher.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AesCbcPkcs7Cipher.Encrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAesCbcPkcs7Cipher_Decrypt(t *testing.T) {
	type args struct {
		encryptedTextBase64 []byte
	}
	tests := []struct {
		name    string
		c       *AesCbcPkcs7Cipher
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.c.Decrypt(tt.args.encryptedTextBase64)
			if (err != nil) != tt.wantErr {
				t.Errorf("AesCbcPkcs7Cipher.Decrypt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("AesCbcPkcs7Cipher.Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
