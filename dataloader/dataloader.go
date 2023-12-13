package dataloader

import (
	"context"
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/models"
	"net/http"
	"time"
)

const userLoaderKey = "userLoader"

func DataloaderMiddleWare(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User
				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()
				return users, []error{err}
			},
			wait:     1 * time.Millisecond,
			maxBatch: 100,
		}

		ctx := context.WithValue(r.Context(), userLoaderKey, &userLoader)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(userLoaderKey).(*UserLoader)
}
