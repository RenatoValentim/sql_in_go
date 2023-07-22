package main

import (
	"context"
	"database/sql"
	"log"
	"sql_in_go/internal/db"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConn, err := sql.Open(`sqlite3`, `./example.db`)
	if err != nil {
		panic(err.Error())
	}
	defer dbConn.Close()

	dt := db.New(dbConn)

	ctx := context.Background()

	cup := db.CreateUserParams{
		ID:   int64(uuid.New().ID()),
		Name: "John Doe",
	}
	err = dt.CreateUser(ctx, cup)
	if err != nil {
		log.Printf("Failed to create user. %v\n", err)
		return
	}

	users, err := dt.GetUsers(ctx)
	if err != nil {
		log.Printf("Failed to get users. %v\n", err)
		return
	}

	log.Printf("users: %v\n", users)
}
