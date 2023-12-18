package resolver

import (
	"context"
	"errors"
	"github.com/neimen-95/go-graphql/graphql"
	"github.com/neimen-95/go-graphql/models"
	"log"
)

type userResolver struct{ *Resolver }

func (r *Resolver) User() graphql.UserResolver {
	return &userResolver{r}
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {
	return r.UserRepository.GetById(id)
}

func (u *userResolver) Meetup(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return u.MeetupRepository.GetByUserId(obj.ID)
}

func (m *mutationResolver) Register(ctx context.Context, input *models.RegisterInput) (*models.AuthResponse, error) {
	_, err := m.UserRepository.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in used")
	}

	_, err = m.UserRepository.GetUserByName(input.Username)
	if err == nil {
		return nil, errors.New("username already in used")
	}

	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hasing password: %v", err)
		return nil, errors.New("something went wrong")
	}

	// TODO: create verification code
	tx, err := m.UserRepository.DB.Begin()
	if err != nil {
		log.Printf("error createing a transaction: %v", err)
		return nil, errors.New("something went wrong")
	}

	defer tx.Rollback()
	if _, err := m.UserRepository.Save(tx, user); err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating the token: %v", err)
		return nil, errors.New("something went wrong")
	}
	authResponse := &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, err
	}

	return authResponse, nil
}
