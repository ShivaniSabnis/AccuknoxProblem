package models

import (
	"log"
	"time"
)

type User struct {
	Name       string    `json:"name"`
	Email      string    `json:"email" validate:"required,email"`
	Password   string    `json:"password" validate:"required"`
	Sid        string    `json:"sid"`
	ExpiryTime time.Time `json:"-"`
}

type UserStore struct {
	Users []*User
}

func (u *UserStore) CheckExistingUser(email string) bool {
	if u == nil {
		return false
	}
	for _, user := range u.Users {
		if user.Email == email {
			return true
		}
	}
	return false
}

func (u *UserStore) CheckIfSignedUp(email string) (bool, *User) {
	if u == nil {
		return false, nil
	}
	for _, user := range u.Users {
		if user.Email == email {
			return true, user
		}
	}
	return false, nil
}

func (u *UserStore) UpdateUserDetails(user *User) {
	for i, userDetails := range u.Users {
		if userDetails.Email == user.Email {
			u.Users[i].ExpiryTime = user.ExpiryTime
			u.Users[i].Sid = user.Sid
		}
	}
}

func (u *UserStore) CheckSidExpired(sid string) bool {
	if u == nil {
		log.Println("Entered nil")
		return true
	}
	for _, user := range u.Users {
		if user.Sid == sid {
			if user.ExpiryTime.After(time.Now()) {
				return false
			}
		}
	}
	return true
}
