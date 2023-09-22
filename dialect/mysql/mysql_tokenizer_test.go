//
// File: mysql_tokenizer_test.go
// Project: mysql
// File Created: 2023-09-21
// Author: machunxiao (machunxiao@shizhuang-inc.com)
// -----
// Last Modified By:  machunxiao (machunxiao@shizhuang-inc.com)
// Last Modified Time: 2023-09-21 20:19:39
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

package mysql

import (
	"kdb/dialect"
	"reflect"
	"testing"
)

func Test_mysqlTokenizer_tokenize(t *testing.T) {
	type fields struct {
		sql    string
		line   int
		offset int
		length int
	}
	genField := func(sql string) fields {
		return fields{
			sql:    sql,
			line:   1,
			offset: 0,
			length: len(sql),
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    *dialect.TokenList
		wantErr bool
	}{
		{
			name: "tokenize",
			fields: genField(`
			hello 123 ? + - * / '123'
			`),
			want: &dialect.TokenList{
				Offset: 0,
				Tokens: []*dialect.Token{
					dialect.NewToken("hello", dialect.Ident, 0, 0, 0),
					dialect.NewToken("123", dialect.Integer, 0, 0, 0),
					dialect.NewToken("?", dialect.Symbol, 0, 0, 0),
					dialect.NewToken("+", dialect.Symbol, 0, 0, 0),
					dialect.NewToken("-", dialect.Symbol, 0, 0, 0),
					dialect.NewToken("*", dialect.Symbol, 0, 0, 0),
					dialect.NewToken("/", dialect.Symbol, 0, 0, 0),
					dialect.NewToken("123", dialect.String, 0, 0, 0),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tokenizer := &mysqlTokenizer{
				sql:    tt.fields.sql,
				line:   tt.fields.line,
				offset: tt.fields.offset,
				length: tt.fields.length,
			}
			got, err := tokenizer.tokenize()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlTokenizer.tokenize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlTokenizer.tokenize() = %v, want %v", got, tt.want)
			}
		})
	}
}
