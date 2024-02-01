package schemas

import (
	"gorm.io/gorm"
)

// Student representa a estrutura de dados de um estudante
type Student struct {
	gorm.Model
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var Students []Student

