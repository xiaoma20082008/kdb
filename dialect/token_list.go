//
// File: token_list.go
// Project: dialect
// File Created: 2023-09-23
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2023-09-23 00:33:08
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

package dialect

type TokenList struct {
	Offset int
	Tokens []*Token
}

func (list *TokenList) Next() *Token {
	if list.Offset < len(list.Tokens) {
		list.Offset++
		return list.Tokens[list.Offset]
	} else {
		return nil
	}
}

func (list *TokenList) Peek() *Token {
	if list.Offset+1 < len(list.Tokens) {
		return list.Tokens[list.Offset+1]
	} else {
		return nil
	}
}

func (list *TokenList) Curr() *Token {
	if list.Offset < len(list.Tokens) {
		return list.Tokens[list.Offset]
	} else {
		return nil
	}
}

func NewTokenList(tokens []*Token) *TokenList {
	return &TokenList{
		Tokens: tokens,
		Offset: 0,
	}
}
