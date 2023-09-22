//
// File: mysql_parser_test.go
// Project: mysql
// File Created: 2023-09-21
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-21 20:19:29
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

package mysql

import (
	"kdb/dialect"
	"reflect"
	"testing"
)

func Test_mysqlParser_parse(t *testing.T) {
	type fields struct {
		offset int
		stream *dialect.TokenList
	}
	tests := []struct {
		name    string
		fields  fields
		want    dialect.SqlStmt
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &mysqlParser{
				offset: tt.fields.offset,
				stream: tt.fields.stream,
			}
			got, err := parser.parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlParser.parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlParser.parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
