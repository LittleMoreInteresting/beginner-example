package main

type IntNumber interface {
	int | int8 | int32 | int64
}

func Min[T IntNumber](a, b T) T {
	if a > b {
		return b
	}
	return a
}
