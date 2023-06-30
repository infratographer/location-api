package graphapi_test

import (
	"context"

	"github.com/brianvoe/gofakeit/v6"
	"go.infratographer.com/x/gidx"

	ent "go.infratographer.com/location-api/internal/ent/generated"
)

type LocationBuilder struct {
	Name        string
	Description *string
	OwnerID     gidx.PrefixedID
}

func (l *LocationBuilder) MustNew(ctx context.Context) *ent.Location {
	if l.Name == "" {
		l.Name = gofakeit.AppName()
	}

	if l.Description == nil {
		desc := gofakeit.HipsterSentence(10)
		l.Description = &desc
	}

	if l.OwnerID == "" {
		l.OwnerID = gidx.MustNewID(ownerPrefix)
	}

	return EntClient.Location.Create().SetName(l.Name).SetDescription(*l.Description).SetOwnerID(l.OwnerID).SaveX(ctx)
}
