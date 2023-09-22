//
// File: logical_plan.go
// Project: logicalplan
// File Created: 2023-09-07
// Author: machunxiao (machunxiao@shizhuang-inc.com)
// -----
// Last Modified By:  machunxiao (machunxiao@shizhuang-inc.com)
// Last Modified Time: 2023-09-07 21:39:59
// -----
//
// Copyright (C) 2023, Shanghai Poizon Information Technology Co., Ltd.
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

type LogicalPlan interface {
	fmt.Stringer

	GetTable() *datatype.Table
	GetChildren() []LogicalPlan
}

func Format(plan LogicalPlan, indent int) string {
	out := ""
	for i := 0; i < indent; i++ {
		out += "\t"
	}
	out += plan.String() + "\n"
	for _, sub := range plan.GetChildren() {
		out += Format(sub, indent+1)
	}
	return out
}
