//  Copyright (c) 2017-2018 Uber Technologies, Inc.
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

// Package memstore has to put test factory here since otherwise we will have a
// memstore -> utils -> memstore import cycle.

package tests

import (
	"github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/vectors"
	"github.com/uber/aresdb/utils"

	"sync"
)

// TestFactoryT creates memstore test objects from text file
type TestFactoryBase struct {
	RootPath string
	utils.FileSystem
	// functions to do real vp conversion, need to pass in from caller
	ToArchiveVectorParty func(common.VectorParty, sync.Locker) common.ArchiveVectorParty
	ToLiveVectorParty    func(common.VectorParty) common.LiveVectorParty
	ToVectorParty        func(*RawVectorParty, bool) (common.VectorParty, error)
}

type rawBatch struct {
	Columns []string `yaml:"columns"`
}

type RawVectorParty struct {
	DataType  string   `yaml:"data_type"`
	Length    int      `yaml:"length"`
	Values    []string `yaml:"values"`
	HasCounts bool     `yaml:"has_counts"`
}

type rawVector struct {
	DataType string   `yaml:"data_type"`
	Length   int      `yaml:"length"`
	Values   []string `yaml:"values"`
}

// rawUpsertBatch represents the upsert batch format in a yaml file. Each individual upsert batch consists of two parts
// columns and rows. Each column needs to specify the column type and column id. Column type need to be a valid type
// defined in common.data_type.go. Each row consists of column values which are comma splitted.
type rawUpsertBatch struct {
	Columns []struct {
		ColumnID int    `yaml:"column_id"`
		DataType string `yaml:"data_type"`
	} `yaml:"columns"`
	Rows []string `yaml:"rows"`
}

// ReadArchiveBatch read batch and do pruning for every columns.
func (t TestFactoryBase) ReadArchiveBatch(name string) (*common.Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadLiveBatch read batch and skip pruning for every columns.
func (t TestFactoryBase) ReadLiveBatch(name string) (*common.Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadBatch returns a batch given batch name. Batch will be searched
// under testing/data/batches folder. Prune tells whether need to prune
// the columns after column contruction.
func (t TestFactoryBase) ReadBatch(name string, forLiveVP bool) (*common.Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TestFactoryBase) readBatchFromFile(path string, forLiveVP bool) (*common.Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rb *rawBatch) toBatch(t TestFactoryBase, forLiveVP bool) (*common.Batch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadArchiveVectorParty loads a vector party and prune it after construction.
func (t TestFactoryBase) ReadArchiveVectorParty(name string, locker sync.Locker) (common.ArchiveVectorParty, error) {
	_ = "STUB: not implemented"
	return *new(common.ArchiveVectorParty), nil
}

// ReadLiveVectorParty loads a vector party and skip pruning.
func (t TestFactoryBase) ReadLiveVectorParty(name string) (common.LiveVectorParty, error) {
	_ = "STUB: not implemented"
	return *new(common.LiveVectorParty), nil
}

// ReadVectorParty returns a vector party given vector party name. Vector party
// will be searched under testing/data/vps folder. Prune tells whether to prune this
// column.
func (t TestFactoryBase) ReadVectorParty(name string, forLiveVP bool) (common.VectorParty, error) {
	_ = "STUB: not implemented"
	return *new(common.VectorParty), nil
}

func (t TestFactoryBase) readVectorPartyFromFile(path string, forLiveVP bool) (common.VectorParty, error) {
	_ = "STUB: not implemented"
	return *new(common.VectorParty), nil
}

// ReadVector returns a vector given vector name. Vector will
// be searched under testing/data/vectors folder.
func (t TestFactoryBase) ReadVector(name string) (*vectors.Vector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setDataValue(v *vectors.Vector, idx int, val common.DataValue) {
	_ = "STUB: not implemented"
	return
}

func (rv *rawVector) toVector() (*vectors.Vector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TestFactoryBase) readVectorFromFile(path string) (*vectors.Vector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadUpsertBatch returns a pointer to UpsertBatch given the upsert batch name.
func (t TestFactoryBase) ReadUpsertBatch(name string) (*common.UpsertBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (ru *rawUpsertBatch) toUpsertBatch() (*common.UpsertBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t TestFactoryBase) readUpsertBatchFromFile(path string) (*common.UpsertBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
