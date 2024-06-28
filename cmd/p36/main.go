package main

import "errors"

func getUserMap(names []string, phoneNumbers []int) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		var err error = errors.New("invalid sizes")
		return nil, err
	}
	users := make(map[string]user)
	for i := 0; i < len(names); i++ {
		var u user
		u.name = names[i]
		u.phoneNumber = phoneNumbers[i]
		users[names[i]] = u
	}
	return users, nil
}

type user struct {
	name        string
	phoneNumber int
}
