//
// File: analyzer.go
// Project: analyzer
// File Created: 2023-09-07
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-07 18:21:39
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

package analyzer

import (
	"fmt"
	"kdb/dialect"
	"kdb/storage"
)

type analyzer struct {
	dialect.Visitor

	Metadata storage.MetadataProvider
	Analysis Analysis
	Error    error
}

func (a *analyzer) Visit(node dialect.SqlNode) bool {
	switch n := node.(type) {
	case *dialect.SqlSelect:
		return a.VisitSelect(n)
	}
	return false
}

func (a *analyzer) VisitSelect(n *dialect.SqlSelect) bool {
	//
	for _, tb := range n.From.List() {
		if a.Metadata.GetTable(tb.String()) == nil {
			a.Error = fmt.Errorf("table not found")
			return false
		}
	}
	return true
}

func Analyze(ast dialect.SqlStmt) (*Analysis, error) {
	analyzer := analyzer{}
	analyzer.Visit(ast)
	if analyzer.Error != nil {
		return nil, analyzer.Error
	}
	return &Analysis{
		Ast: ast,
	}, nil
}
