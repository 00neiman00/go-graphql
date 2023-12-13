package resolver

import (
	"context"
	"github.com/neimen-95/go-graphql/dataloader"
	"github.com/neimen-95/go-graphql/graphql"
	"github.com/neimen-95/go-graphql/models"
)

type meetupResolver struct{ *Resolver }

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return dataloader.GetUserLoader(ctx).Load(obj.UserId)
}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() graphql.MeetupResolver { return &meetupResolver{r} }
