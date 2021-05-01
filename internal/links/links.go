package links

import (
	"log"

	"github.com/krish8learn/Graphql_Mysql_GO/internal/pkg/db/database"
	"github.com/krish8learn/Graphql_Mysql_GO/internal/users"
)

type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

func (link Link) Save() int64 {
	stament, err := database.Db.Prepare("INSERT INTO Links(Title, Address) VALUES (?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stament.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error", err.Error())
	}

	log.Print("row inserted")
	return id
}

func GetAll() []Link {
	statement, err := database.Db.Prepare("SELECT L.ID, L.Title, L.Address, L.UserID, U.Username, FROM Links INNER JOIN Users U ON L.UserID=U.ID")
	if err != nil {
		log.Fatal(err)
	}

	defer statement.Close()
	rows, err := statement.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var username string
	var id string

	var links []Link
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}
