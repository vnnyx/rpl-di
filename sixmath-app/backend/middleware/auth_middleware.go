package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"rpl-sixmath/model"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

type DecodedStructure struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func ValidateToken(encodedToken string) (token *jwt.Token, errData error) {
	jwtPublicKey, err := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))

	if err != nil {
		return token, err
	}

	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err = jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return token, err
	}
	if !token.Valid {
		return token, errors.New("invalid token")
	}
	return token, nil
}

func DecodeToken(encodedToken string) (decodedResult DecodedStructure, errData error) {
	jwtPublicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(os.Getenv("JWT_PUBLIC_KEY")))
	tokenString := encodedToken
	claims := jwt.MapClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtPublicKey, nil
	})
	if err != nil {
		return decodedResult, err
	}
	if !token.Valid {
		return decodedResult, errors.New("invalid token")
	}

	jsonbody, err := json.Marshal(claims)
	if err != nil {
		return decodedResult, err
	}

	var obj DecodedStructure
	if err := json.Unmarshal(jsonbody, &obj); err != nil {
		return decodedResult, err
	}

	return obj, nil
}
func CheckToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		tokenSlice := strings.Split(c.Get("Authorization"), "Bearer ")

		var tokenString string
		if len(tokenSlice) == 2 {
			tokenString = tokenSlice[1]
		}

		// validate token
		_, err := ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err)
			response := model.Response{
				Code:   401,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// extract data from token
		decodedRes, err := DecodeToken(tokenString)
		if err != nil {
			fmt.Println(err)
			response := model.Response{
				Code:   401,
				Status: "Unauthorized",
				Error: map[string]interface{}{
					"general": "UNAUTHORIZED",
				},
			}
			return c.Status(http.StatusUnauthorized).JSON(response)
		}

		// set to global var
		c.Locals("currentUserId", decodedRes.UserId)
		c.Locals("currentUsername", decodedRes.Username)
		c.Locals("currentRole", decodedRes.Role)

		return c.Next()
	}
}
