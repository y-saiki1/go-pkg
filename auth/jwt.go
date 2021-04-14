package auth

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	TokenTypeLogin = iota + 1
	TokenTypeRegister
	TokenTypeChangeEmail
)

type TmpToken struct {
	ID   string
	Type int
	Err  error
}

func (t *TmpToken) IsVerificationLoginToken() bool {
	return t.Type == TokenTypeLogin
}
func (t *TmpToken) IsChangesEmailToken() bool {
	return t.Type == TokenTypeChangeEmail
}
func (t *TmpToken) TokenTypeAsStr() string {
	switch t.Type {
	case TokenTypeLogin:
		return "ログイン"
	case TokenTypeRegister:
		return "新規登録"
	case TokenTypeChangeEmail:
		return "メールアドレス変更"
	}
	return ""
}

func CreateTokens(id interface{}) (string, string, error) {
	acTk := jwt.New(jwt.SigningMethodHS256)
	acClaims := acTk.Claims.(jwt.MapClaims)
	acClaims["id"] = id
	acClaims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	acToken, err := acTk.SignedString([]byte(os.Getenv("ACCESS_JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	// TODO: こっちはAccessTokenと紐付ける形にしてあげないといけない
	refTk := jwt.New(jwt.SigningMethodHS256)
	refClaims := refTk.Claims.(jwt.MapClaims)
	refClaims["id"] = id
	refClaims["exp"] = time.Now().Add(time.Hour * 160).Unix()
	refToken, err := refTk.SignedString([]byte(os.Getenv("REFRESH_JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	return acToken, refToken, nil
}

func CreateTempURLAndToken(id interface{}, tmpUrl string, tokenType int) (string, string, error) {
	tmTk := jwt.New(jwt.SigningMethodHS256)
	tmClaims := tmTk.Claims.(jwt.MapClaims)
	tmClaims["id"] = id
	tmClaims["type"] = tokenType
	tmClaims["exp"] = time.Now().Add(time.Minute * 10).Unix()
	tk, err := tmTk.SignedString([]byte(os.Getenv("TEMP_URL_JWT_SECRET")))
	if err != nil {
		return "", "", err
	}

	return tmpUrl + "?token=" + tk, tk, nil
}

func ParseTmpToken(token string) *TmpToken {
	tk, err := verifyTokenString(token, os.Getenv("TEMP_URL_JWT_SECRET"))
	if err != nil {
		return &TmpToken{"", 0, err}
	}
	claims := tk.Claims.(jwt.MapClaims)

	return &TmpToken{
		claims["id"].(string),
		int(claims["type"].(float64)),
		nil,
	}
}

func ParseAccessToken(tokenString string) (string, error) {
	token, err := verifyTokenString(tokenString, os.Getenv("ACCESS_JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return parseToken(token)
}

func ParseRefreshToken(tokenString string) (string, error) {
	token, err := verifyTokenString(tokenString, os.Getenv("REFRESH_JWT_SECRET"))
	if err != nil {
		return "", err
	}

	return parseToken(token)
}

func verifyTokenString(tokenString string, secret string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
}

func parseToken(token *jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return "", fmt.Errorf("無効なトークンが検出されました")
	}

	return claims["id"].(string), nil
}
