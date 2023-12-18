package resolver

import (
	"context"
	"github.com/neimen-95/go-graphql/graphql"
	"github.com/neimen-95/go-graphql/models"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() graphql.UserResolver {
	return &userResolver{r}
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepository.FindById(id)
}

func (u *userResolver) Meetup(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return u.MeetupRepository.GetByUserId(obj.ID)
}
