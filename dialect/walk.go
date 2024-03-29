//
// File: walk.go
// Project: dialect
// File Created: 2023-09-18
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-18 11:46:11
// -----
//
// Copyright (C) xiaoma20082008. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dialect

type Visitor interface {
	Visit(node SqlNode) bool
}

func Walk(v Visitor, node SqlNode) {
	if node == nil || !v.Visit(node) {
		return
	}
	switch n := node.(type) {
	case *SqlComment:
	case *SqlIdentifier:
	case *SqlInsert:
	case *SqlDelete:
	case *SqlUpdate:
	case *SqlSelect:
		v.Visit(n.Columns)
		v.Visit(n.From)
		v.Visit(n.Where)
		v.Visit(n.GroupBy)
		v.Visit(n.Having)
		v.Visit(n.OrderBy)
		v.Visit(n.Limit)
		v.Visit(n.Offset)
	}
}
