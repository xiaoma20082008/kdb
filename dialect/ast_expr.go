//
// File: ast_expr.go
// Project: dialect
// File Created: 2023-09-15
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-15 19:50:46
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

func (identifier SqlIdentifier) String() string {
	return identifier.Id
}

func (identifier SqlString) String() string {
	return fmt.Sprintf("'%s'", identifier.Value)
}

func (identifier SqlInteger) String() string {
	return fmt.Sprintf("%d", identifier.Value)
}

func (identifier SqlFloat) String() string {
	return fmt.Sprintf("%f", identifier.Value)
}

func (identifier SqlFunc) String() string {
	return identifier.Name
}

func (identifier SqlBinaryExpr) String() string {
	return identifier.Lhs.String() + " " + identifier.Op.Text + " " + identifier.Rhs.String()
}

func (identifier SqlAlias) String() string {
	return identifier.Expr.String() + " AS " + identifier.Alias.Id
}

func (identifier SqlSort) String() string {
	if identifier.Desc {
		return identifier.Expr.String() + " DESC"
	} else {
		return identifier.Expr.String() + " ASC"
	}
}

func (identifier SqlCast) String() string {
	return ""
}

func (identifier SqlUpdateItem) String() string {
	return identifier.Key.String() + " = " + identifier.Val.String()
}
