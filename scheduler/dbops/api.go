package dbops

// user -> api -> delete video
// api -> scheduler -> write deletion recr
// timer -> runner -> read wvdr -> exec -> actual delete video from database

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT INTO video_del_rec (vdeo_id) VALUES(?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord Error: %v", err)
		return err
	}
	defer stmtIns.Close()
	return nil
}
