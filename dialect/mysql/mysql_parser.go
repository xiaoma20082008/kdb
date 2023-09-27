//
// File: mysql_parser.go
// Project: mysql
// File Created: 2023-09-21
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-21 20:16:34
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
	"fmt"
	"kdb/dialect"
	"strconv"
)

type mysqlParser struct {
	offset int
	stream *dialect.TokenList
}

func (parser *mysqlParser) parseStmt() (dialect.SqlStmt, error) {
	token := parser.stream.Curr()
	if token.Type == dialect.Symbol {
		switch token.Kind {
		case INSERT:
			return parser.parseInsert()
		case DELETE:
			return parser.parseDelete()
		case UPDATE:
			return parser.parseUpdate()
		case SELECT:
			return parser.parseSelect()
		default:
			return nil, fmt.Errorf("unknown TokenKind: %d", token.Kind)
		}

	}
	return nil, fmt.Errorf("expect: insert/delete/update/select, but got: %v", token)
}

func (parser *mysqlParser) parseExpr() (dialect.SqlExpr, error) {
	// https://dev.mysql.com/doc/refman/8.0/en/expressions.html#expression-syntax
	token := parser.stream.Curr()
	if token == nil {
		return nil, nil
	}
	switch token.Type {
	case dialect.Ident:
		return &dialect.SqlIdentifier{Id: token.Text}, nil
	case dialect.Symbol:
		// TODO
		return nil, nil
	case dialect.Float:
		v, _ := strconv.ParseFloat(token.Text, 64)
		return &dialect.SqlFloat{Value: v}, nil
	case dialect.Integer:
		v, _ := strconv.ParseInt(token.Text, 10, 64)
		return &dialect.SqlInteger{Value: v}, nil
	case dialect.String:
		return &dialect.SqlString{Value: token.Text}, nil
	case dialect.Comment:
		return &dialect.SqlComment{Comment: token.Text}, nil
	case dialect.EOF:
		return nil, nil
	}
	return nil, nil
}

func (parser *mysqlParser) parseInsert() (*dialect.SqlInsert, error) {
	// insert into t(a,b,c) values()
	if err := parser.consume(INSERT); err != nil {
		return nil, err
	}
	if err := parser.consume(INTO); err != nil {
		return nil, err
	}
	tb, err := parser.parseExpr()
	if err != nil {
		return nil, err
	}
	table := tb.(dialect.SqlIdentifier)
	if err := parser.consume(LP); err != nil {
		return nil, err
	}
	columns, err := parser.parseExprList()
	if err != nil {
		return nil, err
	}
	if err := parser.consume(RP); err != nil {
		return nil, err
	}
	if err := parser.consume(VALUES); err != nil {
		return nil, err
	}
	if err := parser.consume(LP); err != nil {
		return nil, err
	}
	values, err := parser.parseExprList()
	if err != nil {
		return nil, err
	}
	if err := parser.consume(RP); err != nil {
		return nil, err
	}
	return &dialect.SqlInsert{
		Table:   &table,
		Columns: columns,
		Values:  values,
	}, nil
}

func (parser *mysqlParser) parseDelete() (*dialect.SqlDelete, error) {
	// delete from t1 where xxx
	if err := parser.consume(DELETE); err != nil {
		return nil, err
	}
	if err := parser.consume(FROM); err != nil {
		return nil, err
	}
	tb, err := parser.parseExpr()
	if err != nil {
		return nil, err
	}
	table := tb.(dialect.SqlIdentifier)
	var where dialect.SqlExpr
	if parser.stream.Peek().Kind == WHERE {
		parser.consume(WHERE)
		where, err = parser.parseExpr()
		if err != nil {
			return nil, err
		}
	}
	var orderBy dialect.SqlExpr
	if parser.stream.Peek().Kind == ORDER {
		parser.consume(ORDER)
		parser.consume(BY)
		orderBy, err = parser.parseExpr()
		if err != nil {
			return nil, err
		}
	}
	var limit dialect.SqlExpr
	if parser.stream.Peek().Kind == LIMIT {
		parser.consume(LIMIT)
		limit, err = parser.parseExpr()
		if err != nil {
			return nil, err
		}
	}
	return &dialect.SqlDelete{
		Table:   &table,
		Where:   where,
		OrderBy: orderBy,
		Limit:   limit,
	}, nil
}

func (parser *mysqlParser) parseUpdate() (*dialect.SqlUpdate, error) {
	if err := parser.consume(UPDATE); err != nil {
		return nil, err
	}
	tb, err := parser.parseExpr()
	if err != nil {
		return nil, err
	}
	table := tb.(dialect.SqlIdentifier)
	if err := parser.consume(SET); err != nil {
		return nil, err
	}
	return &dialect.SqlUpdate{
		Table: &table,
	}, nil
}

func (parser *mysqlParser) parseSelect() (*dialect.SqlSelect, error) {
	// select <select_expr>
	// from <table_ref>
	// where <condition>
	// group by <expr>
	// having <condition>
	// order by <order_expr>
	// limit <offset>, <limit>

	// select
	if err := parser.consume(SELECT); err != nil {
		return nil, err
	}
	columns, err := parser.parseExprList()
	if err != nil {
		return nil, err
	}

	// from
	if err := parser.consume(FROM); err != nil {
		return nil, err
	}
	from, err := parser.parseExprList()
	if err != nil {
		return nil, err
	}

	// where
	var where dialect.SqlExpr
	if parser.stream.Peek().Kind == WHERE {
		parser.consume(WHERE)
		where, err = parser.parseExpr()
		if err != nil {
			return nil, err
		}
	}

	// group by
	var groupBy *dialect.SqlExprList
	if parser.stream.Peek().Kind == GROUP {
		parser.consume(GROUP)
		if err := parser.consume(BY); err != nil {
			return nil, err
		}
		groupBy, err = parser.parseExprList()
		if err != nil {
			return nil, err
		}
	}

	// having
	var having dialect.SqlExpr
	if parser.stream.Peek().Kind == HAVING {
		parser.consume(HAVING)
		having, err = parser.parseExpr()
		if err != nil {
			return nil, err
		}
	}

	// order by
	var orderBy *dialect.SqlExprList
	if parser.stream.Peek().Kind == ORDER {
		parser.consume(ORDER)
		if err := parser.consume(BY); err != nil {
			return nil, err
		}
		orderBy, err = parser.parseExprList()
		if err != nil {
			return nil, err
		}
	}

	// limit xx,xx
	var limit, offset dialect.SqlExpr
	if parser.stream.Peek().Kind == LIMIT {
		parser.consume(LIMIT)
		p, err := parser.parseExpr()
		if err != nil {
			return nil, err
		}
		if parser.stream.Peek().Kind == COMMA {
			parser.consume(COMMA)
			limit, err = parser.parseExpr()
			if err != nil {
				return nil, err
			}
			offset = p
		} else {
			limit = p
		}
	}
	return &dialect.SqlSelect{
		Columns: columns,
		From:    from,
		Where:   where,
		GroupBy: groupBy,
		Having:  having,
		OrderBy: orderBy,
		Limit:   limit,
		Offset:  offset,
	}, nil
}

func (parser *mysqlParser) parseExprList() (*dialect.SqlExprList, error) {
	// expr [',' expr]*
	return nil, nil
}

func (parser *mysqlParser) consume(kind dialect.TokenKind) error {
	token := parser.stream.Curr()
	if token != nil && token.Kind == kind {
		parser.stream.Next()
		return nil
	}
	return fmt.Errorf("expect: %d, but got: %v", kind, token)
}

func newMysqlParser(stream *dialect.TokenList) *mysqlParser {
	return &mysqlParser{
		offset: 0,
		stream: stream,
	}
}
