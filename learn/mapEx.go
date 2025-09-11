package main

import (
	"errors"
	"fmt"
)

func inMapEx() {
	ex := make(map[string](int))
	fmt.Println(ex)
	return
}

type userex struct {
	name        string
	phoneNumber int
}

func getUserMap(names []string, phoneNumbers []int) (map[string]userex, error) {
	// ?
	nLen := len(names)
	phLen := len(phoneNumbers)

	if nLen != phLen {
		return nil, errors.New("length mismatch")
	}

	res := make(map[string](userex))

	for i := 0; i < nLen; i++ {
		res[names[i]] = userex{
			name:        names[i],
			phoneNumber: phoneNumbers[i],
		}
	}
	return res, nil

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}

func exdelete() {
	users := make(map[string]user)
	users["roshin"] = user{
		name:                 "roshin",
		number:               1,
		scheduledForDeletion: true,
	}

	users["charlie"] = user{
		name:                 "charlie",
		number:               2,
		scheduledForDeletion: false,
	}

	fmt.Println(deleteIfNecessary(users, "roshin"))
	fmt.Println(deleteIfNecessary(users, "charlie"))
	fmt.Println(deleteIfNecessary(users, "roshin"))
}

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
	elem, ok := users[name]
	if !ok {
		return false, errors.New(" not found ")
	}

	if !elem.scheduledForDeletion {
		return false, nil
	}
	delete(users, name)
	return true, nil

}
