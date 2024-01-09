package collection

func Filter[T any](elms []T, fn func(T) bool) []T {
	outputs := make([]T, 0)
	for _, elm := range elms {
		if fn(elm) {
			outputs = append(outputs, elm)
		}
	}
	return outputs
}

func Reverse[T any](elems []T) []T {
	outputs := make([]T, len(elems))
	for i, elm := range elems {
		outputs[len(elems)-1-i] = elm
	}
	return outputs
}
