//
// File: executor.go
// Project: executor
// File Created: 2023-09-08
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-08 11:05:55
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

package executor

import (
	"fmt"

	"kdb/analyzer"
	"kdb/dialect"
	_ "kdb/dialect/mysql"
	"kdb/optimizer"
	"kdb/planner"
)

func Execute(sql string, args map[int32]any) error {
	// 1. 词法分析
	tokens, err := dialect.Tokenize(sql)
	if err != nil {
		return err
	}
	// 2. 语法解析
	ast, err := dialect.Parse(tokens)
	if err != nil {
		return err
	}
	// 3. 语义分析
	analysis, err := analyzer.Analyze(ast)
	if err != nil {
		return err
	}
	// 4. 创建逻辑执行计划
	dataFrame, err := planner.Plan(analysis)
	if err != nil {
		return err
	}
	// 5. 优化逻辑执行计划
	logicalPlan := optimizer.Optimize(dataFrame.Plan())
	// 6. 创建物理执行计划
	physicalPlan, err := planner.CreatePhysicalPlan(logicalPlan)
	if err != nil {
		return err
	}
	// 7. 物理计划执行
	if err := physicalPlan.Open(); err != nil {
		return err
	}
	for physicalPlan.HasNext() {
		data := physicalPlan.Next()
		fmt.Println(data)
	}
	return physicalPlan.Close()

}
