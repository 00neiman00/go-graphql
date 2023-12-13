package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/models"
)

type UserRepository struct {
	DB *pg.DB
}

func (u *UserRepository) FindById(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}

	return &user, nil
}
