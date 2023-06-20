package agg

type Admin interface {
	CheckPassword(pwd string) bool
	UpdatePassword(password string)
}
