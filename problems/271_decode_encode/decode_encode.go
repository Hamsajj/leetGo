// Package decode_encode
// difficulty: medium
// link: https://leetcode.com/problems/encode-and-decode-strings/
// free-link: https://www.lintcode.com/problem/659/
package decode_encode

import "fmt"

func encode(strs []string) string {
	encoded := ""
	for _, str := range strs {
		encoded += fmt.Sprintf("%d%s", len(str), str)
	}
	return encoded
}

func decode(str string) []string {
	var decoded []string
	for i := 0; i < len(str); {
		nextStrLength := int(str[i] - '0')
		decoded = append(decoded, str[i+1:i+nextStrLength+1])
		i += nextStrLength + 1
	}
	return decoded
}
