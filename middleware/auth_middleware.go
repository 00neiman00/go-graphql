package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/neimen-95/go-graphql/models"
	"github.com/neimen-95/go-graphql/postgres"
	"github.com/pkg/errors"
	"net/http"
	"os"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(repo *postgres.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)

			if !ok || !token.Valid {
				next.ServeHTTP(w, r)
				return
			}

			user, err := repo.FindById(claims["jti"].(string))

			if err != nil {
				next.ServeHTTP(w, r)
				return
			}

			ctx := context.WithValue(r.Context(), CurrentUserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func stripBearerPrefixFromToken(token string) (string, error) {
	if len(token) > 6 && token[:7] == "Bearer " {
		return token[7:], nil
	}
	return token, nil
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	token, err := request.ParseFromRequest(r, authExtractor, getJwtSecret)

	return token, errors.Wrap(err, "parseToken error: ")
}

func getJwtSecret(token *jwt.Token) (interface{}, error) {
	t := []byte(os.Getenv("JWT_SECRET"))
	return t, nil
}

func GetCurrentUserFromContext(ctx context.Context) (*models.User, error) {
	if ctx.Value(CurrentUserKey) == nil {
		return nil, errors.New("no user in context")
	}

	user, ok := ctx.Value(CurrentUserKey).(*models.User)
	if !ok {
		return nil, errors.New("user is of invalid type")
	}

	return user, nil
}
