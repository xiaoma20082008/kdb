//
// File: ast.go
// Project: dialect
// File Created: 2023-09-15
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-15 19:39:21
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

import "fmt"

type SqlNode interface {
	fmt.Stringer
	Position() Position
}

type SqlStmt interface {
	SqlNode
}

type SqlExpr interface {
	SqlNode
}

// SqlStmt
type (
	SqlSelect struct {
		SqlStmt
	}

	SqlInsert struct {
		SqlStmt

		table   *SqlIdentifier
		columns []SqlExpr
		values  []SqlExpr
	}

	SqlDelete struct {
		SqlStmt

		table *SqlIdentifier
		where SqlExpr

		orderBy SqlExpr
		limit   SqlExpr
	}

	SqlUpdate struct {
		SqlStmt

		table *SqlIdentifier
		items []*SqlUpdateItem
		where SqlExpr

		orderBy SqlExpr
		limit   SqlExpr
	}
)

// SqlExpr
type (
	SqlIdentifier struct {
		SqlExpr
		Id string
	}

	SqlString struct {
		SqlExpr
		Value string
	}

	SqlLong struct {
		SqlExpr
		Value int64
	}

	SqlFloat struct {
		SqlExpr
		Value float64
	}

	SqlFunc struct {
		SqlExpr

		Name string
		Args []SqlExpr
	}

	SqlBinaryExpr struct {
		SqlExpr

		Lhs SqlExpr
		Op  Token
		Rhs SqlExpr
	}

	SqlAlias struct {
		SqlExpr

		Expr  SqlExpr
		Alias SqlIdentifier
	}

	SqlSort struct {
		SqlExpr

		Expr SqlExpr
		Desc bool
	}

	SqlCast struct {
		SqlExpr

		Expr SqlExpr
		Cast SqlIdentifier
	}

	SqlUpdateItem struct {
		SqlExpr

		Key SqlExpr
		Val SqlExpr
	}

	SqlHint struct {
		SqlExpr

		Hint string
	}

	SqlComment struct {
		SqlExpr

		Comment string
	}
)
