package auth

import (
	"reflect"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func TestTmpToken_IsVerificationLoginToken(t *testing.T) {
	tests := []struct {
		name string
		tr   *TmpToken
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.IsVerificationLoginToken(); got != tt.want {
				t.Errorf("TmpToken.IsVerificationLoginToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTmpToken_IsChangesEmailToken(t *testing.T) {
	tests := []struct {
		name string
		tr   *TmpToken
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.IsChangesEmailToken(); got != tt.want {
				t.Errorf("TmpToken.IsChangesEmailToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTmpToken_TokenTypeAsStr(t *testing.T) {
	tests := []struct {
		name string
		tr   *TmpToken
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.TokenTypeAsStr(); got != tt.want {
				t.Errorf("TmpToken.TokenTypeAsStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateTokens(t *testing.T) {
	type args struct {
		id interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := CreateTokens(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTokens() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTokens() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateTokens() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestCreateTempURLAndToken(t *testing.T) {
	type args struct {
		id        interface{}
		tmpUrl    string
		tokenType int
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := CreateTempURLAndToken(tt.args.id, tt.args.tmpUrl, tt.args.tokenType)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateTempURLAndToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CreateTempURLAndToken() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("CreateTempURLAndToken() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseTmpToken(t *testing.T) {
	type args struct {
		token string
	}
	tests := []struct {
		name string
		args args
		want *TmpToken
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseTmpToken(tt.args.token); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseTmpToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseAccessToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAccessToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAccessToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseAccessToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseRefreshToken(t *testing.T) {
	type args struct {
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseRefreshToken(tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseRefreshToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_verifyTokenString(t *testing.T) {
	type args struct {
		tokenString string
		secret      string
	}
	tests := []struct {
		name    string
		args    args
		want    *jwt.Token
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := verifyTokenString(tt.args.tokenString, tt.args.secret)
			if (err != nil) != tt.wantErr {
				t.Errorf("verifyTokenString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("verifyTokenString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseToken(t *testing.T) {
	type args struct {
		token *jwt.Token
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseToken(tt.args.token)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
