//
// File: mysql_parser.go
// Project: mysql
// File Created: 2023-09-21
// Author: machunxiao (machunxiao@shizhuang-inc.com)
// -----
// Last Modified By:  machunxiao (machunxiao@shizhuang-inc.com)
// Last Modified Time: 2023-09-21 20:16:34
// -----
//
// Copyright (C) 2023, Shanghai Poizon Information Technology Co., Ltd.
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

import "kdb/dialect"

type mysqlParser struct {
	offset int
	stream *dialect.TokenList
}

func (parser *mysqlParser) parse() (dialect.SqlStmt, error) {
	return nil, nil
}

func newMysqlParser(stream *dialect.TokenList) *mysqlParser {
	return &mysqlParser{
		offset: 0,
		stream: stream,
	}
}
