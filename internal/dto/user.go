package dto

type CreateUserParams struct {
	Name     string `valid:"required"`
	Username string `valid:"required"`
	Password string `valid:"required"`
}
