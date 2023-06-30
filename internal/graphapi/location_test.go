package graphapi_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.infratographer.com/permissions-api/pkg/permissions"
	"go.infratographer.com/x/gidx"

	"go.infratographer.com/location-api/internal/ent/generated"
)

func TestQuery_IPAddress(t *testing.T) {
	client := graphTestClient()
	ctx := context.Background()

	// Permit request
	ctx = context.WithValue(ctx, permissions.CheckerCtxKey, permissions.DefaultAllowChecker)

	loc1 := (&LocationBuilder{}).MustNew(ctx)

	testCases := []struct {
		name     string
		queryID  gidx.PrefixedID
		expected *generated.Location
		errorMsg string
	}{
		{
			name:     "happy path location",
			queryID:  loc1.ID,
			expected: loc1,
		},
		{
			name:     "invalid-id",
			queryID:  gidx.MustNewID("testing"),
			errorMsg: "location not found",
		},
	}

	for _, tc := range testCases {
		t.Run("Get "+tc.name, func(t *testing.T) {
			resp, err := client.GetLocationByID(ctx, tc.queryID)

			if tc.errorMsg != "" {
				require.Error(t, err)
				assert.ErrorContains(t, err, tc.errorMsg)
				assert.Nil(t, resp)

				return
			}

			require.NoError(t, err)
			require.NotNil(t, resp)
			require.NotNil(t, resp.Location)
		})
	}
}

// TODO: test crud operations
