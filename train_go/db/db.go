package db

import (
	"fmt"
	"log"
	"time"

	"database/sql"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Username  string
	Email     string
	CreatedAt time.Time
}

func Connect() *sql.DB {
	conn, err := sql.Open("postgres", "host=localhost port=5432 user=admin password=admin dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	err = conn.Ping()
	if err != nil {
		log.Fatal("Exit 1: %v\n", err)
	}

	fmt.Println("Connected to PostgreSQL!")
	return conn
}

func List(conn *sql.DB) []User {
	rows, err := conn.Query("Select * from users")
	if err != nil {
		log.Fatal("Exit 2: %v\n", err)
	}

	defer rows.Close()
	for rows.Next() {
		var user User
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		if err != nil {
			log.Fatal("Exit 3: %v\n", err)
		}
		fmt.Println(user.ID, user.Username, user.Email, user.CreatedAt)
	}
	return nil
}

func Insert(conn *sql.DB, user User) {
	_, err := conn.Exec("Insert into users(username, email,created_at) values($1, $2,$3)", user.Username, user.Email, user.CreatedAt)
	if err != nil {
		log.Fatal("Exit 4: %v\n", err)
	}
	fmt.Println("User inserted successfully")
}

func Delete(conn *sql.DB, id int) {
	_, err := conn.Exec("Delete from users where id = $1", id)
	if err != nil {
		log.Fatal("Exit 5: %v\n", err)
	}
	fmt.Println("User deleted successfully")
}
