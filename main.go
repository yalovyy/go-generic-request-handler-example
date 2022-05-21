package main

import (
	"context"
	"fmt"
	"ghandler/environment/prod"
	"ghandler/router"
)

func main() {
	env := prod.NewProdEnvironment()
	r := router.New(env)

	requests := map[string]string{
		"GET /bar": `{"Id": "bar id"}`,
		"PUT /foo": `{"Id": "foo id", "Value": "foo value"}`,
	}

	ctx := context.Background()
	for path, request := range requests {
		response, err := r.Route(ctx, path, request)
		if err != nil {
			panic(err)
		}
		fmt.Println(response)
	}
}
