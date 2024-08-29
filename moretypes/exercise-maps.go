package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	// strings.Fields でスペース区切りで文字列を分割して配列に格納
	for _, w := range strings.Fields(s) {
		m[w]++
	}

	return m
}

func main(){
	wc.Test(WordCount)
}