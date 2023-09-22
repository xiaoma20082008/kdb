package logicalplan

import (
	"kdb/datasource"
	"kdb/datatype"
)

type Scan struct {
	LogicalPlan

	path string

	dataSource datasource.DataSource
}

func (s *Scan) GetTable() *datatype.Table {
	return s.dataSource.GetTable()
}

func (s *Scan) GetChildren() []LogicalPlan {
	return make([]LogicalPlan, 0)
}

func (s *Scan) String() string {
	return "Scan: " + s.path
}

func NewScan(path string, dataSource datasource.DataSource) LogicalPlan {
	f := new(Scan)
	f.path = path
	f.dataSource = dataSource
	return f
}
