package auth

import (
	"net/http"
	"reflect"
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
			name: "Secure Cookie?",
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
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateSecureCookie(tt.args.key, tt.args.val, tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateSecureCookie() = %v, want %v", got, tt.want)
			}
		})
	}
}
