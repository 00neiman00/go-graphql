package main

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"context"
	"errors"

	"github.com/neimen-95/go-graphql/models"
)

var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "GraphQL Meetup",
		Description: "First meetup",
		UserId:      "1",
	},
	{
		ID:          "2",
		Name:        "Go Meetup",
		Description: "Second meetup",
		UserId:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "neimen",
		Email:    "neimen@gmail.com",
	},
	{
		ID:       "2",
		Username: "john doe",
		Email:    "john@gmail.com",
	},
}

type Resolver struct{}

// User is the resolver for the user field.
func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)
	for _, u := range users {
		if u.ID == obj.UserId {
			user = u
			break
		}
	}

	if user != nil {
		return user, nil
	}

	return nil, errors.New("user not found")
}

// CreateMeetup is the resolver for the createMeetup field.
func (r *mutationResolver) CreateMeetup(ctx context.Context, newMeetup NewMeetup) (*models.Meetup, error) {
	panic("not implemented")
}

// Meetups is the resolver for the meetups field.
func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
}

type userResolver struct{ *Resolver }

func (u *userResolver) Meetup(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var userMeetups []*models.Meetup
	for _, m := range meetups {
		if m.UserId == obj.ID {
			userMeetups = append(userMeetups, m)
		}
	}

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
