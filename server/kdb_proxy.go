//
// File: kdb_proxy.go
// Project: server
// File Created: 2023-09-14
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-14 21:45:55
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

package server

import (
	"fmt"
	"kdb/conf"
)

type KdbProxy struct {
	config *conf.ServerConf
	admin  *KdbAdmin
	server *kdbServer
}

func (proxy *KdbProxy) Start() error {
	if err := proxy.server.Start(); err != nil {
		return nil
	}
	if err := proxy.admin.Start(); err != nil {
		return nil
	}
	return nil
}

func (proxy *KdbProxy) Close() error {
	if err := proxy.server.Close(); err != nil {
		fmt.Println(err)
	}
	if err := proxy.admin.Close(); err != nil {
		fmt.Println(err)
	}
	return nil
}

func (proxy *KdbProxy) Await() {
	proxy.server.Await()
	proxy.admin.Await()
}

func CreateProxy(conf *conf.ServerConf) *KdbProxy {
	proxy := new(KdbProxy)
	proxy.config = conf
	proxy.admin = createAdmin(conf)
	proxy.server = createServer(conf)
	return proxy
}
