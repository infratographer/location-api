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

package eventhooks

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	"go.infratographer.com/location-api/internal/ent/generated"
	"go.infratographer.com/location-api/internal/ent/generated/hook"
	"go.infratographer.com/x/events"
	"go.infratographer.com/x/gidx"
)

func LocationHooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.LocationFunc(func(ctx context.Context, m *generated.LocationMutation) (ent.Value, error) {
					var err error
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					changeset := []events.FieldChange{}
					cv_created_at := ""
					created_at, ok := m.CreatedAt()

					if ok {
						cv_created_at = created_at.Format(time.RFC3339)
						pv_created_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldCreatedAt(ctx)
							if err != nil {
								pv_created_at = "<unknown>"
							} else {
								pv_created_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "created_at",
							PreviousValue: pv_created_at,
							CurrentValue:  cv_created_at,
						})
					}

					cv_updated_at := ""
					updated_at, ok := m.UpdatedAt()

					if ok {
						cv_updated_at = updated_at.Format(time.RFC3339)
						pv_updated_at := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldUpdatedAt(ctx)
							if err != nil {
								pv_updated_at = "<unknown>"
							} else {
								pv_updated_at = ov.Format(time.RFC3339)
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "updated_at",
							PreviousValue: pv_updated_at,
							CurrentValue:  cv_updated_at,
						})
					}

					cv_name := ""
					name, ok := m.Name()

					if ok {
						cv_name = fmt.Sprintf("%s", fmt.Sprint(name))
						pv_name := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldName(ctx)
							if err != nil {
								pv_name = "<unknown>"
							} else {
								pv_name = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "name",
							PreviousValue: pv_name,
							CurrentValue:  cv_name,
						})
					}

					cv_description := ""
					description, ok := m.Description()

					if ok {
						cv_description = fmt.Sprintf("%s", fmt.Sprint(description))
						pv_description := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldDescription(ctx)
							if err != nil {
								pv_description = "<unknown>"
							} else {
								pv_description = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "description",
							PreviousValue: pv_description,
							CurrentValue:  cv_description,
						})
					}

					cv_owner_id := ""
					owner_id, ok := m.OwnerID()
					if !ok && !m.Op().Is(ent.OpCreate) {
						// since we are doing an update or delete and these fields didn't change, load the "old" value
						owner_id, err = m.OldOwnerID(ctx)
						if err != nil {
							return nil, err
						}
					}
					additionalSubjects = append(additionalSubjects, owner_id)

					if ok {
						cv_owner_id = fmt.Sprintf("%s", fmt.Sprint(owner_id))
						pv_owner_id := ""
						if !m.Op().Is(ent.OpCreate) {
							ov, err := m.OldOwnerID(ctx)
							if err != nil {
								pv_owner_id = "<unknown>"
							} else {
								pv_owner_id = fmt.Sprintf("%s", fmt.Sprint(ov))
							}
						}

						changeset = append(changeset, events.FieldChange{
							Field:         "owner_id",
							PreviousValue: pv_owner_id,
							CurrentValue:  cv_owner_id,
						})
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
						FieldChanges:         changeset,
					}

					// complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					if err := m.EventsPublisher.PublishChange(ctx, "location", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),

		// Delete Hook
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.LocationFunc(func(ctx context.Context, m *generated.LocationMutation) (ent.Value, error) {
					additionalSubjects := []gidx.PrefixedID{}

					objID, ok := m.ID()
					if !ok {
						return nil, fmt.Errorf("object doesn't have an id %s", objID)
					}

					dbObj, err := m.Client().Location.Get(ctx, objID)
					if err != nil {
						return nil, fmt.Errorf("failed to load object to get values for event, err %w", err)
					}

					additionalSubjects = append(additionalSubjects, dbObj.OwnerID)

					// we have all the info we need, now complete the mutation before we process the event
					retValue, err := next.Mutate(ctx, m)
					if err != nil {
						return retValue, err
					}

					msg := events.ChangeMessage{
						EventType:            eventType(m.Op()),
						SubjectID:            objID,
						AdditionalSubjectIDs: additionalSubjects,
						Timestamp:            time.Now().UTC(),
					}

					if err := m.EventsPublisher.PublishChange(ctx, "location", msg); err != nil {
						return nil, fmt.Errorf("failed to publish change: %w", err)
					}

					return retValue, nil
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

func EventHooks(c *generated.Client) {
	c.Location.Use(LocationHooks()...)

}

func eventType(op ent.Op) string {
	switch op {
	case ent.OpCreate:
		return string(events.CreateChangeType)
	case ent.OpUpdate, ent.OpUpdateOne:
		return string(events.UpdateChangeType)
	case ent.OpDelete, ent.OpDeleteOne:
		return string(events.DeleteChangeType)
	default:
		return "unknown"
	}
}
