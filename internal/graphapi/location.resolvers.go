package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.31-dev

import (
	"context"

	"entgo.io/contrib/entgql"
	"go.infratographer.com/location-api/internal/ent/generated"
	"go.infratographer.com/location-api/internal/ent/generated/location"
	"go.infratographer.com/x/gidx"
)

// Owner is the resolver for the owner field.
func (r *locationResolver) Owner(ctx context.Context, obj *generated.Location) (*ResourceOwner, error) {
	return &ResourceOwner{ID: obj.OwnerID}, nil
}

// LocationCreate is the resolver for the locationCreate field.
func (r *mutationResolver) LocationCreate(ctx context.Context, input generated.CreateLocationInput) (*LocationCreatePayload, error) {
	// TODO: auth check
	loc, err := r.client.Location.Create().SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &LocationCreatePayload{Location: loc}, nil
}

// LocationDelete is the resolver for the locationDelete field.
func (r *mutationResolver) LocationDelete(ctx context.Context, id gidx.PrefixedID) (*LocationDeletePayload, error) {
	// TODO: auth check
	// TODO: check metadata for references to this location
	if err := r.client.Location.DeleteOneID(id).Exec(ctx); err != nil {
		return nil, err
	}

	return &LocationDeletePayload{DeletedID: id}, nil
}

// LocationUpdate is the resolver for the locationUpdate field.
func (r *mutationResolver) LocationUpdate(ctx context.Context, id gidx.PrefixedID, input generated.UpdateLocationInput) (*LocationUpdatePayload, error) {
	// TODO: auth check
	loc, err := r.client.Location.UpdateOneID(id).SetInput(input).Save(ctx)
	if err != nil {
		return nil, err
	}

	return &LocationUpdatePayload{Location: loc}, nil
}

// Location is the resolver for the location field.
func (r *queryResolver) Location(ctx context.Context, id gidx.PrefixedID) (*generated.Location, error) {
	// TODO: auth check
	return r.client.Location.Get(ctx, id)
}

// Locations is the resolver for the locations field.
func (r *resourceOwnerResolver) Locations(ctx context.Context, obj *ResourceOwner, after *entgql.Cursor[gidx.PrefixedID], first *int, before *entgql.Cursor[gidx.PrefixedID], last *int, orderBy *generated.LocationOrder, where *generated.LocationWhereInput) (*generated.LocationConnection, error) {
	return r.client.Location.Query().Where(location.OwnerID(obj.ID)).Paginate(ctx, after, first, before, last, generated.WithLocationOrder(orderBy), generated.WithLocationFilter(where.Filter))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// ResourceOwner returns ResourceOwnerResolver implementation.
func (r *Resolver) ResourceOwner() ResourceOwnerResolver { return &resourceOwnerResolver{r} }

type (
	mutationResolver      struct{ *Resolver }
	resourceOwnerResolver struct{ *Resolver }
)