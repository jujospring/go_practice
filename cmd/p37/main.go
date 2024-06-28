package main

import "errors"

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	elem, ok := users[name]
	if !ok {
		var err error = errors.New("not found")
		return false, err
	} else if !elem.scheduledForDeletion {
		return false, nil
	} else {
		delete(users, name)
		return true, nil
	}

}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
