package resolver

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

import (
	"github.com/neimen-95/go-graphql/postgres"
)

type Resolver struct {
	MeetupRepository *postgres.MeetupRepository
	UserRepository   *postgres.UserRepository
}
