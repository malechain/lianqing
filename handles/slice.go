package handles

import "strings"

// RemoveDuplicates 切片去重
func RemoveDuplicates(a []string) (ret []string) {
	aLen := len(a)
	for i := 0; i < aLen; i++ {
		if (i > 0 && strings.ToLower(a[i-1]) == strings.ToLower(a[i])) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}
