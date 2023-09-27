//
// File: physical_plan.go
// Project: physicalplan
// File Created: 2023-09-08
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-08 11:14:56
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

package physicalplan

import (
	"fmt"
	"io"
)

type PhysicalPlan interface {
	fmt.Stringer
	io.Closer
	Open() error
	Next() any
	Children() []PhysicalPlan
}

func format0(plan PhysicalPlan) string {
	return format(plan, 0)
}

func format(plan PhysicalPlan, ident int) string {
	out := ""
	for i := 0; i < ident; i++ {
		out += "\t"
	}
	out += plan.String()
	out += "\n"
	for _, child := range plan.Children() {
		out += format(child, ident+1)
	}
	return out
}
