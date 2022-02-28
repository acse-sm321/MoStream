package dbops

import (
	"MoStream/api/defs"
	"MoStream/api/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"time"
)

// Database operations
// if we open multiple conn here it might lead to close-wait issue

// AddUserCredential Add new user credential to database
func AddUserCredential(loginName string, pwd string) error {
	stmIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES  (?,?)")
	if err != nil {
		return err
	}
	_, err = stmIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmIns.Close()
	return nil
}

// GetUserCredential Get the password of user by its username
func GetUserCredential(loginName string) (string, error) {
	stmOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
	if err != nil {
		log.Printf("%scripts", err)
		return "", err
	}
	var pwd string
	err = stmOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmOut.Close()

	return pwd, nil
}

// DeleteUser Delete a user by its username and password
func DeleteUser(loginName string, pwd string) error {
	stmDel, err := dbConn.Prepare("DELETE  FROM  users WHERE login_name=? AND pwd=? ")
	if err != nil {
		log.Printf("%scripts", err)
		return err
	}
	_, err = stmDel.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	stmDel.Close()
	return nil
}

// AddNewVideo Add new video data to database
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	// create uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}

	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05") // MM DD YY, hh:mm:ss
	stmIns, err := dbConn.Prepare("INSERT INTO video_info (id,author_id,name,display_ctime) VALUES(?,?,?,?)")
	if err != nil {
		return nil, err
	}

	_, err = stmIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: ctime}
	defer stmIns.Close()

	return res, nil
}

func GetVideoInfo(vid string) (*defs.VideoInfo, error) {
	stmtOut, err := dbConn.Prepare("SELECT author_id, name, display_ctime FROM video_info WHERE id=?")

	var aid int
	var dct string
	var name string

	err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dct)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmtOut.Close()

	res := &defs.VideoInfo{Id: vid, AuthorId: aid, Name: name, DisplayCtime: dct}

	return res, nil
}

func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE FROM video_info WHERE id=?")
	if err != nil {
		return err
	}

	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}

	defer stmtDel.Close()
	return nil
}

func AddNewComments(vid string, aid int, content string) error {
	id, err := utils.NewUUID()
	if err != nil {
		return err
	}

	stmtIns, err := dbConn.Prepare("INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return err
	}

	defer stmtIns.Close()
	return nil
}

func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(` SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)
		ORDER BY comments.time DESC`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}

		c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
		res = append(res, c)
	}
	defer stmtOut.Close()

	return res, nil
}
