package day4

import "strconv"

func ValidPasswords(min, max int) []string {
	var passwords []string

	for passNum := min; passNum <= max; passNum++ {
		passStr := strconv.Itoa(passNum)
		if IsValidPassword(passStr) {
			passwords = append(passwords, passStr)
		}
	}

	return passwords
}

func IsValidPassword(s string) bool {
	lastNum := rune(s[0])
	var hasDouble bool
	for _, c := range s[1:] {
		if c < lastNum {
			return false
		}
		if c == lastNum {
			hasDouble = true
		}
		lastNum = c
	}

	return hasDouble
}
