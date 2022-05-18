package helper

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateJwtAccessToken(userId string, username string) (accessToken string, err error) {
	key, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_SECRET_KEY")))
	if err != nil {
		return accessToken, err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["user_id"] = userId
	claims["username"] = username
	claims["exp"] = now.AddDate(0, 0, 30).Unix() // The expiration time after which the token must be disregarded.
	claims["iat"] = now.Unix()                   // The time at which the token was issued.
	claims["iss"] = os.Getenv("JWT_ISS")
	claims["aud"] = os.Getenv("JWT_AUD")

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	accessToken, err = token.SignedString(key)
	if err != nil {
		return accessToken, err
	}

	return accessToken, nil
}
