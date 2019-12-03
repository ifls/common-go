package string

import (
	"log"
	"strings"
)

//76.07% 24.33%
//once 20191203 23:46
func isPalindrome(s string) bool {
	filtered := filter(s)
	lowered := strings.ToLower(filtered)
	log.Println(lowered)

	reserved := reverse(lowered)
	log.Println(reserved)
	return reserved == lowered
}

func filter(s string) string {
	s2 := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] <= 57 && s[i] >= 48) || (s[i] >= 65 && s[i] <= 90) || (s[i] >= 97 && s[i] <= 122) {
			s2 = append(s2, s[i])
		}
	}
	return string(s2)
}

func reverse(s string) string {
	s2 := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		s2 = append(s2, s[i])
	}
	return string(s2)
}
