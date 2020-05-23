package repository

import (
	"errors"
	"github.com/laqiiz/graphql-go-learning/suburi/model"
)

type EventRepository struct {
	events []*model.Event
}

func NewEventRepository() *EventRepository {
	return &EventRepository{}
}

// store event to repository
func (r *EventRepository) Store(event *model.Event) {
	r.events = append(r.events, event)
}

func (r EventRepository) FindById(userId string) (*model.Event, error) {
	for _, val := range r.events {
		if val.EventId == userId {
			return val, nil
		}
	}
	return nil, errors.New("user not found")
}

func (r EventRepository) List() []*model.Event {
	return r.events
}
