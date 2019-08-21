package model

import "time"

type User struct {
	Id int
	UserName string
	Password string
	Avatar string
	Website string
	Introduction string
	CreatedAt time.Time
	UpdatedAt time.Time
	Status uint8
	Email string
	Mobile string
	Github string
}

func (u *User) Create() (error) {
	err := db.Create(u).Error
	return err
}

func (u *User) DetailByName(name string) (user *User, err error) {
	err = db.Where("user_name=?", name).First(user).Error
	return
}

func (u *User) DetailById(id int) (user *User, err error) {
	err = db.Find(user, id).Error
	return
}