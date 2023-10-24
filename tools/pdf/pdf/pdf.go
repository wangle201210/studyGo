package pdf

import (
	"context"
	"encoding/json"
	"github.com/google/go-tika/tika"
	"golang.org/x/net/html"
	"log"
	"os"
	"strings"
	"unicode"
)

type Tika struct {
	XTIKAContent string `json:"X-TIKA:content"`
}

func Diff(path1, path2 string, minLength int) {
	s1 := read(path1)
	s2 := read(path2)
	DoFind(s1, s2, minLength)
}

func Self(path1 string, minLength int) {
	s1 := read(path1)
	println(s1)

	DoFindSelf(s1, minLength)
}

func read(path string) string {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	client := tika.NewClient(nil, "http://localhost:9998/tika")
	body, err := client.Parse(context.Background(), f)
	c := &Tika{}
	err = json.Unmarshal([]byte(body), c)
	if err != nil {
		panic(err)
	}
	content := c.XTIKAContent
	// content = removeHTMLTags(content)
	// content = removeOther(content)
	// content = removeLetters(content)
	// Print the extracted text
	return content
}

func removeOther(s string) string {
	s = strings.ReplaceAll(s, "     ", "") // 大量的空格需要被删除
	s = strings.ReplaceAll(s, "\n", "")    // 大量的换行符也需要被删除
	s = strings.ReplaceAll(s, "�", "")     // 乱码符号被替换

	return s
}

func removeHTMLTags(htmlContent string) string {
	doc, _ := html.Parse(strings.NewReader(htmlContent))
	var result strings.Builder
	var removeHTML func(*html.Node)
	// 匿名递归函数
	removeHTML = func(n *html.Node) {
		if n.Type == html.TextNode {
			result.WriteString(n.Data)
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			removeHTML(c)
		}
	}

	removeHTML(doc)
	return result.String()
}

func removeLetters(input string) string {
	var result []rune
	for _, c := range input {
		if !unicode.IsLetter(c) {
			result = append(result, c)
		}
	}
	return string(result)
}
