//
// File: sql_stmt.go
// Project: stmt
// File Created: 2025-01-05
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-05 18:19:23
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

type SqlStmt interface {
	SqlExpr
}

// dml

type SqlInsert interface {
	SqlStmt
}

type SqlDelete interface {
	SqlStmt
}

type SqlUpdate interface {
	SqlStmt
}

type SqlUpsert interface {
	SqlStmt
}

// dql

type SqlSelect interface {
	SqlStmt
}

// ddl

type SqlCreateDatabase interface {
	SqlStmt
}

type SqlCreateTable interface {
	SqlStmt
}

type SqlCreateIndex interface {
	SqlStmt
}

type SqlAlterTable interface {
	SqlStmt
}

type SqlDropTable interface {
	SqlStmt
}

type SqlDropDatabase interface {
	SqlStmt
}

// dcl

type SqlGrantStmt interface {
	SqlStmt
}

type SqlRevokeStmt interface {
	SqlStmt
}
