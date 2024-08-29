package encryption

import (
	"github.com/MauricioGZ/CRUD-GO/internal/model"
	jwt "github.com/golang-jwt/jwt/v5"
)

func SignedLoginToken(u *model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": u.Email,
		"role":  u.Role,
	})
	return token.SignedString([]byte(key))
}

func ParseLoginJWT(_token string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(_token, func(t *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})

	if err != nil {
		return nil, err
	}

	return token.Claims.(jwt.MapClaims), nil
}
