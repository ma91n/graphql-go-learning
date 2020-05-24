package gq

import (
	"github.com/graphql-go/graphql"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user":      UserField,
		"userList":  UserListField,
		"event":     EventField,
		"eventList": EventListField,
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser":  CreateUserField,
		"createEvent": CreateEventField,
	},
})
