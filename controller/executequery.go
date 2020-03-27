package controller

import (
	"context"
	resolver "eaciit/gopher/resolver"
	schemas "eaciit/gopher/schemas"
	types "eaciit/gopher/types"
	"github.com/dgrijalva/jwt-go"
	graphql "github.com/graph-gophers/graphql-go"
	"net/http"
	"os"
)

var (
	// We can pass an option to the schema so we don’t need to
	// write a method to access each type’s field:
	opts = []graphql.SchemaOpt{graphql.UseFieldResolvers()}
	// Schema get schemas

)

// ExecuteQuery djfsdjf
func ExecuteQuery(query string, request *http.Request) *graphql.Response {
	ctx := context.Background()
	q1 := types.ClientQuery{
		Query:     query,
		Variables: nil,
	}
	tokenClaim := types.TokenClaim{}
	headertoken := request.Header.Get("Token-Key")

	if headertoken != "" {
		token, _ := jwt.ParseWithClaims(headertoken, &types.TokenClaim{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SIGNING_KEY")), nil
		})
		if token != nil && token.Valid {
			claims := token.Claims.(*types.TokenClaim)
			tokenClaim.Email = claims.Email
		}

	}

	Schema := graphql.MustParseSchema(schemas.Schema, &resolver.RootResolver{
		Session: tokenClaim.Email,
	}, opts...)
	resp1 := Schema.Exec(ctx, q1.Query, "", q1.Variables)

	return resp1
}
