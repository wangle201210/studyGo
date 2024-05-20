package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"github.com/google/go-tika/tika"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func tikaDo() {
	filePath := "./bt.pdf"
	// 调用函数ReadPdf解析pdf文件
	content, err := ReadPdf(filePath) // Read local pdf file
	if err != nil {
		panic(err)
	}
	// 将pdf的所有内容写入html文件)
	err = ioutil.WriteFile("./out.html", []byte(content), 0666)
	if err != nil {
		log.Fatal(err)
	}
	// 先将html中的<title>标签去掉,因为此标签中含有特殊字符,会导致xml语法出错
	delerr := deleteTitle("out.html")
	if delerr != nil {
		log.Fatal(delerr)
	}
	err = ReadHtml("out.html")
	if err != nil {
		log.Fatal(err)
	}
}

// 删除html中的title标签
func deleteTitle(filename string) error {
	cmd := exec.Command("bash", "-c", fmt.Sprintf("sed -i '65d' %s", filename))
	_, err := cmd.Output()
	if err != nil {
		return err
	}
	return nil
}

// 解析PDF文件
func ReadPdf(path string) (string, error) {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return "", err
	}
	client := tika.NewClient(nil, "http://127.0.0.1:9998")
	return client.Parse(context.TODO(), f)
}

// 定义多个结构体，用来接收反序列化数据
type (
	html struct {
		XMLNAME xml.Name `xml:"html"`
		Body    htmlBody `xml:"body"`
	}
	htmlBody struct {
		XMLNAME xml.Name `xml:"body"`
		Div     htmlDiv  `xml:"div"`
	}
	htmlDiv struct {
		P []string `xml:"p"`
	}
)

// 读取html文件，并反序列化到结构体中
func ReadHtml(filename string) error {
	rf, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	html := html{}
	err = xml.Unmarshal(rf, &html)
	if err != nil {
		return err
	}
	var b []byte
	for _, v := range html.Body.Div.P {
		b = append(b, []byte(v)...)
	}
	err = ioutil.WriteFile("res.doc", b, 0644)
	if err != nil {
		return err
	}
	return nil
}
