package algorithms

type Ordered interface {
	~string | ~int | ~uint | ~int8
}

func Min[T Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}
