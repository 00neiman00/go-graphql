package resolver

import (
	"context"
	"errors"
	"fmt"

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

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input graphql.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := r.MeetupRepository.GetById(id)
	if err != nil {
		return nil, errors.New("meetup not found")
	}

	if input.Name != nil {
		meetup.Name = *input.Name
	}

	if input.Description != nil {
		meetup.Description = *input.Description
	}

	meetup, err = r.MeetupRepository.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error updating meetup: %v", err)
	}
	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupRepository.GetById(id)
	if err != nil {
		return false, errors.New("meetup not found")
	}

	err = r.MeetupRepository.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error deleting meetup: %v", err)
	}
	return true, nil
}
