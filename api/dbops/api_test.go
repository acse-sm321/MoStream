package dbops

import (
	"testing"
)

// init dbConn, truncate tables, unit tests, clear truncate tables

// clear all test conn
func clearTables() {
	dbConn.Exec("TRUNCATE  users")
	dbConn.Exec("TRUNCATE  video_info")
	dbConn.Exec("TRUNCATE  comments")
	dbConn.Exec("TRUNCATE  sessions")
}

func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", TestAddUserCredential)
	t.Run("Get", TestGetUserCredential)
	t.Run("Del", TestDeleteUser)
	t.Run("Reget", TestReGetUserCredential)
}

func TestAddUserCredential(t *testing.T) {
	err := AddUserCredential("mossi", "123")
	if err != nil {
		t.Errorf("Error of add user credential: %v", err)
	}
}

func TestGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("mossi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of get user credential: %v", err)
	}
}

func TestDeleteUser(t *testing.T) {
	err := DeleteUser("mossi", "123")
	if err != nil {
		t.Errorf("Error of delete user credential: %v", err)
	}
}

func TestReGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("mossi")
	if err != nil {
		t.Errorf("Error of re-get user credential: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}
