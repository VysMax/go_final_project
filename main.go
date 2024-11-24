package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"VysMax/DBManip"
	"VysMax/database"
	"VysMax/handlers"

	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	_ "modernc.org/sqlite"
)

func auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		pass := os.Getenv("TODO_PASSWORD")
		if len(pass) > 0 {
			var jwtFromCookie string

			cookie, err := r.Cookie("token")
			if err == nil {
				jwtFromCookie = cookie.Value
			}

			secret := []byte("peacock")

			jwtToken := jwt.New(jwt.SigningMethodHS256)

			signedToken, err := jwtToken.SignedString(secret)
			if err != nil {
				http.Error(w, "failed to sign jwt", http.StatusUnauthorized)
			}

			if signedToken != jwtFromCookie {
				http.Error(w, "Authentification required", http.StatusUnauthorized)
				return
			}
		}
		next(w, r)
	})
}

func main() {
	err := godotenv.Load()
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
	case false:
		db = database.CreateDB(os.Getenv("TODO_DBFILE"))
	case true:
		db = database.ConnectDB(os.Getenv("TODO_DBFILE"))
	}

	repo := DBManip.NewRepository(db)
	handler := handlers.NewHandler(repo)
	defer db.Close()

	http.Handle("/*", http.FileServer(http.Dir("./web")))
	http.HandleFunc("/api/nextdate", handler.GetNextDate)
	http.HandleFunc("/api/task", auth(handler.OneTaskHandler))
	http.HandleFunc("/api/tasks", auth(handler.GetTasks))
	http.HandleFunc("/api/task/done", auth(handler.MarkAsDone))
	http.HandleFunc("/api/signin", handler.SignIn)

	err = http.ListenAndServe(":"+os.Getenv("TODO_PORT"), nil)
	if err != nil {
		log.Fatal(err)
	}
}
