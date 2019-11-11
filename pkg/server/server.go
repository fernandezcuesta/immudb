/*
Copyright 2019 vChain, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"context"
	"fmt"
	"net"

	"github.com/dgraph-io/badger/v2"
	"github.com/dgraph-io/badger/v2/options"
	"google.golang.org/grpc"

	"github.com/codenotary/immudb/pkg/db"
	"github.com/codenotary/immudb/pkg/schema"
)

type ImmuServer struct {
	Topic *db.Topic
}

func Run(address string, dir string) error {
	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	b, err := makeBadger(dir, "immudb")
	if err != nil {
		return err
	}
	var serverOptions []grpc.ServerOption
	server := grpc.NewServer(serverOptions...)
	schema.RegisterImmuServiceServer(server, &ImmuServer{
		Topic: db.NewTopic(b),
	})
	return server.Serve(listener)
}

func (s ImmuServer) Set(ctx context.Context, sr *schema.SetRequest) (*schema.SetResponse, error) {
	fmt.Println("Set", sr.Key)
	if err := s.Topic.Set(sr.Key, sr.Value); err != nil {
		return nil, err
	}
	return &schema.SetResponse{
		Status: 0,
	}, nil
}

func (s ImmuServer) Get(ctx context.Context, gr *schema.GetRequest) (*schema.GetResponse, error) {
	fmt.Println("Get", gr.Key)
	value, err := s.Topic.Get(gr.Key)
	if err != nil {
		return nil, err
	}
	return &schema.GetResponse{
		Status: 0,
		Key:    gr.Key,
		Value:  value,
	}, nil
}

func makeBadger(dir string, name string) (*badger.DB, error) {
	opts := badger.
		DefaultOptions(dir + "/" + name).
		WithTableLoadingMode(options.LoadToRAM).
		WithCompressionType(options.None).
		WithSyncWrites(false).
		WithEventLogging(false)
	return badger.Open(opts)
}
