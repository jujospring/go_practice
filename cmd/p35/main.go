package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, msgWord := range msg {
		for j := 0; j < len(badWords); j++ {
			if msgWord == badWords[j] {
				return i
			}
		}
	}
	return -1
}
