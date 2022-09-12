package main

import (
	"log"
	"net/http"
	"simple-mvc/controller"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:rootPassword!@tcp(127.0.0.1:3306)/golang-mvc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if (err != nil) {
		log.Fatal(err.Error())
	}

	log.Println("Successfully connected to database!")

	instanceRepository := controller.NewRepository(db)

	mux := http.NewServeMux()

	mux.HandleFunc("/add-new-data", controller.GetAddDataFormHandler)
	mux.HandleFunc("/post-add-data", instanceRepository.ProcessAddDataFormHandler)

	fileServer := http.FileServer(http.Dir("assets"))
	mux.Handle("/assets/", http.StripPrefix("/assets", fileServer))

	log.Println("Server connected at http://localhost:5000")

	error := http.ListenAndServe(":5000", mux)

	log.Fatal(error)
}
