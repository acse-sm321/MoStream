package session

import (
	"MoStream/api/dbops"
	"MoStream/api/defs"
	"MoStream/api/utils"
	"sync"
	"time"
)

// outdated session should return notice

// in memory cache is fairly enough
// a secure sync mechanism
var sessionMap *sync.Map

func init() {
	sessionMap = &sync.Map{}
}

func nowInMilli() int64 {
	return time.Now().UnixNano() / 100000
}

func deleteExpiredSession(sid string) {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func LoadSessionsFromDB() {
	r, err := dbops.RetrieveAllSessions()
	if err != nil {
		return
	}

	r.Range(func(key, value interface{}) bool {
		ss := value.(*defs.SimpleSession)
		sessionMap.Store(key, ss)
		return true
	})
}

func GenerateNewSessionId(un string) string {
	id, _ := utils.NewUUID()
	ct := nowInMilli()
	ttl := ct + 30*60*1000
	ss := &defs.SimpleSession{Username: un, TTL: ttl}
	sessionMap.Store(id, ss)
	dbops.InsertSession(id, ttl, un)
	return id
}

// session expired
func IsSessionExpired(sid string) (string, bool) {
	ss, ok := sessionMap.Load(sid)
	ct := nowInMilli()
	if ok {
		if ss.(*defs.SimpleSession).TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}
		return ss.(*defs.SimpleSession).Username, false
	} else {
		ss, err := dbops.RetrieveSession(sid)
		if err != nil || ss == nil {
			return "", true
		}

		if ss.TTL < ct {
			deleteExpiredSession(sid)
			return "", true
		}

		sessionMap.Store(sid, ss)
		return ss.Username, false
	}
	return "", true
}
