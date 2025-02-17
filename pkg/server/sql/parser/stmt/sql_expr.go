//
// File: sql_expr.go
// Project: stmt
// File Created: 2025-01-05
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-05 18:17:50
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

package stmt

import (
	"context"
	"fmt"
)

type SqlExpr interface {
	fmt.Stringer
	Accept(v SqlVisitor, ctx context.Context) interface{}
}

type SqlVisitor interface {
	Visit(SqlExpr) bool
}

type SqlIdent struct {
	SqlExpr
	Name string
}

type SqlValue struct {
	SqlExpr
	Value string
}

type SqlBinaryExpr struct {
	SqlExpr
	Lhs SqlExpr
	Rhs SqlExpr
}

type SqlUnaryExpr struct {
	SqlExpr
	Rhs SqlExpr
}

type SqlLikeExpr struct {
	SqlExpr
	Ident SqlIdent
	Regex SqlValue
}

type SqlBetweenExpr struct {
	SqlExpr
	Ident SqlIdent
	Regex SqlValue
}
