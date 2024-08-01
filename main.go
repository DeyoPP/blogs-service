package main

import (
	"log"
	"net/http"
	"blogs/model"
	"blogs/repo"
	"blogs/service"
	"blogs/handler"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func initDB() *gorm.DB {
	connection_url := "user=postgres password=super dbname=SOA port=5432 sslmode=disable"
	database, err := gorm.Open(postgres.Open(connection_url), &gorm.Config{})

	if err != nil {
		print(err)
		return nil
	}
	database.AutoMigrate(&model.Blog{})

	return database
}

func startServer(blogHandler *handler.BlogHandler) {
	router := mux.NewRouter()

	router.HandleFunc("/blogs", blogHandler.CreateBlog).Methods("POST")
	router.HandleFunc("/",  writte)

	println("Server starting")
	log.Fatal(http.ListenAndServe(":8081", router))
}

func writte(w http.ResponseWriter, r *http.Request) {
	println("test")
}

func main() {
	database := initDB()
	if database == nil {
		print("Failed to connect to database!")
		return
	}

	blogRepo := &repo.BlogRepository{DatabaseConnection: database}
	blogService := &service.BlogService{BlogRepo: blogRepo}
	blogHandler := &handler.BlogHandler{BlogService: blogService}

	startServer(blogHandler)
}
