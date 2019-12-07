package string

import (
	"strings"
)

//76.07% 24.33%
//once 20191203 23:46
func IsPalindrome1(s string) bool {
	s = strings.ToLower(s)
	filtered := filter(s)

	//log.Println(filtered)

	reserved := reverse(filtered)
	//log.Println(reserved)
	return reserved == filtered
}

func filter(s string) string {
	s2 := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if (s[i] <= 57 && s[i] >= 48) || (s[i] >= 97 && s[i] <= 122) {
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

func IsPalindrome2(s string) bool {
	//s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for !((s[left] <= 57 && s[left] >= 48) || (s[left] >= 97 && s[left] <= 122) || (s[left] >= 65 && s[left] <= 90)) {
			left++
		}

		for !((s[right] <= 57 && s[right] >= 48) || (s[right] >= 97 && s[right] <= 122) || (s[right] >= 65 && s[right] <= 90)) {
			right--
		}
		bleft := s[left]
		bright := s[right]
		if s[left] >= 97 && s[left] <= 122 {
			bleft = s[left] - 32
		}
		if s[right] >= 97 && s[right] <= 122 {
			bright = s[right] - 32
		}

		if bleft != bright {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func IsPalindrome3(s string) bool {
	s = strings.ToLower(s)
	left := 0
	right := len(s) - 1
	for left < right {
		for !((s[left] <= 57 && s[left] >= 48) || (s[left] >= 97 && s[left] <= 122)) {
			left++
		}

		for !((s[right] <= 57 && s[right] >= 48) || (s[right] >= 97 && s[right] <= 122)) {
			right--
		}

		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}

func isApl(b byte) bool {
	return (b <= 57 && b >= 48) || (b >= 97 && b <= 122)
}
