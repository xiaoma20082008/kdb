//
// File: scanner.go
// Project: scanner
// File Created: 2025-01-15
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-15 22:09:38
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
	"github.com/xiaoma20082008/kdb/pkg/server/sql/tokens"
)

const EOI byte = '\u0000'

type Scanner struct {
	reader *stringReader

	line   uint32
	column uint32
	read   string
}

func (s *Scanner) Next() (*tokens.Token, error) {
	c := s.peek()
	if s.isSpace(c) {
		return s.skipWhitespace()
	} else if s.isChar(c) {
		return s.readIdent()
	} else if s.isDigit(c) {
		return s.readNumber()
	} else if c == '\'' {
		return s.readString()
	} else {
		return s.make(0, s.read, s.line, s.column), nil
	}
}

func (s *Scanner) advance() byte {
	ch := s.reader.current()
	s.reader.advance()

	s.read += string(ch)

	if ch == '\n' {
		s.line++
		s.column = 1
	} else {
		s.column++
	}
	return ch
}

func (s *Scanner) peek() byte {
	return s.reader.current()
}

func (s *Scanner) skipWhitespace() (*tokens.Token, error) {
	for {
		ch := s.peek()
		if s.isSpace(ch) {
			s.advance()
		} else {
			s.read = ""
			return s.Next()
		}
	}
}

func (s *Scanner) readString() (*tokens.Token, error) {
	s.advance()
	for {
		ch := s.peek()
		if ch != '\'' {
			s.advance()
		} else {
			return s.make(0, s.read[1:len(s.read)], s.line, s.column), nil
		}
	}
}

func (s *Scanner) readIdent() (*tokens.Token, error) {
	for {
		ch := s.peek()
		if s.isDigit(ch) || s.isChar(ch) {
			s.advance()
		} else {
			return s.make(0, s.read, s.line, s.column), nil
		}
	}
}

func (s *Scanner) readBinNumber() (*tokens.Token, error) {
	for {
		ch := s.peek()
		if s.isBin(ch) {
			s.advance()
		} else {
			return s.make(0, s.read, s.line, s.column), nil
		}
	}
}

func (s *Scanner) readOctNumber() (*tokens.Token, error) {
	for {
		ch := s.peek()
		if s.isOct(ch) {
			s.advance()
		} else {
			return s.make(0, s.read, s.line, s.column), nil
		}
	}
}

func (s *Scanner) readHexNumber() (*tokens.Token, error) {
	for {
		ch := s.peek()
		if s.isHex(ch) {
			s.advance()
		} else {
			return s.make(0, s.read, s.line, s.column), nil
		}
	}
}

func (s *Scanner) readNumber() (*tokens.Token, error) {
	c := s.advance()
	flag := false
	if c == '0' {
		e := s.advance()
		if e == 'b' || e == 'B' {
			return s.readBinNumber()
		} else if s.isOct(e) {
			return s.readOctNumber()
		} else if e == 'x' || e == 'X' {
			return s.readHexNumber()
		}
	}
	for {
		if s.isDigit(s.peek()) {
			s.advance()
		} else if s.peek() == '.' && !flag {
			flag = true
			s.advance()
		} else {
			return s.make(0, s.read, s.line, s.column), nil
		}
	}
}

func (s *Scanner) make(kind int, val string, ln, col uint32) *tokens.Token {
	s.read = ""
	return tokens.New(0, val, ln, col)
}

func (s *Scanner) isChar(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_' || ch == '$'
}

func (s *Scanner) isBin(ch byte) bool {
	return '0' <= ch && ch <= '1'
}

func (s *Scanner) isOct(ch byte) bool {
	return '0' <= ch && ch <= '7'
}

func (s *Scanner) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (s *Scanner) isHex(ch byte) bool {
	return ('0' <= ch && ch <= '9') || ('a' <= ch && ch <= 'f') || ('A' <= ch && ch <= 'F')
}

func (s *Scanner) isSpace(ch byte) bool {
	return ch == ' ' || ch == '\r' || ch == '\n' || ch == '\t' || ch == '\f'
}

func New(text string) *Scanner {
	s := new(Scanner)
	s.reader = newReader(text)
	s.line = 1
	s.column = 1
	s.read = ""
	return s
}
