package taskrunner

import (
	"MoStream/scheduler/dbops"
	"MoStream/scheduler/ossops"
	"errors"
	"log"
	"sync"
)

// task can be customized
// delay delete

func deleteVideo(vid string) error {
	//err := os.Remove(VIDEO_PATH + vid)

	// now delete from OSS

	//if err != nil && os.IsNotExist(err) {
	//	log.Printf("Deleting video error: %v", err)
	//	return err
	//}

	ossfn := "videos/" + vid
	bn := "mostream-videos"
	ok := ossops.DeleteObject(ossfn, bn)

	if !ok {
		log.Printf("Deleting video error, oss operation failed")
		return errors.New("Deleting video error")
	}
	return nil
}

func VideoClearDispatcher(dc dataChan) error {
	res, err := dbops.ReadVideoDeletionRecord(3) // how manny records to read
	if err != nil {
		log.Printf("Video clear dispatcher error: %v", err)
		return err
	}

	if len(res) == 0 {
		return errors.New("All tasks finished here.")
	}

	for _, id := range res {
		dc <- id
	}

	return nil
}

func VideoClearExecutor(dc dataChan) error {
	errMap := &sync.Map{}
	var err error
forloop:
	for {
		select {
		case vid := <-dc:
			// could have duplicate rw
			go func(id interface{}) {
				if err := deleteVideo(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
				if err := dbops.DelVideoDeletionRecord(id.(string)); err != nil {
					errMap.Store(id, err)
					return
				}
			}(vid)
		default:
			break forloop
		}
	}

	errMap.Range(func(key, value interface{}) bool {
		err = value.(error)
		if err != nil {
			return false
		}
		return true
	})

	return err
}
