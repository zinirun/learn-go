package theories

func Sum[K comparable, V int | float32](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
