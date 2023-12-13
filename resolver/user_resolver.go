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

func (u *userResolver) Meetup(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var userMeetups []*models.Meetup
	return userMeetups, nil
}
