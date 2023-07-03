// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
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
