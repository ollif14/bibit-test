package main

import (
	"fmt"
	"sort"
)

func main() {
	arrStrs := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	results := make(map[string][]string)
	for _, arrStr := range arrStrs {
		strBytes := []byte(arrStr)
		sort.SliceStable(strBytes, func(i, j int) bool {
			return strBytes[i] < strBytes[j]
		})
		str := string(strBytes)
		results[str] = append(results[str], arrStr)
	}
	var newArrStr [][]string
	for result := range results {
		newArrStr = append(newArrStr, results[result])
	}

	fmt.Println(newArrStr)
}
