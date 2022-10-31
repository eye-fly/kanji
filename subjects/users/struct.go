package users

type User struct {
	UserId int `db:"user_id,key"`
	Level  int `db:"level"`
}

func (*User) TableName() string { return UserTable }
