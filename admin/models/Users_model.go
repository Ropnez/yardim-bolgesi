package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username, Password string
}

func (user User) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.AutoMigrate(&User{})
}

func (user User) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Create(&user)
}

func (user User) Get(where ...interface{}) User {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return User{}
	}
	db.First(&user, where...)
	return user
}

func (user User) GetAll(where ...interface{}) []User {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return []User{}
	}
	var users []User
	db.Find(&users, where...)
	return users
}

func (user User) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&user).Update(column, value)
}

func (user User) Updates(data User) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&user).Updates(data)
}

func (user User) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Delete(&user, user.ID)
}
