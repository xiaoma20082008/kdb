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
}

type SqlStmt interface {
	SqlNode
}

type SqlExpr interface {
	SqlNode
}

type SqlExprList interface {
	SqlExpr

	List() []SqlExpr
}

// SqlStmt
type (

	// select * from t1,t2 where xx = ? group by xx having xx = xx order by xx limit
	SqlSelect struct {
		SqlStmt

		Columns SqlExprList
		From    SqlExprList
		Where   SqlExpr
		GroupBy SqlExprList
		Having  SqlExpr
		OrderBy SqlExprList
		Limit   SqlExpr
		Offset  SqlExpr
	}

	SqlInsert struct {
		SqlStmt

		Table   *SqlIdentifier
		Columns SqlExprList
		Values  SqlExprList
	}

	SqlDelete struct {
		SqlStmt

		Table *SqlIdentifier
		Where SqlExpr

		OrderBy SqlExpr
		Limit   SqlExpr
	}

	SqlUpdate struct {
		SqlStmt

		Table *SqlIdentifier
		Items SqlUpdateItems
		Where SqlExpr

		OrderBy SqlExpr
		Limit   SqlExpr
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

	SqlInteger struct {
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

type SqlUpdateItems []*SqlUpdateItem
