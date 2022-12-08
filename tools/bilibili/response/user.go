package response

type User struct {
	Code    int64  `json:"code"`
	Message string `json:"message"`
	Ttl     int64  `json:"ttl"`
	Data    struct {
		Mid  int64  `json:"mid"`
		Name string `json:"name"`
		Sex  string `json:"sex"`
		Face string `json:"face"`
	} `json:"data"`
}

type NameSearch struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Ttl     int    `json:"ttl"`
	Data    struct {
		Result []struct {
			Type  string `json:"type"`
			Mid   int    `json:"mid"`
			Uname string `json:"uname"`
		} `json:"result"`
	} `json:"data"`
}

type UserInfo struct {
	Mid     int32  `json:"mid"`
	Name    string `json:"name"`
	Sex     string `json:"sex"`
	Face    string `json:"face"`
	Archive int64  `json:"archive"`
	Likes   int64  `json:"likes"`
}
