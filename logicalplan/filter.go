//
// File: filter.go
// Project: logicalplan
// File Created: 2023-09-07
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-07 21:51:50
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

package logicalplan

import (
	"fmt"

	"kdb/datatype"
)

type Filter struct {
	LogicalPlan

	input LogicalPlan
	expr  LogicalExpr
}

func (f *Filter) GetTable() *datatype.Table {
	return f.input.GetTable()
}

func (f *Filter) GetChildren() []LogicalPlan {
	return f.input.GetChildren()
}

func (f *Filter) String() string {
	return fmt.Sprintf("Filter: %s", f.expr.String())
}

func NewFilter(input LogicalPlan, expr LogicalExpr) *Filter {
	f := new(Filter)
	f.input = input
	f.expr = expr
	return f
}
