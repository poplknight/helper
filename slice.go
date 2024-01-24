package helper

func Distinct[T any, V comparable](arr []T, fc func(T) V) []T {
	m := make(map[V]struct{}, len(arr)/2)
	res := make([]T, 0, len(arr))
	for _, item := range arr {
		key := fc(item)
		if _, ok := m[key]; ok {
			continue
		}
		m[key] = struct{}{}
		res = append(res, item)
	}
	return res
}
