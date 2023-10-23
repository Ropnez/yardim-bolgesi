package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Content, City, Distict, Adress, Number, Description string
	Id                                                  int
}

func (list List) Migrate() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.AutoMigrate(&List{})
}

func (list List) Add() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Create(&list)
}

func (list List) Get(where ...interface{}) List {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return List{}
	}
	db.First(&list, where...)
	return list
}

func (list List) GetAll(where ...interface{}) []List {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return []List{}
	}
	var lists []List
	db.Find(&lists, where...)
	return lists
}

func (list List) Update(column string, value interface{}) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&list).Update(column, value)
}

func (list List) Updates(data List) {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Model(&list).Updates(data)
}

func (list List) Delete() {
	db, err := gorm.Open(mysql.Open(Dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	db.Delete(&list, list.Id)
}
