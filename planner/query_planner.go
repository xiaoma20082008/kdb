//
// File: query_planner.go
// Project: planner
// File Created: 2023-09-08
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-08 11:28:00
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

package planner

import (
	"kdb/logicalplan"
	"kdb/physicalplan"
)

func CreatePhysicalPlan(logical logicalplan.LogicalPlan) (physicalplan.PhysicalPlan, error) {
	return nil, nil
}