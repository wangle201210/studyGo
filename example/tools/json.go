package main

import (
	"flag"
	"github.com/wangle201210/studyGo/tools/devlop"
)

var (
	js string
)

func init() {
	flag.StringVar(&js, "js", "", "json string")
}

func JsonDecode() {
	flag.Parse()
	js = "{\"id\":1663301542005,\"content\":\"\\346\\265\\267\\345\\206\\273\\264\\346\\273\\264\",\"title\":\"\\346\\265\\267\\345\\206\\2332(\\346\\265\\267\\345\\206\\2332)\",\"extra\":\"\\/hospital\",\"targetId\":[60724675],\"pushExt\":null,\"channelId\":\"normal\"}\n"
	s := devlop.NewJsonData(js).GetDecode()
	println(s)
}
