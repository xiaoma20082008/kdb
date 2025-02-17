//
// File: kdb_server.go
// Project: proxy
// File Created: 2025-01-20
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-20 22:28:58
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

package proxy

import (
	"io"
	"os"

	config "github.com/xiaoma20082008/kdb/pkg/server/config/v1"
)

type KdbServer struct {
	io.Closer
	Config *config.KdbConfig
	Cache  any
	Proxy  *KdbProxy
	Pid    int
}

func (srv *KdbServer) InitConfig() {
}

func (srv *KdbServer) LoadPlugins() error {
	return nil
}

func (srv *KdbServer) InitPlugins() error {
	return nil
}

func (srv *KdbServer) InitSignals() {
}

func (srv *KdbServer) InitMonitors() error {
	return nil
}

func (srv *KdbServer) InitCache() {
}

func (srv *KdbServer) InitProxy() error {
	return nil
}

func (srv *KdbServer) InitStorage() error {
	return nil
}

func (srv *KdbServer) Start() error {
	return nil
}

func (srv *KdbServer) Close() error {
	return nil
}

func NewKdbServer(cfg *config.KdbConfig) *KdbServer {
	srv := new(KdbServer)
	srv.InitConfig()
	srv.InitCache()
	srv.Pid = os.Getpid()
	return srv
}
