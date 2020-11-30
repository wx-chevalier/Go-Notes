package mergearray

// MergeArray - 使用 Copy 合并多个数组
func MergeArray(s ...[]interface{}) (slice []interface{}) {

	switch len(s) {
	case 0:
		break
	case 1:
		slice = s[0]
		break
	default:
		s1 := s[0]
		s2 := MergeArray(s[1:]...) //...将数组元素打散
		slice = make([]interface{}, len(s1)+len(s2))
		copy(slice, s1)
		copy(slice[len(s1):], s2)
		break
	}

	return
}
