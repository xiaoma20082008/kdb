//
// File: url.go
// Project: network
// File Created: 2025-01-20
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-01-20 22:07:43
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

import "fmt"

type URL struct {
	fmt.Stringer
	protocol string
	username string
	password string
	hostname string
	port     int32
	path     string
	params   map[string]string
}

func Url() *URL {
	return &URL{}
}

func (url *URL) Protocol() string { return url.protocol }

func (url *URL) Username() string { return url.username }

func (url *URL) Password() string { return url.password }

func (url *URL) Hostname() string { return url.hostname }

func (url *URL) Port() int32 { return url.port }

func (url *URL) Path() string { return url.path }

func (url *URL) Params() map[string]string { return url.params }

func (url *URL) String() string {
	var out string
	if url.protocol != "" {
		out += url.protocol
		out += "://"
	}
	if url.username != "" {
		out += url.username
		if url.password != "" {
			out += ":"
			out += url.password
		}
		out += "@"
	}
	out += url.hostname
	out += ":"
	out += string(url.port)
	out += "/"
	out += url.path
	return out
}
