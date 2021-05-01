package users

import (
	"database/sql"
	"log"

	"github.com/krish8learn/Graphql_Mysql_GO/internal/pkg/db/database"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"name"`
	Password string `json:"password"`
}

func (user *User) Create() {
	statement, err := database.Db.Prepare("INSERT INTO Users(Username, Password) VALUES (?,?)")

	print(statement)
	if err != nil {
		log.Fatal(err)
	}

	hashedPassword, err := HashPassword(user.Password)
	_, err = statement.Exec(user.Username, hashedPassword)
	if err != nil {
		log.Fatal(err)
	}
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GetUserIdByUsername(username string) (int, error) {
	statement, err := database.Db.Prepare("SELECT ID FROM Users WHERE Username=?")
	if err != nil {
		log.Fatal(err)
	}
	row := statement.QueryRow(username)

	var Id int

	err = row.Scan(&Id)
	if err != nil {
		if err != sql.ErrNoRows {
			log.Print(err)
		}
		return 0, err
	}
	return Id, nil
}

func (user *User) Auhtentication() bool {
	statement, err := database.Db.Prepare("SELECT Password FROM Users WHERE Username = ?")
	if err != nil {
		log.Fatal(err)
	}

	row := statement.QueryRow(user.Username)

	var hashedPassword string

	err = row.Scan(&hashedPassword)

	if err != nil {
		if err == sql.ErrNoRows {
			return false
		} else {
			log.Fatal(err)
		}
	}
	return CheckPasswordHash(user.Password, hashedPassword)
}
