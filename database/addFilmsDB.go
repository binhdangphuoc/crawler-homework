package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	crawFilm "project1/crawler"
	film "project1/model"
)

func AddData(url string) {
	db, err := gorm.Open(sqlite.Open("test"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("connect mysql cuccess")
	err = db.AutoMigrate(&film.Film{})
	if err != nil {
		panic("failed to create table")
	}
	fmt.Println("create table oke")
	data, err := crawFilm.GetInfoPage(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(len(data))

	for i := range data {
		db.Create(&film.Film{Id: data[i].Id, Rank: data[i].Rank, Name: data[i].Name, Rating: data[i].Rating})
	}

}
