package user

type userCredentials struct {
	Name        string `json:"name"`
	PasswordSha string `json:"password_sha"`
}

type progress struct {
	Progress [5]int `json:"progress"`
}
