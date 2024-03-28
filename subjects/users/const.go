package users

import "time"

const UserTable string = "users"
const UserDataTable string = "users_data"
const SesionTable string = "sesion_ids"

const UserIdColumn string = "user_id"
const UsernameColumn string = "user_name"
const SesionIdColumn string = "sesion_id"
const SesionCreateAtColumn string = "created_at"

const CookieSesionIdName = "sessionId"
const SesionCookieTimeSpan time.Duration = time.Hour * 24 * 30 * 2
