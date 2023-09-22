//
// File: mysql_tokenizer.go
// Project: mysql
// File Created: 2023-09-21
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-21 20:11:32
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
	"errors"
	"fmt"
	"kdb/dialect"

	"golang.org/x/exp/slog"
)

type mysqlTokenizer struct {
	sql string

	line   int
	offset int
	length int
}

func (tokenizer *mysqlTokenizer) tokenize() (*dialect.TokenList, error) {
	list := dialect.TokenList{
		Tokens: make([]*dialect.Token, 0),
		Offset: 0,
	}
	for {
		token, err := tokenizer.next()
		if err == nil {
			if token == nil {
				break
			} else {
				list.Add(token)
			}
		} else {
			return nil, err
		}
	}
	return &list, nil
}

func (tokenizer *mysqlTokenizer) next() (*dialect.Token, error) {
	tokenizer.skip()
	if tokenizer.offset >= tokenizer.length {
		return nil, nil
	}
	ch := tokenizer.sql[tokenizer.offset]
	if isLetter(ch) || ch == '`' {
		return tokenizer.scanIdent()
	} else if isNumber(ch) {
		return tokenizer.scanNumber()
	} else if ch == '\'' {
		return tokenizer.scanString()
	} else if isSymbol(ch) {
		return tokenizer.scanSymbol()
	}
	return nil, nil
}

func (tokenizer *mysqlTokenizer) skip() {
out:
	for tokenizer.offset < tokenizer.length {
		ch := tokenizer.sql[tokenizer.offset]
		switch ch {
		case ' ', '\t', '\f':
			tokenizer.offset++
		case '\n':
			tokenizer.offset++
			tokenizer.line++
		default:
			break out
		}
	}
}

func (tokenizer *mysqlTokenizer) scanIdent() (*dialect.Token, error) {
	flag := false
	if tokenizer.sql[tokenizer.offset] == '`' {
		flag = true
		tokenizer.offset++
	}
	curr := tokenizer.offset
	for tokenizer.offset < tokenizer.length {
		ch := tokenizer.sql[tokenizer.offset]
		slog.Info(fmt.Sprintf("ident: %c", ch))
		if isLetter(ch) {
			tokenizer.offset++
		} else {
			break
		}
	}
	if err := tokenizer.consume('`'); flag && err != nil {
		return nil, err
	}
	text := tokenizer.sql[curr:tokenizer.offset]
	return &dialect.Token{
		Text: text,
		Type: dialect.Ident,

		Position: dialect.Position{
			Offset: curr,
			Line:   0,
			Column: 0,
		},
	}, nil
}

func (tokenizer *mysqlTokenizer) scanNumber() (*dialect.Token, error) {
	curr := tokenizer.offset
	for tokenizer.offset < tokenizer.length {
		ch := tokenizer.sql[tokenizer.offset]
		slog.Info(fmt.Sprintf("number: %c", ch))
		if isNumber(ch) {
			tokenizer.offset++
		} else {
			break
		}
	}
	return &dialect.Token{
		Text: tokenizer.sql[curr:tokenizer.offset],
		Type: dialect.Float,
	}, nil
}

func (tokenizer *mysqlTokenizer) scanString() (*dialect.Token, error) {
	tokenizer.offset++
	curr := tokenizer.offset
	flag := false
	for tokenizer.offset < tokenizer.length {
		ch := tokenizer.sql[tokenizer.offset]
		slog.Info(fmt.Sprintf("string: %c", ch))
		if ch == '\'' {
			tokenizer.offset++
			flag = true
			break
		} else {
			tokenizer.offset++
		}
	}
	if !flag && tokenizer.offset == tokenizer.length {
		return nil, errors.New("expect: ', but got none")
	}
	return &dialect.Token{
		Text: tokenizer.sql[curr:tokenizer.offset],
		Type: dialect.String,
	}, nil
}

func (tokenizer *mysqlTokenizer) scanSymbol() (*dialect.Token, error) {
	mark := tokenizer.offset
	ch := tokenizer.sql[mark]
	tokenizer.offset++
	peek := tokenizer.offset
	for i := mark; i < tokenizer.length; i++ {
	}
	return nil, nil
}

func (tokenizer *mysqlTokenizer) consume(chars ...uint8) error {
	for _, ch := range chars {
		if tokenizer.sql[tokenizer.offset] == ch {
			tokenizer.offset++
			return nil
		} else {
			return errors.New(fmt.Sprintf("expect: %c, actual: %c", ch, tokenizer.sql[tokenizer.offset]))
		}
	}
	return nil
}

func newMysqlTokenizer(sql string) *mysqlTokenizer {
	return &mysqlTokenizer{
		sql:    sql,
		line:   1,
		offset: 0,
		length: len(sql),
	}
}

func isLetter(ch uint8) bool {
	return ch == '_' || ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')
}

func isNumber(ch uint8) bool {
	return '0' <= ch && ch <= '9'
}

func isSymbol(ch uint8) bool {
	for i := op_start; i < op_end; i++ {
		if str, ok := tokens[i]; ok && str[0] == ch {
			return true
		}
	}
	return false
}
