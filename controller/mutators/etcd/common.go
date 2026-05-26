//	Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package etcd

import (
	"github.com/golang/protobuf/proto"
	"github.com/m3db/m3/src/cluster/kv"
	pb "github.com/uber/aresdb/controller/generated/proto"
)

func addEntity(entityList pb.EntityList, name string) (res pb.EntityList, incarnation int, exist bool) {
	_ = "STUB: not implemented"
	return *new(pb.EntityList), 0, false
}

func readValue(etcdStore kv.TxnStore, key string, out proto.Message) (version int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func readEntityList(etcdStore kv.TxnStore, key string) (entityList pb.EntityList, version int, err error) {
	_ = "STUB: not implemented"
	return *new(pb.EntityList), 0, nil
}

func deleteEntity(entityList pb.EntityList, name string) (pb.EntityList, bool) {
	_ = "STUB: not implemented"
	return *new(pb.EntityList), false
}

func updateEntity(entityList pb.EntityList, name string) (pb.EntityList, bool) {
	_ = "STUB: not implemented"
	return *new(pb.EntityList), false
}

func find(entityList pb.EntityList, name string, doWithEntity func(*pb.EntityName)) (pb.EntityList, bool) {
	_ = "STUB: not implemented"
	return *new(pb.EntityList), false
}

func getHash(etcdStore kv.TxnStore, key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
