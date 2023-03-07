package user

import (
	"database/sql"
	"net/http"
	"sql_filler/subjects/users"

	"github.com/samonzeweb/godb"
	log "github.com/sirupsen/logrus"
)

func CheckCookieAndGetUserId(db *godb.DB, r *http.Request) (int, error) {
	cookie, err := r.Cookie(users.CookieSesionIdName)
	if err != nil {
		return 0, err
	}

	session, err := users.IsSessionIdOk(db, cookie.Value)
	if err == sql.ErrNoRows {
		return 0, err
	} else if err != nil {
		log.Errorf("GetCheckCookie: %s", err)
		return 0, err
	}

	return session.UserId, nil
}
