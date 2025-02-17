//
// File: config.go
// Project: v1
// File Created: 2025-01-20
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-20 23:07:14
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

type KdbConfig struct {
	Path    string         `json:"path,omitempty"`
	Pid     string         `json:"pid,omitempty"`
	Server  *ServerConfig  `json:"server"`
	Admin   *AdminConfig   `json:"admin"`
	Plugin  *PluginConfig  `json:"plugin,omitempty"`
	Monitor *MonitorConfig `json:"monitor,omitempty"`
	Trace   *TraceConfig   `json:"trace,omitempty"`
	Metrics *MetricsConfig `json:"metrics,omitempty"`
}

type MonitorConfig struct {
	Enable bool           `json:"enable,omitempty"`
	Host   string         `json:"host,omitempty"`
	Port   int            `json:"port,omitempty"`
	Config map[string]any `json:"config,omitempty"`
}

type TraceConfig struct {
	Enable bool           `json:"enable,omitempty"`
	Driver string         `json:"driver,omitempty"`
	Config map[string]any `json:"config,omitempty"`
}

type MetricsConfig struct {
	Enable bool           `json:"enable,omitempty"`
	Config map[string]any `json:"config,omitempty"`
}

type PluginConfig struct {
	Enable bool           `json:"enable,omitempty"`
	Path   string         `json:"path,omitempty"`
	Config map[string]any `json:"config,omitempty"`
}

func NewKdbConfig(file string) (*KdbConfig, error) {
	return nil, nil
}
