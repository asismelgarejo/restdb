package restdb

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	ID        int
	Username  string
	Password  string
	LastLogin int64
	Admin     int
	Active    int
}

// PostgreSQL Connection details
var (
	Hostname = "postgres"
	Port     = 5432
	Username = "mtsouk"
	Password = "pass"
	Database = "restapi"
)

func connectPostgres() *sql.DB {
	conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Hostname, Port, Username, Password, Database)

	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Println(err)
		return nil
	}

	return db
}

func AdminPass() error {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
	}
	defer db.Close()

	return nil
}

// DeleteUser is for deleting a user defined by ID
func DeleteUser(ID int) bool {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return false
	}
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM users WHERE ID = ?")
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}

	_, err = stmt.Exec(ID)
	if err != nil {
		log.Println("DeleteUser:", err)
		return false
	}

	return true
}

// ReturnAllUsers is for returning all users from the database table
func ReturnAllUsers() []User {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users \n")
	if err != nil {
		log.Println(err)
		return nil
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		temp := User{c1, c2, c3, c4, c5, c6}
		all = append(all, temp)
	}

	log.Println("All:", all)
	return all
}

// FindUserID is for returning a user record defined by ID
func FindUserID(ID int) User {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where ID = $1 \n", ID)
	if err != nil {
		log.Println("Query:", err)
		return User{}
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}
		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}
	return u
}

// FindUserUsername is for returning a user record defined by username
func FindUserUsername(username string) User {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return User{}
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users where Username = $1 \n", username)
	if err != nil {
		log.Println("FindUserUsername Query:", err)
		return User{}
	}
	defer rows.Close()

	u := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return User{}
		}
		u = User{c1, c2, c3, c4, c5, c6}
		log.Println("Found user:", u)
	}
	return u
}

// ReturnLoggedUsers is for returning all logged in users
func ReturnLoggedUsers() []User {
	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return nil
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Active = 1 \n")
	if err != nil {
		log.Println(err)
		return nil
	}

	all := []User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return []User{}
		}
		temp := User{c1, c2, c3, c4, c5, c6}
		log.Println("temp:", all)
		all = append(all, temp)
	}

	log.Println("Logged in:", all)
	return all
}

// IsUserAdmin determines whether a user is
// an administrator or not
func IsUserAdmin(u UserPass) bool {
	err := u.Validate()
	if err != nil {
		log.Println("IsUserAdmin - Validate:", err)
		return false
	}

	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Username = $1 \n", u.Username)
	if err != nil {
		log.Println(err)
		return false
	}

	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	// If there exist multiple users with the same username,
	// we will get the FIRST ONE only.
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return false
		}
		temp = User{c1, c2, c3, c4, c5, c6}
	}

	if u.Username == temp.Username && u.Password == temp.Password && temp.Admin == 1 {
		return true
	}
	return false
}

func IsUserValid(u UserPass) bool {
	err := u.Validate()
	if err != nil {
		log.Println("IsUserValid - Validate:", err)
		return false
	}

	db := connectPostgres()
	if db == nil {
		fmt.Println("Cannot connect to PostgreSQL!")
		db.Close()
		return false
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM users WHERE Username = $1 \n", u.Username)
	if err != nil {
		log.Println(err)
		return false
	}

	temp := User{}
	var c1 int
	var c2, c3 string
	var c4 int64
	var c5, c6 int

	// If there exist multiple users with the same username,
	// we will get the FIRST ONE only.
	for rows.Next() {
		err = rows.Scan(&c1, &c2, &c3, &c4, &c5, &c6)
		if err != nil {
			log.Println(err)
			return false
		}
		temp = User{c1, c2, c3, c4, c5, c6}
	}

	if u.Username == temp.Username && u.Password == temp.Password {
		return true
	}
	return false
}

// Validate method validates the data of UserPass
func (p *UserPass) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}
