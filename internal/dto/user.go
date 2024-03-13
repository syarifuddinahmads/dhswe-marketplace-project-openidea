package dto

type (
	UserResponse struct {
		ID       uint   `json:"user_id"`
		Name     string `json:"name" `
		Username string `json:"username"`
	}
)
