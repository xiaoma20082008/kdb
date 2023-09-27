//
// File: rewriter.go
// Project: rewriter
// File Created: 2023-09-26
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-26 17:17:17
// -----
//
// Copyright (C) 2023, xiaoma20082008. All rights reserved.
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

package rewriter

import (
	"kdb/analyzer"
	"kdb/dialect"
	"kdb/logicalplan"
	"kdb/storage"
)

type rewriter struct {
	dialect.Visitor

	mp    storage.MetadataProvider
	frame logicalplan.DataFrame
	err   error
}

func (r *rewriter) Visit(node dialect.SqlNode) bool {
	switch n := node.(type) {
	case *dialect.SqlSelect:
		return r.visitSelect(n)
	}
	return false
}

func (r *rewriter) visitSelect(sqlSelect *dialect.SqlSelect) bool {
	// 这里我们以单表为例
	tb := sqlSelect.From.List()[0].(dialect.SqlIdentifier)

	// table scan
	scan := logicalplan.NewScan(tb.Id, r.mp)
	df := logicalplan.NewDataFrame(scan)

	// filter
	filter, err := createLogicalExpr(sqlSelect.Where)
	if err != nil {
		r.err = err
		return false
	}
	r.frame = df.Filter(filter)
	return false
}

func createLogicalExpr(expr dialect.SqlExpr) (logicalplan.LogicalExpr, error) {
	return nil, nil
}

func Rewrite(analysis *analyzer.Analysis) (logicalplan.DataFrame, error) {
	r := rewriter{}
	r.Visit(analysis.Ast)
	if r.err != nil {
		return nil, r.err
	}
	return r.frame, nil
}
