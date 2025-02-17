//
// File: jdbc_conn.go
// Project: network
// File Created: 2025-01-22
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-22 21:33:51
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

package network

import (
	"net"

	_ "github.com/xiaoma20082008/kdb/pkg/server/network/jdbc/mssql"
	_ "github.com/xiaoma20082008/kdb/pkg/server/network/jdbc/mysql"
	_ "github.com/xiaoma20082008/kdb/pkg/server/network/jdbc/oracle"
)

type jdbcConnection struct {
	Connection
	id  string
	raw net.Conn
}

func (jdbc *jdbcConnection) Id() string { return jdbc.id }

func (jdbc *jdbcConnection) Handshake() error { return nil }

func (jdbc *jdbcConnection) serve() error {
	defer func() {
		if err := recover(); err != nil {
			jdbc.raw.Close()
		}
	}()
	jdbc.Handshake()

	return nil
}
