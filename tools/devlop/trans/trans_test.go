package trans

import (
	"encoding/json"
	"testing"
)

var data = &Base{Str: "\x227"}

func TestTrans(t *testing.T) {
	a := "{\x22jsonrpc\x22:\x222.0\x22,\x22method\x22:\x22search\x22,\x22params\x22:[null,\x227\x22,\x22\x5Cu4e0a\x5Cu6d77\x5Cu4e2d\x5Cu533b\x5Cu533b\x5Cu9662\x5Cu77f3\x5Cu95e8\x5Cu4e00\x5Cu8def\x5Cu95e8\x5Cu8bca\x5Cu90e8\x22,0,15,0,0,\x228.8.9\x22,false,{\x22uniqueHospital\x22:true}],\x22id\x22:1}"
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(a), &m)
	if err != nil {
		t.Errorf("err (%+v)", err)
		return
	}
	t.Logf("data (%+v)", m)
	//data.Trans()
}
