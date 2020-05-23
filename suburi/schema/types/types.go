package types

import "github.com/graphql-go/graphql"

var EventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Event",
	Fields: graphql.Fields{
		"eventId": &graphql.Field{
			Type: graphql.String,
		},
		"user": &graphql.Field{
			Type: UserType,
		},
		"eventName": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"location": &graphql.Field{
			Type: graphql.String,
		},
		"startTime": &graphql.Field{
			Type: graphql.String,
		},
		"endTime": &graphql.Field{
			Type: graphql.String,
		},
		"participants": &graphql.Field{
			Type: graphql.NewList(UserType),
		},
		"registeredTime": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"userId": &graphql.Field{
			Type: graphql.String,
		},
		"userName": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"photoURL": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})
