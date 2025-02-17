//
// File: location.go
// Project: tokens
// File Created: 2025-01-18
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-18 01:02:48
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

type Location struct {
	fmt.Stringer
	line   uint32
	column uint32
}

func (l *Location) Line() uint32 {
	return l.line
}

func (l *Location) Column() uint32 {
	return l.column
}

func (l *Location) String() string {
	return fmt.Sprintf("Line: %d, Column: %d", l.line, l.column)
}

func Position(ln, col uint32) *Location {
	location := new(Location)
	location.line = ln
	location.column = col
	return location
}
