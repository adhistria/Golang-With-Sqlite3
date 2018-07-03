package main

import (
	"database/sql"
	"fmt"
)

type User struct {
	Id       int    `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
	Name     string `json:"name,omitempty"`
}

func (u *User) getUser(db *sql.DB) error {
	statement := fmt.Sprintf("SELECT * FROM USERS WHERE ID = '%d", u.Id)
	return db.QueryRow(statement).Scan(&u.Id, &u.Username, &u.Name)
}

func (u *User) getUsers(db *sql.DB) ([]User, error) {
	statement := fmt.Sprintf("SELECT * FROM USERS")
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	users := []User{}
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.Id, &u.Username, &u.Name); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (u *User) CreateUser(db *sql.DB) error {
	statement := fmt.Sprintf("INSERT INTO USERS(username,name) VALUES(%s,%d)", u.Username, u.Name)
	_, err := db.Exec(statement)
	if err != nil {
		return err
	}
	err = db.QueryRow("SELECT LAST_INSERT_ID()").Scan(&u.Id)
	fmt.Println("err", err)
	if err != nil {
		return err
	}
	return nil

}
func (u *User) UpdateUser(db *sql.DB) error {
	statement := fmt.Sprintf("UPDATE USERS SET USERNAME = '%s' NAME = '%s' WHERE ID= '%id", u.Username, u.Name, u.Id)
	_, err := db.Exec(statement)
	return err
}
func (u *User) DeleteUser(db *sql.DB) error {
	statement := fmt.Sprintf("DELETE FROM USERS WHERE ID = %d", u.Id)
	_, err := db.Exec(statement)
	return err
}
