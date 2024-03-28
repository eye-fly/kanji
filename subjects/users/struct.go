package users

import "time"

type User struct {
	UserId int `db:"user_id,key"`
	Level  int `db:"level"`
}

func (*User) TableName() string { return UserTable }

type UserData struct {
	Username   string `db:"user_name,key"`
	UserId     int    `db:"user_id"`
	PaswordSha string `db:"pasword_sha"`
}

func (*UserData) TableName() string { return UserDataTable }

type Sesion struct {
	SesionID  string    `db:"sesion_id,key"`
	UserId    int       `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
}

func (*Sesion) TableName() string { return SesionTable }
