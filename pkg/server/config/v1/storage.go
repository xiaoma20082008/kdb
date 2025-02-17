//
// File: storage.go
// Project: v1
// File Created: 2025-01-21
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-21 23:46:33
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

package v1

import (
	"github.com/xiaoma20082008/kdb/pkg/server/common"
)

type StorageConfig struct {
	BaseDir  string `json:"base_dir,omitempty"`
	DataDir  string `json:"data_dir,omitempty"`
	LogsDir  string `json:"logs_dir,omitempty"`
	PageSize int    `json:"page_size,omitempty"`
	//
	Engine string      `json:"engine,omitempty"`
	Config *common.Ini `json:"config,omitempty"`
}
