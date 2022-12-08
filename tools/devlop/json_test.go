package devlop

import (
	"encoding/base64"
	"encoding/json"
	"net/url"
	"testing"
)

type ItemEnv struct {
	Type  string `json:"type"`  // 环境变量类型
	Key   string `json:"key"`   // 环境变量名称
	Value string `json:"value"` // 环境变量键值
	Brief string `json:"brief"` // 环境变量描述
}

func TestJson(t *testing.T) {
	//s := "{\"title\":\"\346\270\270\345\256\242\"}"
	//println(NewJsonData(s).GetDecode())

	// [{"type":"","key":"REGION","value":"cd","brief":""},{"type":"","key":"ZONE","value":"cd001","brief":""}]
	u := "http:bai.com?a=b&c="
	parse, err := url.Parse(u)
	if err != nil {
		t.Error(err)
		return
	}
	get := parse.Query().Get("c")
	e := []ItemEnv{}
	bytes, err := base64.StdEncoding.DecodeString(get)
	err = json.Unmarshal(bytes, &e)
	if err != nil {
		t.Error(err)
		return
	}
	t.Logf("%+v", e)
}
