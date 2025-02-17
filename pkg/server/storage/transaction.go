//
// File: transaction.go
// Project: storage
// File Created: 2025-02-06
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
//
// ------------------------------------------------------------------------
// Last Modified At: 2025-02-06 20:15:12
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

package storage

import "io"

// 事务状态
type TransactionStatus int

const (
	TransactionNone TransactionStatus = iota
	TransactionActive
	TransactionCommitting
	TransactionCommitted
	TransactionAborted
)

func (status TransactionStatus) String() string {
	switch status {
	case TransactionActive:
		return "Active"
	case TransactionCommitting:
		return "Committing"
	case TransactionCommitted:
		return "Committed"
	case TransactionAborted:
		return "Aborted"
	default:
		return "None"
	}
}

// 事务隔离级别
type IsolationLevel int

const (
	Unknown IsolationLevel = iota
	ReadUncommitted
	ReadCommitted
	RepeatableRead
	Serializable
)

func (level IsolationLevel) String() string {
	switch level {
	case ReadUncommitted:
		return "READ-UNCOMMITTED"
	case ReadCommitted:
		return "READ-COMMITTED"
	case RepeatableRead:
		return "REPEATABLE-READ"
	case Serializable:
		return "SERIALIZABLE"
	default:
		return "UNKNOWN"
	}
}

type Transaction interface {
	Id() uint64

	Status() TransactionStatus

	Commit() error

	Rollback() error

	Stats() map[string]interface{}
}

type TransactionManager interface {
	io.Closer

	Current() Transaction

	Begin() (Transaction, error)

	Commit() error

	Rollback() error
}
