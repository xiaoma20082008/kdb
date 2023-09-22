//
// File: executor_test.go
// Project: executor
// File Created: 2023-09-21
// Author: machunxiao (machunxiao@shizhuang-inc.com)
// -----
// Last Modified By:  machunxiao (machunxiao@shizhuang-inc.com)
// Last Modified Time: 2023-09-21 21:47:09
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

package executor

import (
	_ "kdb/dialect/mysql"
	"testing"
)

func TestExecute(t *testing.T) {
	type args struct {
		sql  string
		args map[int32]any
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Execute(tt.args.sql, tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("Execute() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
