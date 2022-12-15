package services

import (
	"brewess2/models"
	"brewess2/utils"
	"context"
	"fmt"
	"log"
	"strings"
	"time"
)

type Table models.Table

const sequence string = `create sequence IF NOT EXISTS %s_id_seq start with 1 increment by 1;`
const createTableStart string = `create table IF NOT EXISTS public.%s ( id integer NOT null default nextval('%s_id_seq'), title text, `
const createTableColumn string = ` %s %s,`
const createTableEnd string = `constraint "%s_pkey" primary key (id)) with (oids = false) tablespace pg_default;`

const ownership string = `alter table public.%s owner to` + utils.User + `;`

func ColumnCreator(name string, datatype string) *models.Column {
	c := new(models.Column)
	c.Name = name
	c.DataType = datatype
	return c
}

func TableCreator(name string, columns []*models.Column) *Table {
	t := new(Table)
	t.Name = name
	t.Columns = columns

	var sequenceWriter strings.Builder
	fmt.Fprintf(&sequenceWriter, sequence, t.Name)
	t.SequenceQuery = sequenceWriter.String()

	var createQueryWriter strings.Builder
	fmt.Fprintf(&createQueryWriter, createTableStart, t.Name, t.Name)
	for _, c := range columns {
		fmt.Fprintf(&createQueryWriter, createTableColumn, c.Name, c.DataType)
	}
	fmt.Fprintf(&createQueryWriter, createTableEnd, t.Name)
	t.CreateTableQuery = createQueryWriter.String()

	var ownershipWriter strings.Builder
	fmt.Fprintf(&ownershipWriter, sequence, t.Name)
	t.OwnershipQuery = ownershipWriter.String()

	return t
}

func (t Table) ExecuteQuery(q string) {
	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	res, err := dbconn.ExecContext(ctx, q)
	if err != nil {
		log.Printf("Error %s when creating %s table", err, t.Name)
		return
	}
	rows, err := res.RowsAffected()
	if err != nil {
		log.Printf("Error %s when getting rows affected in table %s", err, t.Name)
		return
	}
	log.Printf("Rows affected when creating %s table: %d", t.Name, rows)
	return
}

func (t Table) CreateResourceTable() {
	log.Printf(t.SequenceQuery)
	log.Printf(t.CreateTableQuery)
	log.Printf(t.OwnershipQuery)
	t.ExecuteQuery(t.SequenceQuery)
	t.ExecuteQuery(t.CreateTableQuery)
	t.ExecuteQuery(t.OwnershipQuery)
}
