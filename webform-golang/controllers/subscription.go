package controllers

import "webform-golang/db"

type Subscription struct {
	Name  string
	Email string
}

func Create(name, email string) error {
	newSub := Subscription{Name: name, Email: email}

	return db.Insert("newsletter", newSub)
}
