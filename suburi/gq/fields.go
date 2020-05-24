package gq

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/laqiiz/graphql-go-learning/suburi/dao"
	"github.com/laqiiz/graphql-go-learning/suburi/model"
)

var (
	ur = dao.NewUserRepository()
	er = dao.NewEventRepository()
)

var UserField = &graphql.Field{
	Type:        UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userId, ok := p.Args["id"].(string)
		if ok {
			return ur.FindById(userId)
		}
		return nil, errors.New("no userId")
	},
}

var UserListField = &graphql.Field{
	Type:        graphql.NewList(UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return ur.List(), nil
	},
}

var CreateUserField = &graphql.Field{
	Type:        UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"userName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"photoURL": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userName, _ := p.Args["userName"].(string)
		description, _ := p.Args["description"].(string)
		photoURL, _ := p.Args["photoURL"].(string)
		email, _ := p.Args["email"].(string)

		newUser, err := model.NewUser(userName, description, photoURL, email)
		if err != nil {
			return nil, err
		}
		ur.Store(newUser)
		return newUser, nil
	},
}

var EventField = &graphql.Field{
	Type:        EventType,
	Description: "Get single event",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		eventId, ok := p.Args["id"].(string)
		if ok {
			return er.FindById(eventId)
		}
		return nil, errors.New("no eventId")
	},
}

var EventListField = &graphql.Field{
	Type:        graphql.NewList(EventType),
	Description: "List of events",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return er.List(), nil
	},
}

var CreateEventField = &graphql.Field{
	Type:        EventType,
	Description: "Create new event",
	Args: graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"eventName": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"location": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"startTime": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"endTime": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		userId, _ := p.Args["userId"].(string)
		eventName, _ := p.Args["eventName"].(string)
		description, _ := p.Args["description"].(string)
		location, _ := p.Args["location"].(string)
		startTime, _ := p.Args["startTime"].(string)
		endTime, _ := p.Args["endTime"].(string)

		newEvent, err := model.NewEvent(userId, eventName, description, location, startTime, endTime)
		if err != nil {
			return nil, err
		}
		er.Store(newEvent)

		return newEvent, nil
	},
}
