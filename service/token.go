package service

import (
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/ndphu/token-service/domain"
	"os"
	"strconv"
	"time"
)

var (
	secret          = os.Getenv("TOKEN_SERVICE_SECRET")
	validTimeSecond = strconv.Atoi(os.Getenv("TOKEN_SERVICE_TTL")) // 2 hours
)

func CreateToken(input *domain.Token) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":        int64(time.Now().Add(time.Second * 7200).Unix()),
		"appId":      input.AppId,
		"id":         input.Id,
		"iss":        "Mightly Token Service",
		"sub":        input.Sub,
		"externalId": input.ExternalId,
		"firstName":  input.FirstName,
		"lastName":   input.LastName,
		"roles":      input.Roles,
	})

	return token.SignedString([]byte(secret))
}

func ValidateToken(tokenString string) (*domain.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("%v\n", claims)

		exp := int64(claims["exp"].(float64))
		if time.Now().Unix() > exp {
			return nil, errors.New("token is expired")
		}

		claimRoles := claims["roles"].([]interface{})
		roles := make([]string, len(claimRoles))
		for idx, role := range claimRoles {
			roles[idx] = role.(string)
		}

		tokenClaim := domain.Token{
			Exp:        exp,
			AppId:      claims["appId"].(string),
			Sub:        claims["sub"].(string),
			Iss:        claims["iss"].(string),
			Id:         claims["id"].(string),
			ExternalId: claims["externalId"].(string),
			FirstName:  claims["firstName"].(string),
			LastName:   claims["lastName"].(string),
			Roles:      roles,
		}
		return &tokenClaim, nil

	} else {
		return nil, errors.New("token validate failed")
	}
}
