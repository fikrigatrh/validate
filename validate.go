package gtr_validate

import "regexp"

func Num(s ...string) bool {
	var regex, _ = regexp.Compile(`[a-z]+|[A-Z]+|\s|\W`)
	for _, i2 := range s {
		var isMatch = regex.MatchString(i2)
		if isMatch == true{
			return true
		}
	}
	return false
}

func Space(s ...string) bool {
	var regex, _ = regexp.Compile(`\s`)
	for _, i2 := range s {
		var isMatch = regex.MatchString(i2)
		if isMatch == true{
			return true
		}
	}
	return false
}
