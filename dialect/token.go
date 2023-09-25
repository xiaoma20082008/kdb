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

type TokenType int
type TokenKind int

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
	Kind TokenKind

	Offset int // byte offset, starting at 0
	Line   int // line number, starting at 1
	Column int // column number, starting at 1 (character count per line)
}

func (t Token) String() string {
	return t.Text
}

func NewToken(text string, tt TokenType, kind TokenKind, line, offset, column int) *Token {
	return &Token{
		Text:   text,
		Type:   tt,
		Kind:   kind,
		Line:   line,
		Offset: offset,
		Column: column,
	}
}
