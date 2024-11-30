package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"VysMax/database"
	"VysMax/handlers"

	"github.com/go-chi/chi"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	portNum := os.Getenv("TODO_PORT")
	pass := os.Getenv("TODO_PASSWORD")

	signedToken, err := handlers.SignToken(pass)
	if err != nil {
		log.Fatal(err)
	}

	appPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	dbFile := filepath.Join(filepath.Dir(appPath), os.Getenv("TODO_DBFILE"))
	_, err = os.Stat(dbFile)

	var db *sql.DB

	switch err == nil {
	case true:
		db = database.ConnectDB(dbFile)
	case false:
		db = database.CreateDB(dbFile)
	}

	repo := database.NewRepository(db)
	handler := handlers.NewHandler(repo)
	defer db.Close()

	r := chi.NewRouter()
	r.Handle("/*", http.FileServer(http.Dir("./web")))

	r.Post("/api/signin", handler.SignIn)
	r.Get("/api/nextdate", handler.GetNextDate)
	r.Get("/api/task", handlers.Auth(handler.GetOneTask, signedToken))
	r.Post("/api/task", handlers.Auth(handler.PostOneTask, signedToken))
	r.Put("/api/task", handlers.Auth(handler.PutOneTask, signedToken))
	r.Delete("/api/task", handlers.Auth(handler.DeleteOneTask, signedToken))
	r.Get("/api/tasks", handlers.Auth(handler.GetTasks, signedToken))
	r.Post("/api/task/done", handlers.Auth(handler.MarkAsDone, signedToken))

	fmt.Printf("Сервер запущен на порту %s\n", portNum)

	err = http.ListenAndServe(":"+portNum, r)
	if err != nil {
		log.Fatal(err)
	}

}
