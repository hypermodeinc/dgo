/*
 * SPDX-FileCopyrightText: © Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package dgo_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestQueryNoDiscardTxn(t *testing.T) {
	dg, cancel := getDgraphClient()
	defer cancel()

	txn := dg.NewTxn()
	ctx := context.Background()

	_, err := txn.Query(ctx, `{me(){}me(){}}`)
	require.Error(t, err)

	resp, err := txn.Query(ctx, `{me(){}}`)
	require.NoError(t, err)
	require.GreaterOrEqual(t, len(resp.GetHdrs()), 1)
}
