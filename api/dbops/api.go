package dbops

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

// CURD operations
// if we open multiple conn here it might lead to close-wait issue

// AddUserCredential Add new user credential to database
func AddUserCredential(loginName string, pwd string) error {
	stmIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES  (?,?)")
	if err != nil {
		return err
	}
	stmIns.Exec(loginName, pwd)
	stmIns.Close()
	return nil
}

// GetUserCredential Get the password of user by its username
func GetUserCredential(loginName string) (string, error) {
	stmOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	stmOut.QueryRow(loginName).Scan(&pwd)
	stmOut.Close()

	return pwd, nil
}

// DeleteUser Delete a user by its username and password
func DeleteUser(loginName string, pwd string) error {
	stmDel, err := dbConn.Prepare("DELETE  FROM  users WHERE login_name=? AND pwd=? ")
	if err != nil {
		log.Printf("%s", err)
		return err
	}
	stmDel.Exec(loginName, pwd)
	stmDel.Close()
	return nil
}
