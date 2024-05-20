package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func csvMerge() {
	// CSV文件列表
	fileNames := []string{"file1.csv", "file2.csv", "file3.csv"}
	// 输出CSV文件
	outputFile := "merged.csv"

	// 创建输出文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// 创建CSV写入器
	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	// 设置为true，如果要在合并时保留所有文件的头部（列标题）
	includeHeaders := false
	fileNames = getAllFile("pg")
	// 遍历并处理每个文件
	for idx, fileName := range fileNames {
		inFile, err := os.Open(path.Join("data_export", fileName))
		if err != nil {
			panic(err)
		}
		defer inFile.Close()

		// 创建CSV读取器
		reader := csv.NewReader(inFile)

		// 是否是第一个文件（用于决定是否写入头部）
		isFirstFile := idx == 0

		// 读取和写入CSV内容
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				panic(err)
			}

			// 如果是第一行，检查是否需要写入头部
			if isFirstFile || (idx > 0 && !includeHeaders) {
				writer.Write(record)
				isFirstFile = false
			} else if idx > 0 && includeHeaders {
				// 如果设置为保留所有文件的头部，则每个文件的第一行都写入
				writer.Write(record)
			}

			// 如果不是第一个文件，且设置为不包含头部，则跳过每个文件的第一行
			includeHeaders = false
		}
	}

	fmt.Println("CSV files have been merged into", outputFile)
}

func getAllFile(prefix string) (list []string) {
	// 获取当前工作目录
	currentDir := "data_export"

	// 读取当前目录下的所有文件和文件夹
	files, err := ioutil.ReadDir(currentDir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	// 遍历文件列表
	for _, file := range files {
		// 检查文件是否是常规文件（不是目录）
		if file.Mode().IsRegular() {
			// 如果文件扩展名是.csv，则打印出文件名
			name := file.Name()
			if filepath.Ext(name) == ".csv" {
				if strings.HasPrefix(name, prefix) {
					list = append(list, name)
				}
			}
		}
	}
	return
}
