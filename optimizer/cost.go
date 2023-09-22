//
// File: cost.go
// Project: optimizer
// File Created: 2023-09-06
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-06 18:34:28
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

import "fmt"

type Cost struct {
	Monetary float64
	Network  float64
	Disk     float64
	Memory   float64
	Gpu      float64
	Cpu      float64
	DataSize int64
}

func (cost Cost) String() string {
	return fmt.Sprintf("cpu=%f,gpu=%f,memory=%f,disk=%f,network=%f,monetary=%f",
		cost.Cpu, cost.Gpu, cost.Memory, cost.Disk, cost.Network, cost.Monetary)
}
