package schema

import (
	"github.com/graphql-go/graphql"
	"github.com/laqiiz/graphql-go-learning/suburi/schema/fields"
)

var RootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		"user":      fields.UserField,
		"userList":  fields.UserListField,
		"event":     fields.EventField,
		"eventList": fields.EventListField,
	},
})

var RootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser":  fields.CreateUserField,
		"createEvent": fields.CreateEventField,
	},
})
