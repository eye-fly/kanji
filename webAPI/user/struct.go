package user

type userCredentials struct {
	Name        string `json:"name"`
	PasswordSha string `json:"password_sha"`
}
