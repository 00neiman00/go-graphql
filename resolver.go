package main

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"errors"
	"github.com/neimen-95/go-graphql/postgres"

	"github.com/neimen-95/go-graphql/models"
)

type Resolver struct {
	meetupRepository *postgres.MeetupRepository
	userRepository   *postgres.UserRepository
}

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user, err := r.userRepository.FindById(obj.UserId)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, newMeetup NewMeetup) (*models.Meetup, error) {
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
	return r.meetupRepository.CreateMeetup(meetup)
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return r.meetupRepository.GetMeetups()
}

type userResolver struct{ *Resolver }

func (u *userResolver) Meetup(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var userMeetups []*models.Meetup
	return userMeetups, nil
}

// Meetup returns MeetupResolver implementation.
func (r *Resolver) Meetup() MeetupResolver { return &meetupResolver{r} }

func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
