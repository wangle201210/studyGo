package lib

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
)

const (
	DefaultHttpCallTimeOut = 5 * time.Second
)

type HttpParam struct {
	// Header头
	Header map[string]string
	// 参数
	Param map[string]interface{}
	// cookie
	Cookie *http.Cookie
	// 超时时间设置, 秒, 不传取默认值
	Timeout uint32
}

// HttpPost
func HttpPost(url string, args *HttpParam, result interface{}) error {
	data, err := json.Marshal(args.Param)
	if nil != err {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(data))
	if nil != err {
		return err
	}

	for key, value := range args.Header {
		req.Header.Set(key, value)
	}
	if ok := req.Header.Get("Content-Type"); "" == ok {
		req.Header.Set("Content-Type", "application/json")
	}
	if args.Cookie != nil {
		req.AddCookie(args.Cookie)
	}
	client := &http.Client{Timeout: DefaultHttpCallTimeOut}
	if args.Timeout > 0 {
		client.Timeout = time.Duration(args.Timeout) * time.Second
	}
	resp, err := client.Do(req)
	if nil != err {
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Error("HttpPost err, url: [%v], params: [%s], body: [%s]", url, string(data), string(body))
		return errors.New("request response status code err")
	}

	defer func() {
		if err = resp.Body.Close(); nil != err {
			fmt.Println(err)
		}
	}()

	err = json.Unmarshal(body, &result)
	if nil != err {
		return err
	}
	return nil
}

// HttpGet
func HttpGet(url string, args *HttpParam, result interface{}) error {
	req, err := http.NewRequest("GET", url, nil)
	if nil != err {
		return err
	}
	for key, value := range args.Header {
		req.Header.Set(key, value)
	}
	if ok := req.Header.Get("Content-Type"); "" == ok {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
	if args.Cookie != nil {
		req.AddCookie(args.Cookie)
	}
	params := req.URL.Query()
	for key, value := range args.Param {
		params.Add(key, fmt.Sprintf("%v", value))
	}

	req.URL.RawQuery = params.Encode()

	client := &http.Client{Timeout: DefaultHttpCallTimeOut}
	if args.Timeout > 0 {
		client.Timeout = time.Duration(args.Timeout) * time.Second
	}
	resp, err := client.Do(req)
	if nil != err {
		return err
	}

	body, _ := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log.Error("HttpGet err, url: [%v], params: [%s], body: [%s]", url, req.URL.RawQuery, string(body))
		return errors.New("request response status code err")
	}

	defer func() {
		if err = resp.Body.Close(); nil != err {
			fmt.Println(err)
		}
	}()

	err = json.Unmarshal(body, result)
	if nil != err {
		return err
	}
	return nil
}
