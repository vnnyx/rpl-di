package helper

import (
	"errors"
	"fmt"
	"os"
	"rpl-sixmath/exception"
	"rpl-sixmath/model"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
)

func CreateToken(request model.JwtPayload) *model.TokenDetails {
	accessExpired, _ := strconv.Atoi(os.Getenv("JWT_MINUTE"))

	td := &model.TokenDetails{}

	td.AtExpires = time.Now().Add(time.Minute * time.Duration(accessExpired)).Unix()
	td.AccessUuid = uuid.NewV4().String()

	var err error

	keyAccess, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(os.Getenv("JWT_SECRET_KEY")))
	exception.PanicIfNeeded(err)

	now := time.Now().UTC()

	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = request.UserID
	atClaims["username"] = request.Username
	atClaims["email"] = request.Email
	atClaims["avatar"] = request.Avatar
	atClaims["role"] = request.Role
	atClaims["access_uuid"] = td.AccessUuid
	atClaims["exp"] = td.AtExpires
	atClaims["iat"] = now.Unix()
	atClaims["iss"] = "six-math"
	atClaims["aud"] = "six-math"

	at := jwt.NewWithClaims(jwt.SigningMethodRS256, atClaims)
	at.Header["six-math"] = "jwt"
	td.AccessToken, err = at.SignedString(keyAccess)

	if err != nil {
		fmt.Println(err)
		exception.PanicIfNeeded(errors.New(model.AUTHENTICATION_FAILURE_ERR_TYPE))
	}

	return td

}
