package gql

import (
	"fullstack-course/db"
	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *db.DB
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {

	name, ok := p.Args["name"].(string)

	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}

	return nil, nil
}
