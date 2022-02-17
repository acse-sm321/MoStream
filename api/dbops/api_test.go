package dbops

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

// init dbConn, truncate tables, unit tests, clear truncate tables

var tempvid string

// clear data cache
func clearTables() {
	dbConn.Exec("TRUNCATE  users")
	dbConn.Exec("TRUNCATE  video_info")
	dbConn.Exec("TRUNCATE  comments")
	dbConn.Exec("TRUNCATE  sessions")
}

// Test main
func TestMain(m *testing.M) {
	clearTables()
	m.Run()
	clearTables()
}

// Test API for user services
func TestUserWorkFlow(t *testing.T) {
	t.Run("Add", testAddUserCredential)
	t.Run("Get", testGetUserCredential)
	t.Run("Del", testDeleteUser)
	t.Run("Reget", testReGetUserCredential)
}

// Test adding new user credential
func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("mossi", "123")
	if err != nil {
		t.Errorf("Error of add user credential: %v", err)
	}
}

// Test getting user credential from database
func testGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("mossi")
	if pwd != "123" || err != nil {
		t.Errorf("Error of get user credential: %v", err)
	}
}

// Test delete a user entry
func testDeleteUser(t *testing.T) {
	err := DeleteUser("mossi", "123")
	if err != nil {
		t.Errorf("Error of delete user credential: %v", err)
	}
}

//  Test re-get user information to see whether successfully deleted
func testReGetUserCredential(t *testing.T) {
	pwd, err := GetUserCredential("mossi")
	if err != nil {
		t.Errorf("Error of re-get user credential: %v", err)
	}

	if pwd != "" {
		t.Errorf("Deleting user test failed")
	}
}

// Test add/delete videos
func TestVideoWorkFlow(t *testing.T) {
	clearTables()
	t.Run("PrepareUser", testAddUserCredential)
	t.Run("AddVideo", testAddVideoInfo)
	t.Run("GetVideo", testGetVideoInfo)
	t.Run("DelVideo", testDeleteVideoInfo)
	t.Run("RegetVideo", testRegetVideoInfo)
}

func testAddVideoInfo(t *testing.T) {
	vi, err := AddNewVideo(1, "my-video")
	if err != nil {
		t.Errorf("Error of AddVideoInfo: %v", err)
	}
	tempvid = vi.Id
}

func testGetVideoInfo(t *testing.T) {
	_, err := GetVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of GetVideoInfo: %v", err)
	}
}

func testDeleteVideoInfo(t *testing.T) {
	err := DeleteVideoInfo(tempvid)
	if err != nil {
		t.Errorf("Error of DeleteVideoInfo: %v", err)
	}
}

func testRegetVideoInfo(t *testing.T) {
	vi, err := GetVideoInfo(tempvid)
	if err != nil || vi != nil {
		t.Errorf("Error of RegetVideoInfo: %v", err)
	}
}

func TestComments(t *testing.T) {
	clearTables()
	t.Run("AddUser", testAddUserCredential)
	t.Run("AddComments", testAddComments)
	t.Run("ListComments", testListComments)
}

func testAddComments(t *testing.T) {
	vid := "12345"
	aid := 1
	content := "I like this video"

	err := AddNewComments(vid, aid, content)

	if err != nil {
		t.Errorf("Error of AddComments: %v", err)
	}
}

func testListComments(t *testing.T) {
	vid := "12345"
	from := 1514764800
	to, _ := strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000, 10))

	res, err := ListComments(vid, from, to)
	if err != nil {
		t.Errorf("Error of ListComments: %v", err)
	}

	for i, ele := range res {
		fmt.Printf("comment: %d, %v \n", i, ele)
	}
}
