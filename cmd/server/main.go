package main

import (
	"go-api-crud/configs"
	"go-api-crud/internal/entity"
	"go-api-crud/internal/infra/database"
	"go-api-crud/internal/infra/database/web/handler"
	"net/http"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	_, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&entity.Product{}, &entity.User{})

	productDb := database.NewProduct(db)
	productHandler := handler.NewProductHandler(productDb)

	http.HandleFunc("/products", productHandler.CreateProduct)
	
	http.ListenAndServe(":8080", nil)
}