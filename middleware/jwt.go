package middleware

import (
	"context"
	"eve/database"
	"eve/models"
	"eve/util"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const TOKEN_EXPIRATION = 15 * time.Minute

func GenerateToken(username string) (string, error) {

	claims := jwt.MapClaims{
		"username":   username,
		"Expiration": time.Now().Add(TOKEN_EXPIRATION).Unix(),
		"IssuedAt":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func DecodeToken(tokenString string) (string, error) {
	var err error

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		return string(claims["username"].(string)), nil
	}

	return "", err
}

func AuthMiddleware(handler http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" {
			util.RespondWithError(w, http.StatusUnauthorized, "No auth token in the header")
			return
		}

		tokenString := strings.Split(authHeader, " ")

		if len(tokenString) != 2 {
			util.RespondWithError(w, http.StatusUnauthorized, "Malformed token")
			return
		}

		username, err := DecodeToken(tokenString[1])
		if err != nil || username == "" {
			util.RespondWithError(w, http.StatusUnauthorized, "Bad token or token is expired")
			return
		}

		ctx := context.WithValue(r.Context(), "username", username)
		newReq := r.WithContext(ctx)

		handler.ServeHTTP(w, newReq)

	})
}

func RoleAuthorization(handler http.Handler, perm uint8) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username := r.Context().Value("username")
		var foundUser models.User

		fmt.Println("username: ", username)
		result := database.Database.Db.Where(models.User{Username: username.(string)}).First(&foundUser)

		if result.Error != nil {
			util.RespondWithError(w, http.StatusUnauthorized, "Unauthorized")
			return
		}

		var role models.Role

		result = database.Database.Db.First(&role, foundUser.RoleID)

		if !role.HasPermission(perm) {
			util.RespondWithError(w, http.StatusForbidden, "Forbidden")
			return
		}

		handler.ServeHTTP(w, r)
	})
}
