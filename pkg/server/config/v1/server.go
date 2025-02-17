//
// File: server.go
// Project: v1
// File Created: 2025-01-20
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-20 23:22:15
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

package v1

type ServerConfig struct {
	Host      string           `json:"host,omitempty"`
	Port      int              `json:"port,omitempty"`
	Session   *SessionConfig   `json:"session,omitempty"`
	Parser    *ParserConfig    `json:"parser,omitempty"`
	Optimizer *OptimizerConfig `json:"optimizer,omitempty"`
	Rule      *RuleConfig      `json:"rule,omitempty"`
	Executor  *ExecutorConfig  `json:"executor,omitempty"`
	Storage   *StorageConfig   `json:"storage,omitempty"`
}

type SessionConfig struct {
	ConnectTimeout int  `json:"connect_timeout,omitempty"`
	ReadTimeout    int  `json:"read_timeout,omitempty"`
	WriteTimeout   int  `json:"write_timeout,omitempty"`
	ExpireTimeout  int  `json:"expire_timeout,omitempty"`
	MaxIdle        int  `json:"max_idle,omitempty"`
	MaxConnections int  `json:"max_connections,omitempty"`
	IsolationLevel int  `json:"isolation_level,omitempty"`
	AutoCommit     bool `json:"auto_commit,omitempty"`
	UseSSL         bool `json:"use_ssl,omitempty"`
	//
	Timezone string         `json:"timezone,omitempty"`
	Charset  string         `json:"charset,omitempty"`
	Extended map[string]any `json:"extended,omitempty"`
}

type ParserConfig struct {
}

type OptimizerConfig struct {
}

type RuleConfig struct {
}

type ExecutorConfig struct {
}
