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
package utils

const (
	// AresController sets the name for ares controller
	AresController = "ares-controller"
	// AresSubscriber sets the name for ares subscriber
	AresSubscriber = "ares-subscriber"
	// AresDataNode sets the name for ares datanode
	AresDataNode = "ares-datanode"
	// AresBroker sets the name for ares broker
	AresBroker = "ares-broker"
)

// etcd keys

// NamespaceListKey builds key for namespace list
func NamespaceListKey() string { _ = "STUB: not implemented"; return "" }

// NamespaceKey builds key for namespace
func NamespaceKey(namespace string) string { _ = "STUB: not implemented"; return "" }

// SchemaListKey builds key for schema list
func SchemaListKey(namespace string) string { _ = "STUB: not implemented"; return "" }

// JobListKey builds key for job list
func JobListKey(namespace string) string { _ = "STUB: not implemented"; return "" }

// JobAssignmentsListKey builds key for job assignments
func JobAssignmentsListKey(namespace string) string { _ = "STUB: not implemented"; return "" }

// InstanceListKey builds key for job assignments
func InstanceListKey(namespace string) string { _ = "STUB: not implemented"; return "" }

// SchemaKey builds key for schema
func SchemaKey(namespace, name string) string { _ = "STUB: not implemented"; return "" }

// JobKey builds key for job config
func JobKey(namespace, name string) string { _ = "STUB: not implemented"; return "" }

// JobAssignmentsKey builds key for job assignments
func JobAssignmentsKey(namespace, name string) string { _ = "STUB: not implemented"; return "" }

// InstanceKey builds key for instance
func InstanceKey(namespace, name string) string { _ = "STUB: not implemented"; return "" }

// EnumNodeListKey builds the key for enum node list
func EnumNodeListKey(namespace, table string, incarnation, columnID int) string {
	_ = "STUB: not implemented"
	return ""
}

// EnumNodeKey builds the key for enum node
func EnumNodeKey(namespace, table string, incarnation int, columnID, nodeID int) string {
	_ = "STUB: not implemented"
	return ""
}

// SubscriberServiceName builds the subscriber service name
func SubscriberServiceName(namespace string) string { _ = "STUB: not implemented"; return "" }

// DataNodeServiceName builds the data node service name
func DataNodeServiceName(namespace string) string { _ = "STUB: not implemented"; return "" }

// BrokerServiceName builds the broker service name
func BrokerServiceName(namespace string) string { _ = "STUB: not implemented"; return "" }
