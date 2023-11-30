package usecase

type CreateTaskInput struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
