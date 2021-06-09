package models

type Role struct {
	ID          uint
	Name        string
	Permissions []int
}
