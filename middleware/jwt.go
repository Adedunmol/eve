package middleware

import (
	"eve/util"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username":   username,
		"Expiration": 15 * time.Minute,
		"IssuedAt":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(os.Getenv("SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeToken(token string) bool {

	return true
}

func AuthMiddleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")
		token := strings.TrimPrefix(authHeader, "Bearer ")

		if authHeader == "" || authHeader == token {
			util.RespondWithError(w, http.StatusUnauthorized, "No auth token in the header")
			return
		}

		isValid := DecodeToken(token)

		handler.ServeHTTP(w, r)

		log.Println(isValid)
	})
}
