// pgstmt
package json2pgstmt

import (
	"bytes"
)

type ProcParam struct {
	Name  string
	Value string
	Type  string
	Order string
}

type JsonRequest struct {
	Schema string
	Name   string
	Fields string
	Order  string
	Offset int
	Limit  int

	Params []ProcParam
}

func (r JsonRequest) Statement() (string, error) {

	var b bytes.Buffer

	b.WriteString("select ")

	if r.Fields == "" {
		b.WriteString("* ")
	} else {
		b.WriteString(checkSafeSql(r.Fields))
	}

	b.WriteString("from")
	b.WriteString(r.getQualifiedObjectName())

	if r.Order != "" {
		b.WriteString("order by")
		b.WriteString(checkSafeSql(r.Order))
	}

	if r.Offset != 0 {
		b.WriteString(" offset ")
		b.WriteString(string(r.Offset))
	}

	if r.Limit != 0 {
		b.WriteString(" limit ")
		b.WriteString(string(r.Limit))
	}

	return string(b.Bytes()), nil
}

// TODO Create logic
func (r JsonRequest) getQualifiedObjectName() string {
	return checkSafeSql(r.Name)
}

// TODO Check safe sql in the fields
func checkSafeSql(sql string) string {
	return " " + sql + " "
}
