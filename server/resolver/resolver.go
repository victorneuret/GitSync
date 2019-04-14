package resolver

import (
	"github.com/victorneuret/GitSync/generated"
)

// THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct{}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &MutationResolverType{r}
}
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

type MutationResolverType struct{ *Resolver }

type queryResolver struct{ *Resolver }
