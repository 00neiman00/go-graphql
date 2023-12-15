package dataloader

import (
	"context"
	"net/http"
	"time"

	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/models"
)

const userLoaderKey = "userLoader"

func DataloaderMiddleWare(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userLoader := UserLoader{
			fetch: func(ids []string) ([]*models.User, []error) {
				var users []*models.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					return nil, []error{err}
				}
				u := make(map[string]*models.User, len(users))

				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*models.User, len(ids))
				for i, id := range ids {
					result[i] = u[id]
				}

				return result, []error{err}
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
