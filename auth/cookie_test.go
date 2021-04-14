package auth

import (
	"net/http"
	"testing"
	"time"
)

func TestCreateSecureCookie(t *testing.T) {
	type args struct {
		key string
		val string
		t   time.Duration
	}
	tests := []struct {
		name string
		args args
		want *http.Cookie
	}{
		{
			name: "Success Create Secure Token",
			args: args{
				key: "at",
				val: "test",
				t:   1,
			},
			want: &http.Cookie{
				Name:     "at",
				Value:    "test",
				Expires:  time.Now().Add(1 * time.Hour),
				Secure:   true,
				HttpOnly: true,
				SameSite: http.SameSiteLaxMode,
			},
		},
	}
	// Cookieは構造体の等価比較では同一と判断できないので、URL文字列で比較する
	// https://stackoverflow.com/questions/46246728/testing-cookie-returned-from-function-in-go
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := CreateSecureCookie(tt.args.key, tt.args.val, tt.args.t); !reflect.DeepEqual(got, tt.want) {
			got := CreateSecureCookie(tt.args.key, tt.args.val, tt.args.t)
			if got.String() != tt.want.String() {
				t.Errorf("got  %v", got.String())
				t.Errorf("want %v", tt.want.String())
			}
		})
	}
}
