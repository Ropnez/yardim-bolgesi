package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type UserUpdate struct {
	gorm.Model
	Content, City, Distict, Adress, Number, Description string
	Id                                                  int
	Update_Id                                           int
}

func (userUpdate UserUpdate) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.AutoMigrate(&UserUpdate{})
}

func (userUpdate UserUpdate) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Create(&userUpdate)
}

func (userUpdate UserUpdate) Get(where ...interface{}) UserUpdate {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return UserUpdate{}
	}
	db.First(&userUpdate, where...)
	return userUpdate
}

func (userUpdate UserUpdate) GetAll(where ...interface{}) []UserUpdate {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return []UserUpdate{}
	}
	var userUpdates []UserUpdate
	db.Find(&userUpdates, where...)
	return userUpdates
}

func (userUpdate UserUpdate) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&userUpdate).Update(column, value)
}

func (userUpdate UserUpdate) Updates(data UserUpdate) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&userUpdate).Updates(data)
}

func (userUpdate UserUpdate) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Delete(&userUpdate, userUpdate.Id)
}
