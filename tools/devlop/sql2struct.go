package devlop

import (
	"bytes"
	"github.com/olekukonko/tablewriter"
	"regexp"
	"strings"
)

var (
	typeMap map[string]string
)

type s2s struct {
	sql   string
	table string
	rows  []row
	// 备注的位置，默认在有边(可选择 top,right)
	commentPos string
	// 需不需要加gogo的json后缀
	gogo bool
}

type row struct {
	// 字段名
	name string
	// 类型
	typ string
	// 备注
	comment string
	// 整句sql
	content string
}

type S2sOpt func(*s2s)

func Sql2struct(sql string, opts ...S2sOpt) (s *s2s) {
	s = &s2s{
		sql:        sql,
		commentPos: "right",
	}
	for _, opt := range opts {
		opt(s)
	}
	s.getTableName()
	s.getRows()
	return
}

func (s *s2s) getRows() {
	s.rows = []row{}
	reg, err := regexp.Compile("`" + `([\w_]+)` + "`" + ` ([\w_]+).*`)
	if err != nil {
		panic(err)
	}
	as := reg.FindAllStringSubmatch(s.sql, -1)
	for _, a := range as {
		//a[0] 匹配结果， a[1] 字段名， a[2] 字段类型
		r := row{
			content: a[0],
			name:    a[1],
			typ:     a[2],
			comment: filterComment(a[0]),
		}
		s.rows = append(s.rows, r)
		//fmt.Printf("(%+v)\n", a)
	}
}

// 找到备注
func filterComment(s string) string {
	index := strings.Index(s, "COMMENT")
	if index == -1 {
		return ""
	}
	//去掉COMMENT空格和单引号的位置
	lc := len("COMMENT") + 2
	// 长度不够就不要了（理论上不会存在的）
	if len(s) < index+lc || index+lc > len(s)-2 {
		return ""
	}
	com := s[index+lc : len(s)-2]
	com = strings.TrimSpace(com)
	return com
}

func (s *s2s) getTableName() {
	reg, err := regexp.Compile(`(TABLE|table)\s+` + "`" + `([\w_]+)` + "`" + `\s+\(\s+`)
	if err != nil {
		panic(err)
	}
	as := reg.FindStringSubmatch(s.sql)
	if len(as) < 3 {
		panic("没表名")
	}
	s.table = as[2]
}

func getType(str string) string {
	if t, ok := typeMap[str]; ok {
		return t
	}
	return "string"
}

// camel string, xx_yy to XxYy
func camelString(s string, lower ...bool) string {
	data := make([]byte, 0, len(s))
	l := false
	if len(lower) > 0 {
		l = lower[0]
	}
	j := false
	k := false
	num := len(s) - 1
	for i := 0; i <= num; i++ {
		d := s[i]
		if k == false && d >= 'A' && d <= 'Z' {
			k = true
		}
		if d >= 'a' && d <= 'z' && (j || k == false) {
			if i != 0 || !l {
				d = d - 32
			}
			j = false
			k = true
		}
		if k && d == '_' && num > i && s[i+1] >= 'a' && s[i+1] <= 'z' {
			j = true
			continue
		}
		data = append(data, d)
	}
	return string(data[:])
}

func init() {
	typeMap = map[string]string{
		"tinyint":    "int64",
		"smallint":   "int64",
		"int":        "int64",
		"mediumint":  "int64",
		"bigint":     "int64",
		"float":      "float64",
		"double":     "float64",
		"decimal":    "float64",
		"char":       "string",
		"varchar":    "string",
		"text":       "string",
		"mediumtext": "string",
		"longtext":   "string",
		"time":       "time.Time",
		"date":       "time.Time",
		"datetime":   "time.Time",
		"timestramp": "int64",
		"enum":       "string",
		"set":        "string",
		"blob":       "string",
	}
}

func (s *s2s) String() (str string) {
	structResult := "type " + camelString(s.table) + " struct {"
	var rows [][]string
	for _, r := range s.rows {
		// 字段名 + 字段类型 + 备注
		row := []string{
			camelString(r.name), getType(r.typ),
		}
		if len(r.comment) > 0 {
			row = append(row, "//"+r.comment)
		}
		rows = append(rows, row)
	}
	buffer := bytes.NewBuffer(nil)
	tw := tablewriter.NewWriter(buffer)
	tw.SetBorder(false)
	tw.SetRowLine(false)
	tw.SetAutoWrapText(false)
	tw.SetColumnSeparator("")
	tw.AppendBulk(rows)
	tw.Render()
	defineContent := buffer.String()
	buffer.Reset()
	buffer.WriteString(defineContent)
	structResult += "\n" + buffer.String()
	structResult += "}"
	return structResult
}

func WithCommentPos(str string) S2sOpt {
	return func(s *s2s) {
		s.commentPos = str
	}
}
