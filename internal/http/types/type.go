package types

type Student struct {
	ID    int64
	Name  string `json:"name" validate:"required, min=3"`
	Email string `json:"email" validate:"required,email"`
}
