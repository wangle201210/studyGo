package devlop

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Mod struct {
	FilePath string
}

func (m *Mod) ModCheck() {
	file, err := os.OpenFile(m.FilePath, os.O_RDWR, 0666)
	if err != nil {
		if os.IsNotExist(err) {
			return
		} else {
			fmt.Println("Open file error!", err)
		}
		return
	}
	defer file.Close()

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		if !m.checkLine(line) {
			break
		}
	}
}

func (m *Mod) checkLine(s string) bool {
	if !strings.Contains(s, "git.medlinker.com/foundations/med-common-sdk") {
		return true
	}
	split := strings.Split(s, " ")
	if len(split) != 2 {
		return true
	}
	version := split[1]
	// 不能包含任何commit，即 tag-日期-commitID v2.0.0-20211108115335-fcd935171749
	// 只能是 v2.0.0 这种纯tag的模式
	// 即每次上线必须用正式的tag
	if strings.Contains(version, "-") {
		fmt.Println("git.medlinker.com/foundations/med-common-sdk 请勿使用commit，必须使用正式tag")
		return false
	}
	return true
}
