package main

import (
	crawler "project1/crawler"
)

func main() {
	//url := "https://www.imdb.com/chart/top/?ref_=nv_mv_250"
	//database.AddData(url)

	//db, err := gorm.Open(sqlite.Open("test"), &gorm.Config{})
	//if err != nil {
	//	panic("failed to connect database")
	//}
	//fmt.Println("connect mysql cuccess")
	//var f film.Film
	//db.First(&f,"id=250")
	//fmt.Println(f)

	url2 := "https://template-homedecor.onshopbase.com/collections/new-arrivals?sortby=price%3Aasc"
	crawler.GetProducts(url2)
}
