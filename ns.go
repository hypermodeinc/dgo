/*
 * SPDX-FileCopyrightText: © Hypermode Inc. <hello@hypermode.com>
 * SPDX-License-Identifier: Apache-2.0
 */

package dgo

import (
	"context"
	"crypto/x509"
	"fmt"
	"math/rand"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	"github.com/dgraph-io/dgo/v240/protos/api"
	apiv25 "github.com/dgraph-io/dgo/v240/protos/api.v25"
)

type hmAuthCreds struct {
	apiKey string
}

func (a *hmAuthCreds) GetRequestMetadata(ctx context.Context, uri ...string) (
	map[string]string, error) {

	return map[string]string{"Authorization": fmt.Sprintf("Bearer %s", a.apiKey)}, nil
}

func (a *hmAuthCreds) RequireTransportSecurity() bool {
	return true
}

type clientOptions struct {
	gopts []grpc.DialOption
}

// ClientOption is a function that modifies the client options.
type ClientOption func(*clientOptions) error

// WithSystemCertPool will use the system cert pool and setup a TLS connection with Dgraph cluster.
func WithSystemCertPool() ClientOption {
	return func(o *clientOptions) error {
		pool, err := x509.SystemCertPool()
		if err != nil {
			return fmt.Errorf("failed to create system cert pool: %w", err)
		}

		creds := credentials.NewClientTLSFromCert(pool, "")
		o.gopts = append(o.gopts, grpc.WithTransportCredentials(creds))
		return nil
	}
}

// WithDgraphAPIKey will use the provided API key for authentication for Dgraph Cloud.
func WithDgraphAPIKey(apiKey string) ClientOption {
	return func(o *clientOptions) error {
		o.gopts = append(o.gopts, grpc.WithPerRPCCredentials(&authCreds{token: apiKey}))
		return nil
	}
}

// WithHypermodeAPIKey will use the provided API key for authentication for Hypermode.
func WithHypermodeAPIKey(apiKey string) ClientOption {
	return func(o *clientOptions) error {
		o.gopts = append(o.gopts, grpc.WithPerRPCCredentials(&hmAuthCreds{apiKey: apiKey}))
		return nil
	}
}

// WithGrpcOption will add a grpc.DialOption to the client.
// This is useful for setting custom  grpc options.
func WithGrpcOption(opt grpc.DialOption) ClientOption {
	return func(o *clientOptions) error {
		o.gopts = append(o.gopts, opt)
		return nil
	}
}

// NewClient creates a new Dgraph client for a single endpoint.
func NewClient(endpoint string, opts ...ClientOption) (*Dgraph, error) {
	return NewRoundRobinClient([]string{endpoint}, opts...)
}

// NewRoundRobinClient creates a new Dgraph client for a list
// of endpoints. It will round robin among the provided endpoints.
func NewRoundRobinClient(endpoints []string, opts ...ClientOption) (*Dgraph, error) {
	co := &clientOptions{}
	for _, opt := range opts {
		if err := opt(co); err != nil {
			return nil, err
		}
	}

	dc := make([]api.DgraphClient, len(endpoints))
	dcv25 := make([]apiv25.DgraphClient, len(endpoints))
	for i, endpoint := range endpoints {
		conn, err := grpc.NewClient(endpoint, co.gopts...)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to endpoint [%s]: %w", endpoint, err)
		}
		dc[i] = api.NewDgraphClient(conn)
		dcv25[i] = apiv25.NewDgraphClient(conn)
	}

	return &Dgraph{dc: dc, dcv25: dcv25}, nil
}

func (d *Dgraph) anyClientv25() apiv25.DgraphClient {
	//nolint:gosec
	return d.dcv25[rand.Intn(len(d.dcv25))]
}

func (d *Dgraph) CreateNamespace(ctx context.Context, name string, password string) error {
	dc := d.anyClientv25()
	req := &apiv25.CreateNamespaceRequest{NsName: name, Password: password}
	_, err := dc.CreateNamespace(ctx, req)
	return err
}

func (d *Dgraph) DropNamespace(ctx context.Context, name string) error {
	dc := d.anyClientv25()
	_, err := dc.DropNamespace(ctx, &apiv25.DropNamespaceRequest{NsName: name})
	return err
}

func (d *Dgraph) DropNamespaceWithID(ctx context.Context, nsID uint64) error {
	dc := d.anyClientv25()
	_, err := dc.DropNamespace(ctx, &apiv25.DropNamespaceRequest{NsId: nsID})
	return err
}

func (d *Dgraph) RenameNamespace(ctx context.Context, from string, to string) error {
	dc := d.anyClientv25()
	req := &apiv25.RenameNamespaceRequest{FromNs: from, ToNs: to}
	_, err := dc.RenameNamespace(ctx, req)
	return err
}

func (d *Dgraph) ListNamespaces(ctx context.Context) (map[string]*apiv25.Namespace, error) {
	dc := d.anyClientv25()
	resp, err := dc.ListNamespaces(ctx, &apiv25.ListNamespacesRequest{})
	if err != nil {
		return nil, err
	}
	return resp.NsList, nil
}
