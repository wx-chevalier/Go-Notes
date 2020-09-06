package middleware

import (
	"context"
	"log"
	"net/http"

	"github.com/L04DB4L4NC3R/jobs-mhrd/api/views"
	pkg "github.com/L04DB4L4NC3R/jobs-mhrd/pkg"
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/spf13/viper"
)

func Validate(h http.Handler) http.Handler {
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(viper.GetString("jwt_secret")), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	return jwtMiddleware.Handler(h)
}

func ValidateAndGetClaims(ctx context.Context, role string) (map[string]interface{}, error) {

	token, ok := ctx.Value("user").(*jwt.Token)
	if !ok {
		log.Println(token)
		return nil, views.ErrInvalidToken
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok {
		log.Println(claims)
		return nil, views.ErrInvalidToken
	}

	if claims.Valid() != nil {
		return nil, views.ErrInvalidToken
	}

	if claims["role"].(string) != role {
		log.Println(claims["role"])
		return nil, pkg.ErrUnauthorized
	}
	return claims, nil
}
