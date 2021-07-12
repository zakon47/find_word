package find_word

import (
	"errors"
	"strings"
)

type modeType byte

const (
	prefix modeType = 0
	suffix modeType = 1
)

var ErrFindWord = errors.New("no match found")

//найти 1 из вхождений с начала строки
func HasPrefix(str string, list []string) (string, string, error) {
	return hasPrefix(str, list, prefix, false)
}

//найти 1 из вхождений с конца строки
func HasSuffix(str string, list []string) (string, string, error) {
	return hasPrefix(str, list, suffix, false)
}

//найти 1 из вхождений с начала строки - без учета регистра
func HasPrefixWithoutCase(str string, list []string) (string, string, error) {
	return hasPrefix(str, list, prefix, true)
}

//найти 1 из вхождений с конца строки - без учета регистра
func HasSuffixWithoutCase(str string, list []string) (string, string, error) {
	return hasPrefix(str, list, suffix, true)
}

func hasPrefix(str string, list []string, mode modeType, lower bool) (string, string, error) {
	if len(list) == 0 {
		return "", "", ErrFindWord
	}
	asRunes := []rune(str)
	for _, value := range list {
		if len(value) == 0 {
			continue
		}
		if mode == prefix {
			if lower {
				if strings.HasPrefix(strings.ToLower(str), strings.ToLower(value)) {
					return str[:len(value)], str[len(value):], nil
				}
			} else {
				if strings.HasPrefix(str, value) {
					return str[:len(value)], str[len(value):], nil
				}
			}
		} else {
			if lower {
				if strings.HasSuffix(strings.ToLower(str), strings.ToLower(value)) {
					lenWord := []rune(value)
					return string(asRunes[len(asRunes)-len(lenWord):]), string(asRunes[:len(asRunes)-len(lenWord)]), nil
				}
			} else {
				if strings.HasSuffix(str, value) {
					lenWord := []rune(value)
					return string(asRunes[len(asRunes)-len(lenWord):]), string(asRunes[:len(asRunes)-len(lenWord)]), nil
				}
			}
		}
	}
	return "", "", ErrFindWord
}
