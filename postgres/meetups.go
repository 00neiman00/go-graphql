package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/neimen-95/go-graphql/models"
)

type MeetupRepository struct {
	DB *pg.DB
}

func (m *MeetupRepository) GetMeetups(filter *models.MeetupFilter, limit, offset *int) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	query := m.DB.Model(&meetups)
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}
	if limit != nil {
		query.Limit(*limit)
	}
	if offset != nil {
		query.Offset(*offset)
	}
	err := query.Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupRepository) GetById(id string) (*models.Meetup, error) {
	var meetup models.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &meetup, nil
}

func (m *MeetupRepository) GetByUserId(userId string) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", userId).Select()
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

func (m *MeetupRepository) Update(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

func (m *MeetupRepository) Delete(meetup *models.Meetup) error {
	_, err := m.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}
