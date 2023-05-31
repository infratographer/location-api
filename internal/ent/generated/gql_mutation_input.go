// Copyright Infratographer, Inc. and/or licensed to Infratographer, Inc. under one
// or more contributor license agreements. Licensed under the Elastic License 2.0;
// you may not use this file except in compliance with the Elastic License 2.0.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"go.infratographer.com/x/gidx"
)

// CreateLocationInput represents a mutation input for creating locations.
type CreateLocationInput struct {
	Name        string
	Description *string
	OwnerID     gidx.PrefixedID
}

// Mutate applies the CreateLocationInput on the LocationMutation builder.
func (i *CreateLocationInput) Mutate(m *LocationMutation) {
	m.SetName(i.Name)
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
	m.SetOwnerID(i.OwnerID)
}

// SetInput applies the change-set in the CreateLocationInput on the LocationCreate builder.
func (c *LocationCreate) SetInput(i CreateLocationInput) *LocationCreate {
	i.Mutate(c.Mutation())
	return c
}

// UpdateLocationInput represents a mutation input for updating locations.
type UpdateLocationInput struct {
	Name             *string
	ClearDescription bool
	Description      *string
}

// Mutate applies the UpdateLocationInput on the LocationMutation builder.
func (i *UpdateLocationInput) Mutate(m *LocationMutation) {
	if v := i.Name; v != nil {
		m.SetName(*v)
	}
	if i.ClearDescription {
		m.ClearDescription()
	}
	if v := i.Description; v != nil {
		m.SetDescription(*v)
	}
}

// SetInput applies the change-set in the UpdateLocationInput on the LocationUpdate builder.
func (c *LocationUpdate) SetInput(i UpdateLocationInput) *LocationUpdate {
	i.Mutate(c.Mutation())
	return c
}

// SetInput applies the change-set in the UpdateLocationInput on the LocationUpdateOne builder.
func (c *LocationUpdateOne) SetInput(i UpdateLocationInput) *LocationUpdateOne {
	i.Mutate(c.Mutation())
	return c
}
