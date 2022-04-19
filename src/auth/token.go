package auth

import (
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

// custom claims
type Claims struct {
	User         int64 `json:"user_id"`
	IsGoogleUser bool  `json:"is_google_user"`
	jwt.StandardClaims
}

var jwt_token = os.Getenv("jwt_token")

// jwt secret key
var jwtSecret = []byte(jwt_token)

// generate JWT
func SetClaim(user_id int64, is_google_user bool) (string, error) {
	now := time.Now()
	jwtId := strconv.FormatInt(user_id, 10) + strconv.FormatInt(now.Unix(), 10)
	// set claims and sign
	claims := Claims{
		User:         user_id,
		IsGoogleUser: is_google_user,
		StandardClaims: jwt.StandardClaims{
			Audience:  strconv.FormatInt(user_id, 10),
			ExpiresAt: now.Add(24 * time.Hour).Unix(),
			Id:        jwtId,
			IssuedAt:  now.Unix(),
			Issuer:    "ginJWT",
			NotBefore: now.Add(time.Second).Unix(),
			Subject:   strconv.FormatInt(user_id, 10),
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return token, nil
}

// validate JWT
func AuthRequired(c *gin.Context) {
	auth := c.GetHeader("Authorization")
	token := strings.Split(auth, "Bearer ")[1]

	// parse and validate token for six things:
	// validationErrorMalformed => token is malformed
	// validationErrorUnverifiable => token could not be verified because of signing problems
	// validationErrorSignatureInvalid => signature validation failed
	// validationErrorExpired => exp validation failed
	// validationErrorNotValidYet => nbf validation failed
	// validationErrorIssuedAt => iat validation failed
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (i interface{}, err error) {
		return jwtSecret, nil
	})

	if err != nil {
		var message string
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				message = "token is malformed"
			} else if ve.Errors&jwt.ValidationErrorUnverifiable != 0 {
				message = "token could not be verified because of signing problems"
			} else if ve.Errors&jwt.ValidationErrorSignatureInvalid != 0 {
				message = "signature validation failed"
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				message = "token is expired"
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				message = "token is not yet valid before sometime"
			} else {
				message = "can not handle this token"
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": message,
		})
		c.Abort()
		return
	}

	if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
		c.Set("is_google_user", claims.IsGoogleUser)
		c.Set("user_id", claims.User)
		c.Next()
	} else {
		c.Abort()
		return
	}
}
