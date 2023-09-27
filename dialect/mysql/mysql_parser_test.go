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

func Test_mysqlParser_parseStmt(t *testing.T) {
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
		{
			name: "Insert",
			fields: fields{
				offset: 0,
				stream: dialect.NewTokenList([]*dialect.Token{
					dialect.NewToken("INSERT", dialect.Symbol, INSERT, 1, 0, 0),
					dialect.NewToken("INTO", dialect.Symbol, INTO, 1, 0, 0),
					dialect.NewToken("t1", dialect.Ident, -1, 1, 0, 0),
					dialect.NewToken("(", dialect.Symbol, LP, 1, 0, 0),
					dialect.NewToken("a", dialect.Ident, -1, 1, 0, 0),
					dialect.NewToken("b", dialect.Ident, -1, 1, 0, 0),
					dialect.NewToken("c", dialect.Ident, -1, 1, 0, 0),
					dialect.NewToken("d", dialect.Ident, -1, 1, 0, 0),
					dialect.NewToken(")", dialect.Symbol, RP, 1, 0, 0),
					dialect.NewToken("VALUES", dialect.Symbol, VALUES, 1, 0, 0),
					dialect.NewToken("(", dialect.Symbol, LP, 1, 0, 0),
					dialect.NewToken("101", dialect.String, -1, 1, 0, 0),
					dialect.NewToken("100.00", dialect.Float, -1, 1, 0, 0),
					dialect.NewToken("200", dialect.Integer, -1, 1, 0, 0),
					dialect.NewToken("?", dialect.Symbol, -1, 1, 0, 0),
					dialect.NewToken(")", dialect.Symbol, RP, 1, 0, 0),
				}),
			},
			want: &dialect.SqlInsert{
				Table: &dialect.SqlIdentifier{Id: "t1"},
				Columns: &dialect.SqlExprList{
					List: []dialect.SqlExpr{
						&dialect.SqlIdentifier{Id: "a"},
						&dialect.SqlIdentifier{Id: "b"},
						&dialect.SqlIdentifier{Id: "c"},
						&dialect.SqlIdentifier{Id: "d"},
					},
				},
				Values: &dialect.SqlExprList{
					List: []dialect.SqlExpr{
						&dialect.SqlString{Value: "a"},
						&dialect.SqlFloat{Value: 100.00},
						&dialect.SqlInteger{Value: 200},
						&dialect.SqlIdentifier{Id: "?"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parser := &mysqlParser{
				offset: tt.fields.offset,
				stream: tt.fields.stream,
			}
			got, err := parser.parseStmt()
			if (err != nil) != tt.wantErr {
				t.Errorf("mysqlParser.parseStmt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mysqlParser.parseStmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
