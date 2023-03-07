package user

import (
	"database/sql"
	"encoding/json"
	"math"
	"math/rand"
	"net/http"
	"path"
	"sql_filler/internal/utility"
	"sql_filler/subjects/assignment"
	"sql_filler/subjects/users"
	"strconv"
	"time"

	"github.com/samonzeweb/godb"
	log "github.com/sirupsen/logrus"
)

type backEnd struct {
	db *godb.DB
}

func NewBackEnd(db *godb.DB) *backEnd {
	return &backEnd{
		db: db,
	}
}

type serves struct {
	path       string
	handleFunc func(w http.ResponseWriter, r *http.Request)
}

func (bec *backEnd) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	serves := []serves{
		{"register-user", bec.serveRegisterUser},
		{"login-user", bec.serveLoginUser},
		{"progress", bec.serveLvlsProgress},
		{"pleas_log_out_because_js_cant", bec.serveLogOut},
	}

	p := r.URL.Path
	for _, srv := range serves {
		b, err := path.Match(srv.path, p)
		if err != nil {
			log.Errorf("json serveHTTP error: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		if b {
			srv.handleFunc(w, r)
			return
		}
	}
	w.WriteHeader(404)
}

func (bec *backEnd) serveRegisterUser(w http.ResponseWriter, r *http.Request) {
	var credentials userCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Errorf("userRegiser body decoder error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	is, err := users.IsUsernameUsed(bec.db, credentials.Name)
	if err != nil {
		log.Errorf("userRegiser body decoder error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if is {
		serveSimpleJson(w, "error", "sorry username: "+credentials.Name+" is used", http.StatusBadRequest)
		return
	}

	//generate unused userid
	rand.Seed(time.Now().UnixNano())
	var id int
	for true {
		id = rand.Intn(math.MaxInt)
		is, err := users.IsUserIdUsed(bec.db, id)
		if err != nil {
			log.Errorf("userRegiser id generator error: %s", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		if !is {
			break
		}
	}

	var user users.User
	var userData users.UserData
	user.Level = 1
	user.UserId = id
	userData.PaswordSha = credentials.PasswordSha
	userData.UserId = id
	userData.Username = credentials.Name

	err = bec.db.Insert(&user).Do()
	if err != nil {
		log.Errorf("userRegiser user insert error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = bec.db.Insert(&userData).Do()
	if err != nil {
		log.Errorf("userRegiser userData insert error: %s", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	serveSimpleJson(w, "status", "succses", http.StatusAccepted)
}

func (bec *backEnd) serveLoginUser(w http.ResponseWriter, r *http.Request) {
	var credentials userCredentials
	err := json.NewDecoder(r.Body).Decode(&credentials)
	if err != nil {
		log.Errorf("userRegiser body decoder error: %s", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userData, err := users.GetUserData(bec.db, credentials.Name)
	if err == sql.ErrNoRows {
		serveSimpleJson(w, "error", "Username dose not exist", http.StatusBadRequest)
		return
	}
	if err != nil {
		log.Errorf("userLogin userdata get error: %s", err)
		http.Error(w, "err", http.StatusInternalServerError)
		return
	}

	if userData.PaswordSha != credentials.PasswordSha {
		serveSimpleJson(w, "error", "Wrong password", http.StatusBadRequest)
		return
	}
	cred, err := users.CreateNewSession(bec.db, userData.UserId)
	if err != nil {
		log.Errorf("userLogin CreateNewSession error: %s", err)
		http.Error(w, "err", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     users.CookieSesionIdName,
		Value:    cred.SesionID,
		HttpOnly: true,
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	})
	serveSimpleJson(w, "status", "succses", http.StatusAccepted)
}

func (bec *backEnd) serveLogOut(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:     users.CookieSesionIdName,
		Value:    "",
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now(),
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
	})

	http.Redirect(w, r, "/user/login", http.StatusPermanentRedirect)
}

func (bec *backEnd) serveLvlProgress(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	v := r.Header.Get("Lvl")
	lvl, err := strconv.Atoi(v)
	if err != nil || lvl > 5 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	count, err := assignment.GetProgress(bec.db, userId, lvl)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Errorf("serveLvlProgress  error: %s", err)
		return
	}

	serveSimpleJson(w, "count", strconv.Itoa(count), http.StatusOK)
}

func (bec *backEnd) serveLvlsProgress(w http.ResponseWriter, r *http.Request) {
	userId, err := CheckCookieAndGetUserId(bec.db, r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var progress progress
	for i := 0; 5 > i; i++ {
		count, err := assignment.GetProgress(bec.db, userId, i+1)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Errorf("serveLvlProgress  error: %s", err)
			return
		}
		progress.Progress[i] = count
	}
	utility.ServeBodyJson(w, &progress)
}

func serveSimpleJson(w http.ResponseWriter, jsonField, jsonVal string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Write([]byte(
		"{ \"" + jsonField + "\"" + ":" +
			"\"" + jsonVal + "\"}"))
}

// func serveBodyJson(w http.ResponseWriter, jsonStr interface{}) {
// 	buf, err := json.Marshal(jsonStr)
// 	if err != nil {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		w.Write([]byte(err.Error()))
// 		log.Errorf("lesson serveQueue marshal error: %s", err)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(buf)
// }
