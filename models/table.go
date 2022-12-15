package models

type Column struct {
	Name     string
	DataType string
}

type Table struct {
	Name             string
	Columns          []*Column
	SequenceQuery    string
	CreateTableQuery string
	OwnershipQuery   string
}
