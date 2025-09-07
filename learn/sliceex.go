package main

import "errors"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	} else if plan == planFree {
		return messages[:2], nil
	} else {
		return nil, errors.New("Unsupported Plan")
	}

}
