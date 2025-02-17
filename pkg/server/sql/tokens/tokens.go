//
// File: tokens.go
// Project: tokens
// File Created: 2025-01-05
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-07 23:42:09
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

package tokens

import "fmt"

type TokenKind int

type Token struct {
	fmt.Stringer
	kind     TokenKind
	text     string
	location *Location
}

type TokenList []*Token

func (token *Token) String() string {
	return fmt.Sprintf("Kind=%d,Text=%s(%d,%d)", token.kind, token.text, token.location.line, token.location.column)
}

func New(kind int, text string, ln, col uint32) *Token {
	tk := new(Token)
	tk.kind = TokenKind(kind)
	tk.text = text
	tk.location = Position(ln, col)
	return tk
}
