//
// File: reader.go
// Project: scanner
// File Created: 2025-01-15
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-15 22:09:51
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

type stringReader struct {
	stream string
	offset int
	line   int
	column int
	limit  int
}

func newReader(text string) *stringReader {
	sr := new(stringReader)
	sr.stream = text
	sr.offset = 0
	sr.line = 1
	sr.column = 1
	sr.limit = len(text)
	return sr
}

func (r *stringReader) available() bool {
	return r.offset < r.limit
}

func (r *stringReader) advance() {
	if r.available() {
		r.offset++
	}
}

func (r *stringReader) current() byte {
	if r.available() {
		return r.stream[r.offset]
	} else {
		return EOI
	}
}

func (r *stringReader) peek(n int) byte {
	if r.offset+n < r.limit {
		return r.stream[r.offset+n]
	} else {
		return EOI
	}
}
