package authDto

type RegisterRequest struct {
	FullName string `json:"full_name" `
	Email    string `json:"email" `
	Username string `json:"username" `
	Password string `json:"password" `
	Level    int64  `json:"level" `
}

type RegisterResponse struct {
	ID       string `json:"id" `
	Email    string `json:"email" `
	Username string `json:"username" `
	Level    int64  `json:"level" `
}

type LoginRequest struct {
	Email    string `json:"email" `
	Password string `json:"password" `
}

type LoginResponse struct {
	Email    string `json:"email" `
	Username string `json:"username" `
	Level    string `json:"level" `
	Token    string `json:"token"`
}
