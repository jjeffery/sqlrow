package sqlrow

import (
	"testing"
)

func TestPrepare(t *testing.T) {
	defer func() { Default.Dialect = nil }()
	tests := []struct {
		row     interface{}
		sql     string
		queries map[string]string
	}{
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Name string
			}{},
			sql: "insert into tbl({}) values({})",
			queries: map[string]string{
				"mysql":    "insert into tbl(`name`) values(?)",
				"postgres": `insert into tbl("name") values($1)`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Name string
			}{},
			sql: "insert into tbl({all}) values({})",
			queries: map[string]string{
				"mysql":    "insert into tbl(`id`,`name`) values(?,?)",
				"postgres": `insert into tbl("id","name") values($1,$2)`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key"`
				Name string
			}{},
			sql: "insert into tbl({}) values({})",
			queries: map[string]string{
				"mysql":    "insert into tbl(`id`,`name`) values(?,?)",
				"postgres": `insert into tbl("id","name") values($1,$2)`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Name string
			}{},
			sql: "update tbl set {} where {}",
			queries: map[string]string{
				"mysql":    "update tbl set `name`=? where `id`=?",
				"postgres": `update tbl set "name"=$1 where "id"=$2`,
			},
		},
		{
			row: struct {
				ID    string `sql:"primary key auto increment"`
				Hash  string `sql:"pk"`
				Name  string
				Count int
			}{},
			sql: "update [xxx]\nset\n{}\nwhere {}",
			queries: map[string]string{
				"mysql":    "update `xxx` set `name`=?,`count`=? where `id`=? and `hash`=?",
				"postgres": `update "xxx" set "name"=$1,"count"=$2 where "id"=$3 and "hash"=$4`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Name string
			}{},
			sql: "delete from tbl where {}",
			queries: map[string]string{
				"mysql":    "delete from tbl where `id`=?",
				"postgres": `delete from tbl where "id"=$1`,
			},
		},
		{
			row: struct {
				ID    string `sql:"primary key auto increment"`
				Hash  string `sql:"pk"`
				Name  string
				Count int
			}{},
			sql: "delete from `xxx`\n-- this is a comment\nwhere {}",
			queries: map[string]string{
				"mysql":    "delete from `xxx` where `id`=? and `hash`=?",
				"postgres": `delete from "xxx" where "id"=$1 and "hash"=$2`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Name string
			}{},
			sql: "select {} from tbl where {}",
			queries: map[string]string{
				"mysql":    "select `id`,`name` from tbl where `id`=?",
				"postgres": `select "id","name" from tbl where "id"=$1`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Name string
			}{},
			sql: "select {alias t} from tbl t where {pk,alias t}",
			queries: map[string]string{
				"mysql":    "select t.`id`,t.`name` from tbl t where t.`id`=?",
				"postgres": `select t."id",t."name" from tbl t where t."id"=$1`,
			},
		},
		{
			row: struct {
				ID   string `sql:"primary key auto increment"`
				Home struct {
					Postcode string
				}
			}{},
			sql: "select {alias t} from tbl t where {pk,alias t}",
			queries: map[string]string{
				"mysql":    "select t.`id`,t.`home_postcode` from tbl t where t.`id`=?",
				"postgres": `select t."id",t."home_postcode" from tbl t where t."id"=$1`,
			},
		},
		{
			row: struct {
				ID    string `sql:"primary key auto increment"`
				Hash  string `sql:"pk"`
				Name  string
				Count int
			}{},
			sql: "select {} from `xxx`\nwhere {}",
			queries: map[string]string{
				"mysql":    "select `id`,`hash`,`name`,`count` from `xxx` where `id`=? and `hash`=?",
				"postgres": `select "id","hash","name","count" from "xxx" where "id"=$1 and "hash"=$2`,
			},
		},
	}

	for i, tt := range tests {
		for dialect, query := range tt.queries {
			Default.Dialect = DialectFor(dialect)
			fns := []func(interface{}, string) (*Stmt, error){
				Prepare,
				Default.Prepare,
			}
			for _, fn := range fns {
				stmt, err := fn(tt.row, tt.sql)
				if err != nil {
					t.Errorf("%d: expected no error: got %v", i, err)
					continue
				}
				if stmt.String() != query {
					t.Errorf("%d: %s: expected=%q, actual=%q", i, dialect, query, stmt.String())
				}
			}
		}
	}
}
