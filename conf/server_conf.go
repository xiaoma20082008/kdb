//
// File: server_conf.go
// Project: conf
// File Created: 2023-09-14
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-14 11:49:24
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

package conf

type ServerConf struct {
	MainName string `default:"kdb" yaml:"mainName" json:"mainName"`
	MainPort int32  `default:"3306" yaml:"mainPort" json:"mainPort"`
	CtrlPort int32  `default:"8081" yaml:"ctrlPort" json:"ctrlPort"`

	ReadTimeout  int32 `default:"5" yaml:"readTimeout" json:"readTimeout"`
	WriteTimeout int32 `default:"10" yaml:"writeTimeout" json:"writeTimeout"`
	CloseTImeout int32 `default:"10" yaml:"closeTImeout" json:"closeTImeout"`
}
