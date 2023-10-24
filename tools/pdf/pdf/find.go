package pdf

import "fmt"

func DoFind(str1, str2 string, minLength int) {
	results := findCommonSequence(str1, str2, minLength)
	if len(results) == 0 {
		fmt.Println("没找到重复内容")
	} else {
		// fmt.Println("找到了重复内容如下:")
		// for _, str := range results {
		// 	fmt.Println(str)
		// }
	}
	// println(str1)
	// println("==================")
	// println(str2)
}

func findCommonSequence(str1, str2 string, minLength int) []string {
	m, n := len([]rune(str1)), len([]rune(str2))
	lookupTable := make([][]int, m+1)
	for i := range lookupTable {
		lookupTable[i] = make([]int, n+1)
	}
	var matchingSubstrings []string
	runeStr1 := []rune(str1)
	runeStr2 := []rune(str2)
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if runeStr1[i-1] == runeStr2[j-1] {
				lookupTable[i][j] = lookupTable[i-1][j-1] + 1
				if lookupTable[i][j] >= minLength {
					matchingSubstrings = append(matchingSubstrings, string(runeStr1[i-lookupTable[i][j]:i]))
					fmt.Printf("第%d处重复,内容: %s \n", len(matchingSubstrings), string(runeStr1[i-lookupTable[i][j]:i]))
				}
			}
		}
	}
	return matchingSubstrings
}

const Base = 107
const mod = 10007

func DoFindSelf(str1 string, minLength int) {
	getRepeatedSubStr(str1, minLength)
	// fmt.Println(getRepeatedSubStr(str1, minLength))
	// for i := len([]rune(str1)); i > minLength; i-- {
	// 	substring := searchLCS(str1, i)
	// 	if len(substring) > minLength {
	// 		fmt.Printf("公共子串: %s, 长度: %d\n", substring, len(substring))
	// 		break
	// 	}
	// }
}
func searchLCS(s string, length int) string {
	rs := []rune(s)
	if len(rs) < length {
		return ""
	}
	powers := make([]uint64, length)
	powers[0] = 1
	for i := 1; i < length; i++ {
		powers[i] = (powers[i-1] * Base) % mod
	}
	hashes := make([]uint64, len(rs)-length+1)
	hashes[0] = 0
	for i := 0; i < len(rs); i++ {
		hashes[0] = (hashes[0]*Base + uint64(rs[i])) % mod
		if i < length-1 {
			continue
		}
		if i >= length {
			prevHash := hashes[i-length+1]
			firstChar := uint64(rs[i-length])
			hashes[i-length+2] = ((prevHash-firstChar*powers[length-1])*Base + uint64(rs[i])) % mod
		}
	}

	hashesMap := make(map[uint64][]int)
	for i, hash := range hashes {
		if hashesMap[hash] == nil {
			hashesMap[hash] = []int{}
		}
		hashesMap[hash] = append(hashesMap[hash], i)
	}

	for _, positions := range hashesMap {
		if len(positions) <= 1 {
			continue
		}
		for i := 1; i < len(positions); i++ {
			if positions[i]-positions[i-1] >= length {
				return string(rs[positions[i-1] : positions[i-1]+length])
			}
		}
	}
	return ""
}

func getRepeatedSubStr(str string, length int) []string {
	record := make(map[string]int)
	for i := 0; i+length <= len([]rune(str)); i++ {
		subStr := string([]rune(str)[i : i+length])
		record[subStr]++
	}
	result := make([]string, 0)
	for key, value := range record {
		if value > 1 {
			result = append(result, key)
			fmt.Printf("第%d处重复,内容: %s \n", len(result), key)
		}
	}
	return result
}
