package main

import (
	"database/sql"
	"fmt"
	"log"
	"golang/storage"
	_ "github.com/lib/pq"
)

const (
	user = "postgres"
	password = "1234"
	host = "localhost"
	port = 5432
	dbname = "blog"
)

func main(){
	connstr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	_, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatalf("Failed to open connection: %v", err)
	}
	// dbManager := storage.NewDBManager(db)
	// err = InsertFakeDate(dbManager)
	// if err != nil {
	// 	log.Fatalf("Failed while creating blog: %v", err)
	// }
	// blogs, err := dbManager.GetAll(&storage.GetBlogsQueryParam{
	// 	Title: "Consequatur",
	// 	Limit: 20, 
	// 	Page: 2,
	// })
	// if err != nil {
	// 	log.Fatalf("Failed to get all blogs: %v", err)
	// }

	// PrintBlogs(blogs)
}

func PrintBlogs(blogs []*storage.Blog) {
	for _, blog := range blogs {
		PrintBlog(blog)
	}
}

func PrintBlog(blog *storage.Blog) {
	fmt.Println("-------------- Blog --------------")
	fmt.Println("Id:", blog.Id)
	fmt.Println("Title", blog.Title)
	fmt.Println("Description:", blog.Description)
	fmt.Println("Author:", blog.Author)
	fmt.Println("Created At:", blog.CreatedAt.Format("2006-01-02 15:04:03"))
	fmt.Println("----------------------------------")
}