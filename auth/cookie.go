package auth

import (
	"net/http"
	"time"
)

func CreateSecureCookie(key, val string, t time.Duration) *http.Cookie {
	ck := new(http.Cookie)
	ck.Name = key
	ck.Value = val
	ck.Secure = true
	ck.HttpOnly = true
	ck.SameSite = http.SameSiteNoneMode
	ck.Expires = time.Now().Add(t * time.Hour)
	return ck
}
