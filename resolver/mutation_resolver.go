package resolver

import (
	"context"
	"errors"
	"github.com/neimen-95/go-graphql/graphql"
	"github.com/neimen-95/go-graphql/models"
)

type mutationResolver struct{ *Resolver }

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() graphql.MutationResolver { return &mutationResolver{r} }

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, newMeetup graphql.NewMeetup) (*models.Meetup, error) {
	if len(newMeetup.Name) < 3 {
		return nil, errors.New("name must be at least 3 characters")
	}

	if len(newMeetup.Description) < 3 {
		return nil, errors.New("description must be at least 3 characters")
	}

	meetup := &models.Meetup{
		Name:        newMeetup.Name,
		Description: newMeetup.Description,
		UserId:      "1",
	}
	return r.MeetupRepository.CreateMeetup(meetup)
}
