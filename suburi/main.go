package main

import (
	"encoding/json"
	"fmt"
	"github.com/graphql-go/graphql"
	"github.com/laqiiz/graphql-go-learning/suburi/schema"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	gs, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    schema.RootQuery,
		Mutation: schema.RootMutation,
	})
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		defer r.Body.Close()

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			json.NewEncoder(w).Encode(map[string]interface{}{"message": err})
		}
		fmt.Println(string(body))

		resp := graphql.Do(graphql.Params{
			Schema:        gs,
			RequestString: string(body),
		})
		if len(resp.Errors) > 0 {
			fmt.Printf("wrong result, unexpected errors: %v", resp.Errors)
		}
		json.NewEncoder(w).Encode(resp)
	})

	fmt.Println("Now server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
