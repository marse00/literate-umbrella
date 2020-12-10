package middleware

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"log"
)

var (
	ctx      context.Context
	database *sql.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "0000"
	dbname   = "postgres"
)

func OpenDatabase() error {
	fmt.Println("OpenDatabase", database)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error

	database, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println("sql.Open error:", err)
	}
	fmt.Println("OpenDatabase", database)

	err = database.Ping()
	if err != nil {
		log.Println("database.Ping error:", err)
	}

	return err
}

//IsUsernameTaken checks if a username already exists
//It returns a bool depending on the result and any error encountered
func IsUsernameTaken(username string) bool {
	fmt.Println("isUsernameTaken", database)

	sqlQuery := "SELECT id FROM public.\"User\" WHERE name = $1 ;"
	log.Println("Called query row context")
	log.Println(sqlQuery, username)
	var id string
	err := database.QueryRow(sqlQuery, username).Scan(&id)
	log.Println("finished query row context", err, "id:", id)

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		log.Fatalf("Query error in middleware.isUsernameTaken: %v\n", err)
		return true
	default:
		return true
	}
}

func NewUser(username string, email string, passwordHash []byte) error {
	sqlQuery := "INSERT INTO public.\"User\" (id, name, email, passwd) VALUES ($1, $2, $3, $4)"
	userID := uuid.New()
	res, err := database.Exec(sqlQuery, userID, username, email, passwordHash)
	fmt.Println("Err NewUser:", err, "Result: ", res)
	return err
}

