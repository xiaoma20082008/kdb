//
// File: scanner_test.go
// Project: scanner
// File Created: 2025-02-15
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-02-15 01:49:08
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

package scanner

import (
	"reflect"
	"testing"

	"github.com/xiaoma20082008/kdb/pkg/server/sql/tokens"
)

func TestScanner_Next(t *testing.T) {
	type fields struct {
		text string
	}
	tests := []struct {
		name    string
		fields  fields
		want    *tokens.Token
		wantErr bool
	}{
		// {
		// 	name:    "keyword",
		// 	fields:  fields{text: "if"},
		// 	want:    tokens.New(0, "if", 1, 3),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "ident",
		// 	fields:  fields{text: "asdf"},
		// 	want:    tokens.New(0, "asdf", 1, 5),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "number",
		// 	fields:  fields{text: "123"},
		// 	want:    tokens.New(0, "123", 1, 4),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "number",
		// 	fields:  fields{text: "123.456"},
		// 	want:    tokens.New(0, "123.456", 1, 8),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "number",
		// 	fields:  fields{text: "0.1"},
		// 	want:    tokens.New(0, "0.1", 1, 4),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "number",
		// 	fields:  fields{text: "0b101"},
		// 	want:    tokens.New(0, "0b101", 1, 6),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "number",
		// 	fields:  fields{text: "0xABCDEF"},
		// 	want:    tokens.New(0, "0xABCDEF", 1, 9),
		// 	wantErr: false,
		// },
		// {
		// 	name:    "number",
		// 	fields:  fields{text: "01234567"},
		// 	want:    tokens.New(0, "01234567", 1, 9),
		// 	wantErr: false,
		// },
		{
			name:    "number",
			fields:  fields{text: "'hello'"},
			want:    tokens.New(0, "hello", 1, 7),
			wantErr: false,
		},
		{
			name:    "number",
			fields:  fields{text: "\r\n\t\f 'hello'"},
			want:    tokens.New(0, "hello", 2, 10),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.text)
			got, err := s.Next()
			if (err != nil) != tt.wantErr {
				t.Errorf("Scanner.Next() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Scanner.Next() = %s, want %s", got.String(), tt.want.String())
			}
		})
	}
}
