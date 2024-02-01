package handler

// Student representa a estrutura de dados de um estudante
type Student struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

var Students []Student
