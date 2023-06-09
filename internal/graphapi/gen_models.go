// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package graphapi

import (
	"go.infratographer.com/location-api/internal/ent/generated"
	"go.infratographer.com/x/gidx"
)

// Return response from locationCreate
type LocationCreatePayload struct {
	// The created annotation namespace.
	Location *generated.Location `json:"location"`
}

// Return response from locationDelete
type LocationDeletePayload struct {
	// The ID of the deleted annotation namespace.
	DeletedID gidx.PrefixedID `json:"deletedID"`
}

// Return response from locationUpdate
type LocationUpdatePayload struct {
	// The updated annotation namespace.
	Location *generated.Location `json:"location"`
}

type ResourceOwner struct {
	ID        gidx.PrefixedID               `json:"id"`
	Locations *generated.LocationConnection `json:"locations"`
}

func (ResourceOwner) IsEntity() {}
