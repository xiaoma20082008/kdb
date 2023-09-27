//
// File: filter.go
// Project: physicalplan
// File Created: 2023-09-25
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-25 17:10:51
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

package physicalplan

type FilterExec struct {
	PhysicalPlan

	input  PhysicalPlan
	filter any
}

func (filter *FilterExec) String() string {
	return format0(filter)
}

func (filter *FilterExec) Close() error {
	return filter.input.Close()
}

func (filter *FilterExec) Open() error {
	return filter.input.Open()
}

func (filter *FilterExec) Next() any {
	if data := filter.input.Next(); data != nil {
	}
	return nil
}

func (filter *FilterExec) Children() []PhysicalPlan {
	return []PhysicalPlan{
		filter.input,
	}
}
