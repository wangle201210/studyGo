package devlop

import (
	"bytes"
	"fmt"
	"strings"
)

var (
	gogoFmt = "[(gogoproto.jsontag) = \"%s\"]"
)

type s2p struct {
	s2s
}

func Sql2proto(sql string, opts ...S2sOpt) (s *s2p) {
	s = new(s2p)
	s.sql = sql
	s.commentPos = "right"
	for _, opt := range opts {
		opt(&s.s2s)
	}
	s.getTableName()
	s.getRows()
	return
}

func getProtoType(t string) (typeName string) {
	t = strings.ToLower(t)
	switch t {
	case "binary", "varbinary", "blob", "tinyblob", "mediumblob", "longblob":
		typeName = "bytes"
	case "bit", "int", "tinyint", "small_int", "smallint", "medium_int", "mediumint", "serial", "int8", "big_int", "bigint", "bigserial":
		typeName = "int64"
	case "real":
		typeName = "float"
	case "float", "double", "decimal", "smallmoney":
		typeName = "double"
	case "bool":
		typeName = "bool"
	case "datetime", "timestamp", "date", "time":
		typeName = "int64"
	case "text", "char":
		typeName = "string"
	default:
		typeName = "string"
	}
	return
}
func (s *s2p) String() (str string) {
	structResult := "message " + camelString(s.table) + " {"
	buffer := bytes.NewBuffer(nil)
	for _, r := range s.rows {
		row := getProtoType(r.typ) + "\t" + camelString(r.name, true)
		if s.gogo {
			row += " " + fmt.Sprintf(gogoFmt, r.name)
		}
		bpc := combProtoComment(row, r.comment, s.commentPos)
		buffer.WriteString(bpc)
	}
	structResult += "\n" + buffer.String()
	structResult += "}"
	return structResult
}

func combProtoComment(row, comment, pos string) (s string) {
	s = "\t" + row + ";"
	if len(comment) > 0 {
		if pos == "right" || pos == "r" {
			s += " //" + comment
		} else if pos == "top" || pos == "t" {
			s = "\t//" + comment + "\n" + s
		}
	}
	s += "\n"
	return
}

func WithGogo() S2sOpt {
	return func(s *s2s) {
		s.gogo = true
	}
}
