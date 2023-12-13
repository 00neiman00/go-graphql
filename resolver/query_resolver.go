package resolver

import (
	"context"
	"github.com/neimen-95/go-graphql/graphql"
	"github.com/neimen-95/go-graphql/models"
)

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.MeetupRepository.GetMeetups()
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() graphql.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
