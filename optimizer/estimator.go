//
// File: estimator.go
// Project: optimizer
// File Created: 2023-09-13
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-13 18:15:44
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

package optimizer

import (
	"kdb/logicalplan"
)

type Estimator interface {
	Estimate(plan logicalplan.LogicalPlan) *Cost
}

type estimator struct{}

func (e *estimator) Estimate(plan logicalplan.LogicalPlan) *Cost {
	return nil
}

func Estimate(plan logicalplan.LogicalPlan) *Cost {
	return &Cost{}
}
