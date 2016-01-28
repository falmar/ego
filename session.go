package main

import (
	"net/http"
	"time"

	"github.com/nu7hatch/gouuid"
)

type Session struct {
	ID string
	w  http.ResponseWriter
	r  *http.Request
}

func getSession(w http.ResponseWriter, r *http.Request) *Session {
	sess := &Session{w: w, r: r}
	cookie, err := r.Cookie("session-id")
	if err != nil {
		SessionID, _ := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session-id",
			Value: SessionID.String(),
		}
		http.SetCookie(w, cookie)
		sess.ID = cookie.Value
	} else {
		sess.ID = cookie.Value
	}

	return sess
}

func (sess *Session) Set(key, value string, expires time.Time) {
	http.SetCookie(sess.w, &http.Cookie{
		Name:    key,
		Value:   value,
		Expires: expires,
	})
}

func (sess *Session) Get(key string) string {
	cookie, err := sess.r.Cookie(key)
	if err != nil {
		return ""
	}
	return cookie.Value
}
