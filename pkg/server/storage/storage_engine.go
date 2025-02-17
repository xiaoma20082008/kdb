//
// File: storage_engine.go
// Project: storage
// File Created: 2025-02-05
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-02-05 21:42:51
// Last Modified By: xiaoma20082008 (mmccxx2519@gmail.com>)
// ------------------------------------------------------------------------
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

package storage

import "io"

type StorageEngine interface {
	io.Closer
	// 初始化存储引擎
	Init(config map[string]interface{}) error

	// 事务操作
	BeginTx(IsolationLevel) (Transaction, error)

	// 获取存储引擎的统计信息
	Stats() map[string]interface{}
}

type Table interface {
	Insert(Tuple) error
	Delete() error
	Update() error
	TableScan() (Cursor, error)
	IndexScan() (Cursor, error)
	Stats() map[string]interface{}
}

type Tuple interface{}

type Cursor interface {
}
