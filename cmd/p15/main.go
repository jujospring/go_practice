package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed {
		return 2 * (len(e.body))
	} else {
		return 5 * (len(e.body))
	}
}

func (e email) format() string {
	// return fmt.Sprintf("'%s' | ", e.body)
	formatted := fmt.Sprintf("'%s' | ", e.body)
	if e.isSubscribed {
		return formatted + "Subscribed"
	} else {
		return formatted + "Not Subscribed"
	}
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
