// export DATABASE_URL=postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable

/*
docker run \
  -d \
  -e POSTGRES_HOST_AUTH_METHOD=trust \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=postgres \
  -p 5432:5432 \
  postgres:12.5-alpine
 */

/*
CREATE TABLE users (
  id SERIAL PRIMARY KEY,
  age INT,
  first_name TEXT,
  last_name TEXT,
  email TEXT UNIQUE NOT NULL
);
INSERT INTO users (age, email, first_name, last_name)
VALUES (30, 'jon@calhoun.io', 'Jonathan', 'Calhoun');
INSERT INTO users (age, email, first_name, last_name)
VALUES (52, 'bob@smith.io', 'Bob', 'Smith');
INSERT INTO users (age, email, first_name, last_name)
VALUES (15, 'jerryjr123@gmail.com', 'Jerry', 'Seinfeld');

 */

package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v4/stdlib"
)

type User struct {
	ID        int
	Age       int
	FirstName string
	LastName  string
	Email     string
}

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	var greeting string
	err = db.QueryRow("select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(greeting)

	sqlStatement := `SELECT id, email FROM users WHERE id=$1;`
	var email string
	var id int
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	row := db.QueryRow(sqlStatement, 3)
	switch err := row.Scan(&id, &email); err {
	case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
	case nil:
		fmt.Println(id, email)
	default:
		panic(err)
	}

	sqlStatement = `SELECT * FROM users;`

	linhas, _ := db.Query(sqlStatement)

	var users []User

	for linhas.Next() {

		var user User

		err = linhas.Scan(&user.ID, &user.Age, &user.FirstName,
			&user.LastName, &user.Email)
		
		users = append(users,user)
		
	}

	for _,usuario := range users {
		fmt.Println("Usuario: ",usuario.FirstName," - ",usuario.Email)
	}

}


