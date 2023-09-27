package logicalplan

import (
	"kdb/storage"
)

type Scan struct {
	LogicalPlan

	Table      storage.Table
	Projection LogicalExprList
	Condition  LogicalExpr
}

func (s *Scan) GetTable() storage.Table {
	return s.Table
}

func (s *Scan) GetChildren() []LogicalPlan {
	return make([]LogicalPlan, 0)
}

func (s *Scan) String() string {
	return "Scan: " + s.Table.Name() + ""
}

func NewScan(path string, mp storage.MetadataProvider) LogicalPlan {
	tb := mp.GetTable(path)
	return &Scan{
		Table: tb,
	}
}
