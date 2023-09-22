//
// File: dialect.go
// Project: dialect
// File Created: 2023-09-15
// Author: machunxiao (machunxiao@shizhuang-inc.com)
// -----
// Last Modified By:  machunxiao (machunxiao@shizhuang-inc.com)
// Last Modified Time: 2023-09-15 18:07:20
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

package dialect

type Dialect int32

const (
	DB2 Dialect = iota
	MySQL
	Oracle
	PostgreSQL
	SqlServer
	Hive
	SQLite
	MariaDB
)

func (d Dialect) String() string {
	switch d {
	case DB2:
		return "IBM DB2"
	case MySQL:
		return "MySQL"
	case Oracle:
		return "Oracle"
	case PostgreSQL:
		return "PostgreSQL"
	case SqlServer:
		return "Microsoft SQL Server"
	case Hive:
		return "Hive"
	case SQLite:
		return "SQLite"
	case MariaDB:
		return "MariaDB"
	}
	return "UNKNOWN"
}

var Tokenize func(sql string) (*TokenList, error)

var Parse func(tokens *TokenList) (SqlStmt, error)
