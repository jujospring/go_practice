package main

func maxMessages(thresh int) int {
	messages := 0

	for i := 0; ; i++ {
		messages += (100 + i*1)
		if messages > thresh {
			return i
		}
	}

}
