package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/models"
)

type UserRepository struct {
	DB *pg.DB
}

func (u *UserRepository) GetByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

func (u *UserRepository) GetById(id string) (*models.User, error) {
	return u.GetByField("id", id)
}

func (u *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	return u.GetByField("email", email)
}

func (u *UserRepository) GetUserByName(username string) (*models.User, error) {
	return u.GetByField("username", username)
}

func (u *UserRepository) Save(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return user, nil
}
