package jwt

import (
	"anla.io/taizhou-fe-api/config"
	"anla.io/taizhou-fe-api/handler"
	jwt "github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
)

var (
	jwtConf = config.Config.JWT
)

// JwtHandler is
var JwtHandler = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.Secret), nil
	},
	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	SigningMethod: jwt.SigningMethodHS256,
	ErrorHandler:  handler.JWTError,
})

// JwtHandlerAdmin is
var JwtHandlerAdmin = jwtmiddleware.New(jwtmiddleware.Config{
	ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtConf.AdminSecret), nil
	},
	// When set, the middleware verifies that tokens are signed with the specific signing algorithm
	// If the signing method is not constant the ValidationKeyGetter callback can be used to implement additional checks
	// Important to avoid security issues described here: https://auth0.com/blog/2015/03/31/critical-vulnerabilities-in-json-web-token-libraries/
	SigningMethod: jwt.SigningMethodHS256,
	ErrorHandler:  handler.JWTError,
})
