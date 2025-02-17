//
// File: ini.go
// Project: v1
// File Created: 2025-01-21
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-21 23:58:58
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

package common

import (
	"fmt"
	"strconv"
	"strings"
)

type Ini struct {
	fmt.Stringer
	conf map[string]string
}

func (ini *Ini) Get(key, def string) string {
	if val, ok := ini.conf[key]; ok {
		return val
	} else {
		return def
	}
}

func (ini *Ini) GetInt(key string, def int) int {
	if val, err := strconv.Atoi(ini.Get(key, "0")); err != nil {
		return val
	} else {
		return def
	}
}

func (ini *Ini) GetBool(key string, def bool) bool {
	val := ini.Get(key, "")
	switch strings.ToLower(val) {
	case "1", "on", "yes", "enable", "enabled":
		return true
	case "0", "off", "no", "disable", "disabled":
		return false
	}
	return def
}

func (ini *Ini) Put(key, val string) {
	ini.conf[key] = val
}

func (ini *Ini) String() string {
	var b strings.Builder
	for k, v := range ini.conf {
		b.WriteString(k)
		b.WriteString("=")
		b.WriteString(v)
		b.WriteString("\n")
	}
	return b.String()
}

func NewIni() *Ini {
	ini := new(Ini)
	ini.conf = make(map[string]string)
	return ini
}
