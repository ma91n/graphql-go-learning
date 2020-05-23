package fields

import (
	"encoding/json"
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/laqiiz/graphql-go-learning/suburi/model"
	"github.com/laqiiz/graphql-go-learning/suburi/repository"
	"github.com/laqiiz/graphql-go-learning/suburi/schema/types"
	"os/user"
)

// fetch single user
var UserField = &graphql.Field{
	Type:        types.UserType, // 返り値の型
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{ //引数の定義
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) { //実行関数
		userId, isOK := params.Args["id"].(string) // 引数取り出し
		if isOK {
			return repository.NewUserRepository().FindById(userId)
		}

		return nil, errors.New("no userId")
	},
}

// fetch all user
var UserListField = &graphql.Field{
	Type:        graphql.NewList(types.UserType),
	Description: "List of users",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.NewUserRepository().List(), nil
	},
}

// create user
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
			panic(err)
		}
		repository.NewUserRepository().Store(newUser)
		return newUser, nil
	},
}
// fetch single event
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
			return repository.NewEventRepository().FindById(eventId)
		}
		return nil, errors.New("no eventId")
	},
}

// fetch all event
var EventListField = &graphql.Field{
	Type:        graphql.NewList(types.EventType),
	Description: "List of events",
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		return repository.NewEventRepository().List(), nil
	},
}

// create event
var CreateEventField = &graphql.Field{
	Type:        types.EventType,
	Description: "Create new event",
	Args: graphql.FieldConfigArgument{
		"user": &graphql.ArgumentConfig{
			Type: types.UserInput,
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
		givenUser, _ := params.Args["user"]
		eventName, _ := params.Args["eventName"].(string)
		description, _ := params.Args["description"].(string)
		location, _ := params.Args["location"].(string)
		startTime, _ := params.Args["startTime"].(string)
		endTime, _ := params.Args["endTime"].(string)

		parsedJson, err := json.Marshal(givenUser)
		if err != nil {
			panic(err)
		}
		var parsedUser user.User
		err = json.Unmarshal(parsedJson, &parsedUser)
		if err != nil {
			panic(err)
		}

		// create new Event
		newEvent, err := model.NewEvent(&parsedUser, eventName, description, location, startTime, endTime)
		if err != nil {
			panic(err)
		}

		repository.NewEventRepository().Store(newEvent)
		return newEvent, nil
	},
}

