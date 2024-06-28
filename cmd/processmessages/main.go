package main

func processMessages(messages []string) []string {
	num := len(messages)
	doneMessages := make([]string, num)
	mChan := make(chan int, num)
	doneChan := make(chan struct{}, num)

	for i := range messages {
		go func(index int) {
			doneMessages[index] = messages[index] + "-processed"
			mChan <- index
			doneChan <- struct{}{}
		}(i)
	}

	for i := 0; i < num; i++ {
		<-doneChan
	}

	return doneMessages
}
