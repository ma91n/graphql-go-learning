package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/laqiiz/graphql-go-learning/suburi/model"
	"github.com/laqiiz/graphql-go-learning/suburi/repository"
	"github.com/laqiiz/graphql-go-learning/suburi/schema/types"
)

var (
	ur = repository.NewUserRepository()
	er = repository.NewEventRepository()
)

var UserField = &graphql.Field{
	Type:        types.UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userId, isOK := params.Args["id"].(string)
		if isOK {
			return ur.FindById(userId)
		}
		return nil, errors.New("no userId")
	},
}

var UserListField = &graphql.Field{
	Type:        graphql.NewList(types.UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return ur.List(), nil
	},
}

var CreateUserField = &graphql.Field{
	Type:        types.UserType,
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
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userName, _ := params.Args["userName"].(string)
		description, _ := params.Args["description"].(string)
		photoURL, _ := params.Args["photoURL"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := model.NewUser(userName, description, photoURL, email)
		if err != nil {
			return nil, err
		}
		ur.Store(newUser)
		return newUser, nil
	},
}

var EventField = &graphql.Field{
	Type:        types.EventType,
	Description: "Get single event",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		eventId, isOK := params.Args["id"].(string)
		if isOK {
			return er.FindById(eventId)
		}
		return nil, errors.New("no eventId")
	},
}

var EventListField = &graphql.Field{
	Type:        graphql.NewList(types.EventType),
	Description: "List of events",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return er.List(), nil
	},
}

var CreateEventField = &graphql.Field{
	Type:        types.EventType,
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
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userId, _ := params.Args["userId"].(string)
		eventName, _ := params.Args["eventName"].(string)
		description, _ := params.Args["description"].(string)
		location, _ := params.Args["location"].(string)
		startTime, _ := params.Args["startTime"].(string)
		endTime, _ := params.Args["endTime"].(string)

		newEvent, err := model.NewEvent(userId, eventName, description, location, startTime, endTime)
		if err != nil {
			return nil, err
		}
		er.Store(newEvent)

		return newEvent, nil
	},
}
