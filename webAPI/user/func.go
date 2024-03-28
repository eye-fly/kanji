package user

import (
	"net/http"

	"github.com/samonzeweb/godb"
)

func CheckCookieAndGetUserId(db *godb.DB, r *http.Request) (int, error) {
	// cookie, err := r.Cookie(users.CookieSesionIdName)
	// if err != nil {
	// 	return 0, err
	// }

	// session, err := users.IsSessionIdOk(db, cookie.Value)
	// if err == sql.ErrNoRows {
	// 	return 0, err
	// } else if err != nil {
	// 	log.Errorf("GetCheckCookie: %s", err)
	// 	return 0, err
	// }

	// return session.UserId, nil
	return 101, nil
}
