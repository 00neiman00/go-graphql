package postgres

import (
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/models"
)

type MeetupRepository struct {
	DB *pg.DB
}

func (m *MeetupRepository) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupRepository) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()
	if err != nil {
		return nil, err
	}
	return meetup, nil
}
