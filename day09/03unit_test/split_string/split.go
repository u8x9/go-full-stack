package split_string

import "strings"

func Split(str, sep string) []string {
	idx := strings.Index(str, sep)
	ss := make([]string, 0)
	for idx >= 0 {
		item := str[:idx]
		if len(item) > 0 {
			ss = append(ss, item)
		}
		str = str[idx+len(sep):]
		idx = strings.Index(str, sep)
	}
	if len(str) > 0 {
		ss = append(ss, str)
	}
	return ss
}
