//
// File: join.go
// Project: logicalplan
// File Created: 2023-09-26
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-26 17:57:10
// -----
//
// Copyright (C) 2023, xiaoma20082008. All rights reserved.
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
	"kdb/storage"
)

type JoinType int32

const (
	LeftJoin JoinType = iota
	RightJoin
	InnerJoin
	SelfJoin
	CrossJoin
)

type Join struct {
	// left JOIN right ON xx = xx USING xx
	LogicalPlan

	Left      LogicalPlan
	Right     LogicalPlan
	JoinType  JoinType
	Condition LogicalExpr
}

func (j *Join) GetTable() storage.Table {
	return storage.NewTable("", "")
}

func (j *Join) GetChildren() []LogicalPlan {
	return []LogicalPlan{j.Left, j.Right}
}

func (j *Join) String() string {
	return fmt.Sprintf("Join: %s", "")
}
