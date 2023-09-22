//
// File: table.go
// Project: datatype
// File Created: 2023-09-07
// Author: machunxiao (machunxiao@shizhuang-inc.com)
// -----
// Last Modified By:  machunxiao (machunxiao@shizhuang-inc.com)
// Last Modified Time: 2023-09-07 21:33:48
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

package datatype

type Table struct {
	Catalog  string
	Database string
	Name     string

	Columns []*Column
}

func (t *Table) Project(indices []int) *Table {
	n := new(Table)
	n.Catalog = t.Catalog
	n.Database = t.Database
	n.Name = t.Name
	n.Columns = make([]*Column, 0)

	size := len(indices)
	for i := 0; i < size; i++ {
		_ = append(n.Columns, t.Columns[i])
	}
	return n
}

func (t *Table) Filter(names []string) *Table {
	n := new(Table)
	return n
}
