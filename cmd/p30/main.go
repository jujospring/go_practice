package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planFree {
		mySlice := messages[:2]
		return mySlice, nil
	} else if plan == planPro {
		mySlice := messages[:]
		return mySlice, nil
	} else {
		var err error = errors.New("unsupported plan")
		return nil, err
	}
}
