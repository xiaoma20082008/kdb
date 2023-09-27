//
// File: metadata.go
// Project: storage
// File Created: 2023-09-25
// Author: xiaoma20082008 (mmccxx2519@gmail.com)
// -----
// Last Modified By:  xiaoma20082008 (mmccxx2519@gmail.com)
// Last Modified Time: 2023-09-25 19:57:35
// -----
//
// Copyright (C) 2023, xiaoma20082008. All rights reserved.
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

import (
	"fmt"

	"github.com/apache/arrow/go/v14/arrow"
)

type (
	BoolType arrow.BooleanType

	Int8Type   arrow.Int8Type
	Int16Type  arrow.Int16Type
	Int32Type  arrow.Int32Type
	Int64Type  arrow.Int64Type
	UInt8Type  arrow.Int8Type
	UInt16Type arrow.Int16Type
	UInt32Type arrow.Int32Type
	UInt64Type arrow.Int64Type

	VarcharType arrow.StringType

	DateType     arrow.Date32Type
	DateTimeType arrow.Date64Type

	TimeType      arrow.Time32Type
	TimestampType arrow.Time64Type
)

const Catalog = "abc"

type Column interface {
	fmt.Stringer

	// the catalog used. Currently always "def"
	Catalog() string
	// schema name
	Schema() string
	// virtual table name
	Table() string
	// virtual column name
	Name() string
	// maximum length of the field
	Length() int
	// the column character set
	Charset() int16
	// type of the column
	Type() int16
	//
	Flags() int16
	// max shown decimal digits:
	// 	0x00 for integers and static strings
	// 	0x1f for dynamic strings, double, float
	// 	0x00 to 0x51 for decimals
	Decimal() int8
}

type Columns []Column

type Table interface {
	fmt.Stringer
	Catalog() string
	Schema() string
	Name() string
	Columns() Columns
}

type column struct {
	catalog    string
	schemaName string
	tableName  string
	columnName string
	length     int
	charset    int16
	columnType int16
	flags      int16
	decimal    int8
}

type table struct {
	catalog    string
	schemaName string
	name       string
	columns    Columns
}

func (c *column) Catalog() string { return c.catalog }
func (c *column) Schema() string  { return c.schemaName }
func (c *column) Table() string   { return c.tableName }
func (c *column) Name() string    { return c.columnName }
func (c *column) Length() int     { return c.length }
func (c *column) Charset() int16  { return c.charset }
func (c *column) Type() int16     { return c.columnType }
func (c *column) Flags() int16    { return c.flags }
func (c *column) Decimal() int8   { return c.decimal }
func (c *column) String() string  { return "" }

func (t *table) Catalog() string  { return t.catalog }
func (t *table) Schema() string   { return t.schemaName }
func (t *table) Name() string     { return t.name }
func (t *table) Columns() Columns { return t.columns }
func (t *table) String() string   { return "" }

type MetadataProvider interface {
	Dialect() string
	AddTable(name string, table Table)
	GetTable(name string) Table
}

type metadataProvider struct {
	dialect string

	tables map[string]Table
}

func (mp *metadataProvider) Dialect() string {
	return mp.dialect
}

func (mp *metadataProvider) AddTable(name string, t Table) {
	mp.tables[name] = t
}

func (mp *metadataProvider) GetTable(name string) Table {
	if tb, ok := mp.tables[name]; ok {
		return tb
	}
	return nil
}

func NewColumn(schema, table, name string, length int, charset, ctype, flags int16, decimal int8) Column {
	return &column{
		catalog:    Catalog,
		schemaName: schema,
		tableName:  table,
		columnName: name,
		length:     length,
		charset:    charset,
		columnType: ctype,
		flags:      flags,
		decimal:    decimal,
	}
}

func NewTable(schema, name string, columns ...Column) Table {
	return &table{
		catalog:    Catalog,
		schemaName: schema,
		name:       name,
		columns:    columns,
	}
}

func NewMetadataProvider() MetadataProvider {
	return &metadataProvider{
		tables: make(map[string]Table),
	}
}
