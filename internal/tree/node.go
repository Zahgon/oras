/*
Copyright The ORAS Authors.
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

// Package tree pretty prints trees
package tree

// Node represents a tree node.
type Node struct {
	Value any
	Nodes []*Node
}

// New creates a new tree / root node.
func New(value any) *Node { _ = "STUB: not implemented"; return nil }

// Add adds a leaf node.
func (n *Node) Add(value any) *Node { _ = "STUB: not implemented"; return nil }

// AddPath adds a chain of nodes.
func (n *Node) AddPath(values ...any) *Node { _ = "STUB: not implemented"; return nil }

// Find finds the child node with the target value.
// Nil if not found.
func (n *Node) Find(value any) *Node { _ = "STUB: not implemented"; return nil }
