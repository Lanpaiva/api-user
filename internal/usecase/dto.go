package usecase

type CreateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateUserInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
