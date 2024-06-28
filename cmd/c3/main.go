package main

func addEmailsToQueue(emails []string) chan string {
	cap := len(emails)
	ch := make(chan string, cap)
	for i := 0; i < cap; i++ {
		ch <- emails[i]
	}
	return ch
}
