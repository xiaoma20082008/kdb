//
// File: token.go
// Project: dialect
// File Created: 2023-09-16
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-16 10:53:51
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

package dialect

type TokenType int32

const (
	EOF TokenType = -(iota + 1)
	Ident
	Symbol
	Integer
	Float
	String
	Comment
)

func (tt TokenType) String() string {
	switch tt {
	case Ident:
		return "Ident"
	case Symbol:
		return "Symbol"
	case Integer:
		return "Integer"
	case Float:
		return "Float"
	case String:
		return "String"
	case Comment:
		return "Comment"
	case EOF:
		return "EOF"
	default:
		return "Unknown"
	}
}

type Token struct {
	Text string
	Type TokenType

	Position Position
}

type TokenList struct {
	Offset int
	Tokens []*Token
}

func (t Token) String() string {
	return t.Text
}

func (list *TokenList) Next() *Token {
	list.Offset++
	return list.Tokens[list.Offset]
}

func (list *TokenList) Peek() *Token {
	return list.Tokens[list.Offset+1]
}

func (list *TokenList) Add(token *Token) {
	list.Tokens = append(list.Tokens, token)
}

func NewToken(text string, tt TokenType, offset, line, column int) *Token {
	return &Token{
		Text: text,
		Type: tt,
		Position: Position{
			Offset: offset,
			Line:   line,
			Column: column,
		},
	}
}

func NewTokenList(tokens []*Token) *TokenList {
	return &TokenList{
		Tokens: tokens,
		Offset: 0,
	}
}
