package auth

import (
	"crypto/aes"
	"reflect"
	"testing"
)

func TestNewAesCbcPkcs7Cipher(t *testing.T) {
	successBlock, _ := aes.NewCipher([]byte("testtesttesttest"))

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
				key: []byte("testtesttesttest"),
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

func TestAesCbcPkcs7Cipher_Encrypt(t *testing.T) {
	successAes, _ := NewAesCbcPkcs7Cipher([]byte("testtesttesttest"))

	type args struct {
		plain []byte
	}
	tests := []struct {
		name    string
		c       *AesCbcPkcs7Cipher
		args    args
		wantErr bool
	}{
		{
			name: "Success Encrypt",
			c:    successAes,
			args: args{
				plain: []byte("test: this is a test text"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := tt.c.Encrypt(tt.args.plain)
			if err != nil {
				t.Errorf("AesCbcPkcs7Cipher.Encrypt() error = %v", err)
			}
			// if (err != nil) != tt.wantErr {
			// 	t.Errorf("AesCbcPkcs7Cipher.Encrypt() error = %v, wantErr %v", err, tt.wantErr)
			// 	return
			// }

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("AesCbcPkcs7Cipher.Encrypt() = %v, want %v", got, tt.want)
			// }
		})
	}
}

func TestAesCbcPkcs7Cipher_Decrypt(t *testing.T) {
	successAes, _ := NewAesCbcPkcs7Cipher([]byte("testtesttesttest"))
	plain := "test: this is test text"
	encrypt, _ := successAes.Encrypt([]byte(plain))

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
		{
			name: "Success Decrypt",
			c:    successAes,
			args: args{
				encryptedTextBase64: encrypt,
			},
			want:    plain,
			wantErr: false,
		},
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
