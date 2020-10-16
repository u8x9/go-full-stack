package split_string

import "strings"

func Split(str, sep string) []string {
	idx := strings.Index(str, sep)
    // 优化前：112 B/op	       3 allocs/op
    // 优化后：64 B/op	       1 allocs/op
	ss := make([]string, 0, strings.Count(str, sep)+1)
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
