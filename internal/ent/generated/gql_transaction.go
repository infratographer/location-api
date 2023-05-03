// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
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
	"context"
	"database/sql/driver"
	"errors"
)

// OpenTx opens a transaction and returns a transactional
// context along with the created transaction.
func (c *Client) OpenTx(ctx context.Context) (context.Context, driver.Tx, error) {
	tx, err := c.Tx(ctx)
	if err != nil {
		return nil, nil, err
	}
	ctx = NewTxContext(ctx, tx)
	ctx = NewContext(ctx, tx.Client())
	return ctx, tx, nil
}

// OpenTxFromContext open transactions from client stored in context.
func OpenTxFromContext(ctx context.Context) (context.Context, driver.Tx, error) {
	client := FromContext(ctx)
	if client == nil {
		return nil, nil, errors.New("no client attached to context")
	}
	return client.OpenTx(ctx)
}
