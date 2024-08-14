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
	// Retrieve database connection details from environment variabless
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	dbname := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	// Construct the connection URL
	connectionURL := "user=" + user + " password=" + password + " dbname=" + dbname + " host=" + host + " port=" + port + " sslmode=disable"
	database, err := gorm.Open(postgres.Open(connectionURL), &gorm.Config{})

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
