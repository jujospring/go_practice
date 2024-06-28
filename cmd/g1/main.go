package main

func getLast[T any](s []T) T {
	if len(s) == 0 {
		var none T
		return none
	}
	last := s[len(s)-1]
	return last
}
