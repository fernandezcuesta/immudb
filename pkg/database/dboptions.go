/*
Copyright 2022 Codenotary Inc. All rights reserved.

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

package database

import "github.com/codenotary/immudb/embedded/store"

const (
	DefaultDbRootPath     = "./data"
	DefaultReadTxPoolSize = 128
)

//Options database instance options
type Options struct {
	dbRootPath string

	storeOpts *store.Options

	replica bool

	corruptionChecker bool

	readTxPoolSize int
}

// DefaultOption Initialise Db Optionts to default values
func DefaultOption() *Options {
	return &Options{
		dbRootPath:     DefaultDbRootPath,
		storeOpts:      store.DefaultOptions(),
		readTxPoolSize: DefaultReadTxPoolSize,
	}
}

// WithDbRootPath sets the directory in which this database will reside
func (o *Options) WithDBRootPath(Path string) *Options {
	o.dbRootPath = Path
	return o
}

// GetDbRootPath returns the directory in which this database resides
func (o *Options) GetDBRootPath() string {
	return o.dbRootPath
}

// WithCorruptionChecker sets if corruption checker should start for this database instance
func (o *Options) WithCorruptionChecker(cc bool) *Options {
	o.corruptionChecker = cc
	return o
}

// GetCorruptionChecker returns if corruption checker should start for this database instance
func (o *Options) GetCorruptionChecker() bool {
	return o.corruptionChecker
}

// WithStoreOptions sets backing store options
func (o *Options) WithStoreOptions(storeOpts *store.Options) *Options {
	o.storeOpts = storeOpts
	return o
}

// GetStoreOptions returns backing store options
func (o *Options) GetStoreOptions() *store.Options {
	return o.storeOpts
}

// AsReplica sets if the database is a replica
func (o *Options) AsReplica(replica bool) *Options {
	o.replica = replica
	return o
}

func (o *Options) WithReadTxPoolSize(txPoolSize int) *Options {
	o.readTxPoolSize = txPoolSize
	return o
}

func (o *Options) GetTxPoolSize() int {
	return o.readTxPoolSize
}
