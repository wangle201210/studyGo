package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type RequestAccess struct {
	Platform        string `json:"wxPlatform" form:"wxPlatform"`
	RedirectUrl     string `json:"redirectUrl" form:"redirectUrl"  valid:"redirectUrl @required#redirectUrl不能为空"`
	Scope           int    `json:"scope" form:"scope"`
	CurrentPlatform string `json:"currentWxPlatform" form:"currentWxPlatform"`
	SwimLane        string `json:"swimlane" form:"swimlane"`
}

func main() {
	housingFund()
	return
	//JsonDecode()
	//return
	//fmt.Println(decimalToAny(math.MaxInt64, 62))
	//return
	//devlop.Dns()
	//var n1 []int
	//n2 := []int{1, 2, 3}
	//n3 := append(n1, n2...)
	//fmt.Printf("%+v %d", n3, cap(n3))
	//return
	////birthDate:"1973\345\271\26410\346\234\2105\346\227\245"
	//s := "{\"id\":1663301542005,\"content\":\"\346\265\267\345\206\2332:\346\273\264\346\273\264\",\"title\":\"\346\265\267\345\206\2332(\346\265\267\345\206\2332)\",\"extra\":\"\\/hospital\",\"targetId\":[60724675],\"pushExt\":null,\"channelId\":\"normal\"}"
	s := "{\"name\":\"\351\203\255\345\273\272\346\230\216\"}"
	r := make(map[string]interface{})
	err := json.Unmarshal([]byte(s), &r)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", r)
	return
	//var a = -9223372036854775808
	//var b uint = 18446744073709551615
	//
	//fmt.Printf("a的类型为:[%T],a的值为:[%d]\n", a, a)
	//fmt.Printf("b的类型为:[%T],b的值为:[%d]\n", b, b)

	//s := "r=1&n=301&t=0&v=1"
	//stringXml(&s)
	//println(s)
	//img2txt.Img2txt("/Users/med/Downloads/821646621584_.pic.png", 200, []string{"@", "#", "*", "%", "+", ",", ".", " "}, "\n", "./保存的文本.txt")
	housingFund()
	//sql2struct()
	//modCheck()
	//wechatDo()
	//img2txt.Img()
}

var tenToAny map[int]string = map[int]string{0: "0", 1: "1", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9",
	10: "a", 11: "b", 12: "c", 13: "d", 14: "e", 15: "f", 16: "g", 17: "h", 18: "i", 19: "j", 20: "k", 21: "l", 22: "m", 23: "n", 24: "o", 25: "p", 26: "q", 27: "r", 28: "s", 29: "t", 30: "u", 31: "v", 32: "w", 33: "x", 34: "y", 35: "z",
	36: ":", 37: ";", 38: "<", 39: "=", 40: ">", 41: "?", 42: "@", 43: "[", 44: "]", 45: "^", 46: "_", 47: "{", 48: "|", 49: "}",
	50: "A", 51: "B", 52: "C", 53: "D", 54: "E", 55: "F", 56: "G", 57: "H", 58: "I", 59: "J", 60: "K", 61: "L", 62: "M", 63: "N", 64: "O", 65: "P", 66: "Q", 67: "R", 68: "S", 69: "T", 70: "U", 71: "V", 72: "W", 73: "X", 74: "Y", 75: "Z"}

// 10进制转任意进制
func decimalToAny(num, n int) string {
	new_num_str := ""
	var remainder int
	var remainder_string string
	for num != 0 {
		remainder = num % n
		if 76 > remainder && remainder > 9 {
			//1x2QluMLiT
			//1x2ClZu|{iF
			//aZl8N0y58M7
			if remainder >= 36 {
				remainder += 14
			}
			remainder_string = tenToAny[remainder]
		} else {
			remainder_string = strconv.Itoa(remainder)
		}
		new_num_str = remainder_string + new_num_str
		num = num / n
	}
	return new_num_str
}

// map根据value找key
func findkey(in string) int {
	result := -1
	for k, v := range tenToAny {
		if in == v {
			result = k
		}
	}
	return result
}

// 任意进制转10进制
func anyToDecimal(num string, n int) int {
	var new_num float64
	new_num = 0.0
	nNum := len(strings.Split(num, "")) - 1
	for _, value := range strings.Split(num, "") {
		tmp := float64(findkey(value))
		if tmp != -1 {
			new_num = new_num + tmp*math.Pow(float64(n), float64(nNum))
			nNum = nNum - 1
		} else {
			break
		}
	}
	return int(new_num)
}
