package logic

import "regexp"

func IsStringOrNum(ip string) (b bool) {
	if m, _ := regexp.MatchString("[a-zA-Z0-9]+", ip); !m {
		return false
	}
	return true
}

func IsPhone(phone string) (b bool) {
	if m, _ := regexp.MatchString("[0-9]+", phone); !m {
		return false
	}
	return true
}
