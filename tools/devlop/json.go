package devlop

import (
	"encoding/json"
	"fmt"
)

type jsonData struct {
	s string
}

func NewJsonData(s string) *jsonData {
	return &jsonData{
		s: s,
	}
}

func (j *jsonData) GetDecode() (s string) {
	m := make(map[string]interface{})
	if err := json.Unmarshal([]byte(j.s), &m); err != nil {
		panic(err.Error() + "\n" + j.s)
	}
	s = fmt.Sprintf("%+v", m)
	return
}
