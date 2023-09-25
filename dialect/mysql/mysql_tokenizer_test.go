//
// File: mysql_tokenizer_test.go
// Project: mysql
// File Created: 2023-09-21
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-21 20:19:39
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
			123.456 
			'123.456'
			`),
			want: &dialect.TokenList{
				Offset: 0,
				Tokens: []*dialect.Token{
					dialect.NewToken("hello", dialect.Ident, -1, 2, 4, 9),
					dialect.NewToken("123", dialect.Integer, -1, 2, 10, 13),
					dialect.NewToken("?", dialect.Symbol, QM, 2, 14, 15),
					dialect.NewToken("+", dialect.Symbol, PLUS, 2, 16, 17),
					dialect.NewToken("-", dialect.Symbol, MINUS, 2, 18, 19),
					dialect.NewToken("*", dialect.Symbol, STAR, 2, 20, 21),
					dialect.NewToken("/", dialect.Symbol, SLASH, 2, 22, 23),
					dialect.NewToken("123", dialect.String, -1, 2, 25, 29),
					dialect.NewToken("123.456", dialect.Float, -1, 3, 33, 40),
					dialect.NewToken("123.456", dialect.String, -1, 4, 46, 54),
					dialect.NewToken("", dialect.EOF, -1, -1, -1, -1),
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
